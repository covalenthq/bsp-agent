package proof

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"math/big"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func SubmitSpecimenProofTx(config *config.Config, ethclient *ethclient.Client, chainHeight uint64, chainLength uint64, blockSpecimen event.SpecimenEvent) (string, bool, error) {

	ctx := context.Background()
	addr, opts, chainId, secret, err := GetTransactionOpts(config)
	if err != nil {
		log.Error(err.Error())
	}

	contract, err := NewProofChain(addr, ethclient)
	if err != nil {
		log.Error(err.Error())
	}

	jsonSpecimen, err := json.Marshal(blockSpecimen)
	if err != nil {
		log.Error(err.Error())
	}
	sha256Specimen := sha256.Sum256(jsonSpecimen)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLength, uint64(len(jsonSpecimen)), sha256Specimen)
	if err != nil {
		log.Error(err.Error())
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), secret)
	if err != nil {
		log.Error(err.Error())
	}

	receipt, err := bind.WaitMined(ctx, ethclient, signedTx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Specimen proof tx call: %v , to contract failed: %v", tx.Hash(), err)
		return signedTx.Hash().String(), false, err
	}
	log.Info("Proof of specimen tx on: %v signed by : %v has been sent to mempool at nonce: %v", addr, opts.Signer, opts.Nonce)

	return signedTx.Hash().String(), true, err
}

func SubmitResultProofTx(config *config.Config, ethclient *ethclient.Client, chainHeight uint64, chainLength uint64, blockResult event.ResultEvent) (string, bool, error) {

	ctx := context.Background()

	addr, opts, chainId, secret, err := GetTransactionOpts(config)
	if err != nil {
		log.Error(err.Error())
	}

	contract, err := NewProofChain(addr, ethclient)
	if err != nil {
		log.Error(err.Error())
	}

	jsonResult, err := json.Marshal(blockResult)
	if err != nil {
		log.Error(err.Error())
	}
	sha256Result := sha256.Sum256(jsonResult)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLength, uint64(len(jsonResult)), sha256Result)
	if err != nil {
		log.Error(err.Error())
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), secret)
	if err != nil {
		log.Error(err.Error())
	}

	receipt, err := bind.WaitMined(ctx, ethclient, signedTx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Result proof tx call: %v , to contract failed: %v", tx.Hash(), err)
		return signedTx.Hash().String(), false, err
	}

	log.Info("Proof of result tx on: %v signed by : %v has been sent to mempool at nonce: %v", addr, opts.Signer, opts.Nonce)

	return signedTx.Hash().String(), true, err
}

func GetEthClient(config *config.Config) *ethclient.Client {

	cl, err := ethclient.Dial(config.EthConfig.Client)
	if err != nil {
		log.Error(err.Error())
	}
	return cl
}

func GetTransactionOpts(config *config.Config) (common.Address, *bind.TransactOpts, uint64, *ecdsa.PrivateKey, error) {

	sKey := config.EthConfig.Key
	chainId := config.EthConfig.ChainId

	secretKey := crypto.ToECDSAUnsafe(common.FromHex(sKey))
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(secretKey, new(big.Int).SetUint64(chainId))
	if err != nil {
		log.Error(err.Error())
	}

	return addr, opts, chainId, secretKey, err
}

func GetKeyStore(config *config.Config) (*bind.TransactOpts, error) {

	chainId := config.EthConfig.ChainId

	ks := keystore.NewKeyStore(config.EthConfig.Keystore, keystore.StandardScryptN, keystore.StandardScryptP)
	accs := ks.Accounts()
	ks.Unlock(accs[0], config.EthConfig.Password)

	ksOpts, err := bind.NewKeyStoreTransactorWithChainID(ks, accs[0], new(big.Int).SetUint64(chainId))

	return ksOpts, err
}
