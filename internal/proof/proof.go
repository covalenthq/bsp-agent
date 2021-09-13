package proof

import (
	"context"
	"math/big"

	ty "github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

func buildSpecimenProofTx(client *ethclient.Client, contractAddress string, chainId int, hash string, blockSpecimen *ty.BlockSpecimen) (string, error) {
	//todo
	//return string, error

	addr := common.HexToAddress(contractAddress)
	ctr, err := contract.NewContract(addr, client)

	if err != nil {
		log.Error(err.Error())
	}

}

func buildResultProofTx(chainId int, hash string, blockResult *ty.BlockResult) (string, error) {
	//todo
	//return string, error
}

func getClient(address string) *ethclient.Client {

	cl, err := ethclient.Dial(address)
	if err != nil {
		log.Error(err.Error())
	}
	return cl
}

func createProofTx(ethclient *ethclient.Client, sender common.Address, receiver *common.Address) error {
	ctx := context.Background()
	// Use infura
	infura := "wss://goerli.infura.io/ws/v3/xxxxxx"

	client := getClient(infura)
	gasLimit := uint64(21000)
	data := []byte{}
	amount := big.NewInt(10 * params.GWei)

	estimateGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From: sender,
		To:   receiver,
		Gas:  gasLimit,
		Data: data,
	})
	if err != nil {
		log.Error(err.Error())
	}

	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		log.Error(err.Error())
	}

	nonce, err := client.NonceAt(ctx, sender, big.NewInt(int64(blockNum)))
	if err != nil {
		log.Error(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Error(err.Error())
	}

	tx := types.NewTransaction(nonce, *receiver, amount, estimateGas, gasPrice, data)

	err = signProofTx(ctx, tx, client, sender)

	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func signProofTx(ctx context.Context, tx *types.Transaction, client *ethclient.Client, address common.Address) error {

	notSecretKey := "0x0000" //replace with env with secret key

	secretKey := crypto.ToECDSAUnsafe(common.FromHex(notSecretKey))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), secretKey)

	opts := bind.NewKeyedTransactor(secretKey)
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Proof of specimen tx for: %v signed by : %v has been sent to mempool at nonce: %v", addr, opts.Signer, opts.Nonce)

	return err
}

func getKeyStore(ctx context.Context, tx *types.Transaction, client *ethclient.Client, address common.Address) (*bind.TransactOpts, error) {

	ks := keystore.NewKeyStore("/path/to/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// acc, err := ks.NewAccount("passwordToNewAccount")
	accs := ks.Accounts()
	ks.Unlock(accs[0], "passwordToNewAccount")
	ksOpts, err := bind.NewKeyStoreTransactor(ks, accs[0])

	return ksOpts, err
}
