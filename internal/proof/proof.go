package proof

import (
	"context"
	"fmt"
	"math/big"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func SubmitSpecimenProofTx(config *config.Config, chainHeight uint64, blockSpecimen event.SpecimenEvent) {

	ctx := context.Background()
	//var onlyOnce sync.Once

	sendAddr, opts, chainId, err := GetTransactionOpts(config)
	if err != nil {
		log.Error("error in getting transaction ops: %v", err)
	}

	ethclient, err := GetEthClient(config.EthConfig.ProverClient)
	if err != nil {
		log.Error("error in getting prover eth client: ", err.Error())
	}

	number, err := ethclient.BlockNumber(ctx)
	if err != nil {
		log.Error("error in getting block number: %v", err)
	}
	fmt.Println(ethclient.BalanceAt(ctx, sendAddr, new(big.Int).SetUint64(number)))

	contractAddress := common.HexToAddress(config.EthConfig.Contract)
	contract, err := NewProofChain(contractAddress, ethclient)
	if err != nil {
		log.Error("error in binding to deployed contract: %v", err)
	}

	fmt.Println("specimen:", blockSpecimen.ReplicationEvent.Hash, "sender:", sendAddr, "receiver:", common.HexToAddress(config.EthConfig.Contract), "chainid:", chainId)

	// onlyOnce.Do(func() {
	// 	WatchContractResultPublicationProof(contract)
	// })

	// jsonSpecimen, err := json.Marshal(blockSpecimen)
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	//sha256Specimen := sha256.Sum256(jsonSpecimen)

	signedTx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight)
	if err != nil {
		log.Error("error in calling deployed contract: %v", err)
	}

	fmt.Println(signedTx)

	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), secret)
	// if err != nil {
	// 	log.Error("error in signing tx: %v", err)
	// }

	// receipt, err := bind.WaitMined(ctx, ethclient, signedTx)
	// if receipt.Status != types.ReceiptStatusSuccessful {
	// 	panic("Call failed")
	// 	//log.Error("Specimen proof tx call: %v , to contract failed: %v", signedTx.Hash(), err)
	// 	//return signedTx.Hash().String(), false, err
	// }
	log.Info("Proof of specimen signed by : %v has been sent to mempool at nonce: %v", opts.Signer, opts.Nonce)

	//return signedTx.Hash().String(), true, err
}

func SubmitResultProofTx(config *config.Config, chainHeight uint64, blockResult event.ResultEvent) {

	ctx := context.Background()
	//var onlyOnce sync.Once

	sendAddr, opts, chainId, err := GetTransactionOpts(config)
	if err != nil {
		log.Error("error in getting transaction ops: %v", err)
	}

	fmt.Println("result:", blockResult.ReplicationEvent.Hash, sendAddr, chainId)

	ethclient, err := GetEthClient(config.EthConfig.ProverClient)
	if err != nil {
		log.Error("error in getting prover eth client: ", err.Error())
	}
	number, err := ethclient.BlockNumber(ctx)
	if err != nil {
		log.Error("error in getting block number: %v", err)
	}
	fmt.Println(ethclient.BalanceAt(ctx, sendAddr, new(big.Int).SetUint64(number)))

	contractAddress := common.HexToAddress(config.EthConfig.Contract)
	contract, err := NewProofChain(contractAddress, ethclient)
	if err != nil {
		log.Error("error in binding to deployed contract: %v", err)
	}

	// onlyOnce.Do(func() {
	// 	WatchContractSpecimenPublicationProof(contract)
	// })

	fmt.Println("result:", blockResult.ReplicationEvent.Hash, "sender:", sendAddr, "receiver:", common.HexToAddress(config.EthConfig.Contract), "chainid:", chainId)

	// jsonResult, err := json.Marshal(blockResult)
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	// sha256Result := sha256.Sum256(jsonResult)

	signedTx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight)
	if err != nil {
		log.Error("error in calling deployed contract: %v", err)
	}

	fmt.Println(signedTx)

	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), secret)
	// if err != nil {
	// 	log.Error("error in signing tx: %v", err)
	// }

	// receipt, err := bind.WaitMined(ctx, ethclient, signedTx)
	// if receipt.Status != types.ReceiptStatusSuccessful {
	// 	panic("Call failed")
	// 	//log.Error("Result proof tx call: %v , to contract failed: %v", signedTx.Hash(), err)
	// 	//return signedTx.Hash().String(), false, err
	// }

	log.Info("Proof of result tx signed by : %v has been sent to mempool at nonce: %v", opts.Signer, opts.Nonce)

	//return signedTx.Hash().String(), true, err
}

func GetEthClient(address string) (*ethclient.Client, error) {

	cl, err := ethclient.Dial(address)
	if err != nil {
		log.Error("error in signing tx: %v", err)
	}
	return cl, nil
}

func GetTransactionOpts(config *config.Config) (common.Address, *bind.TransactOpts, uint64, error) {

	sKey := config.EthConfig.Key
	chainId := config.EthConfig.ChainId

	secretKey := crypto.ToECDSAUnsafe(common.FromHex(sKey))
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(secretKey, new(big.Int).SetUint64(chainId))
	if err != nil {
		log.Error("error in getting new keyed transactor with chain id: ", err.Error())
	}

	return addr, opts, chainId, err
}

func GetKeyStore(config *config.Config) (*bind.TransactOpts, error) {

	chainId := config.EthConfig.ChainId

	ks := keystore.NewKeyStore(config.EthConfig.Keystore, keystore.StandardScryptN, keystore.StandardScryptP)
	accs := ks.Accounts()
	ks.Unlock(accs[0], config.EthConfig.Password)

	ksOpts, err := bind.NewKeyStoreTransactorWithChainID(ks, accs[0], new(big.Int).SetUint64(chainId))
	if err != nil {
		log.Error("error in getting transaction with chain id: ", err.Error())
	}

	return ksOpts, err
}

// func WatchContractResultPublicationProof(contract *ProofChain) {
// 	ctx := context.Background()
// 	// Watch for a Deposited event
// 	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
// 	// Setup a channel for results
// 	channel := make(chan *ProofChainBlockResultPublicationProofAppended)
// 	// Start a goroutine which watches new events
// 	go func() {
// 		sub, err := contract.WatchBlockResultPublicationProofAppended(watchOpts, channel)
// 		if err != nil {
// 			log.Error("error in watching contract for result proof event: ", err.Error())
// 		}
// 		defer sub.Unsubscribe()
// 	}()
// 	// Receive events from the channel
// 	event := <-channel

// 	log.Info("new result event emitted from prover contract: %v", event)
// }

// func WatchContractSpecimenPublicationProof(contract *ProofChain) {
// 	ctx := context.Background()
// 	// Watch for a Deposited event
// 	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
// 	// Setup a channel for results
// 	channel := make(chan *ProofChainBlockSpecimenPublicationProofAppended)
// 	// Start a goroutine which watches new events
// 	go func() {
// 		sub, err := contract.WatchBlockSpecimenPublicationProofAppended(watchOpts, channel)
// 		if err != nil {
// 			log.Error("error in watching contract for specimen proof event: ", err.Error())
// 		}
// 		defer sub.Unsubscribe()
// 	}()
// 	// Receive events from the channel
// 	event := <-channel

// 	log.Info("new specimen event emitted from prover contract: %v", event)
// }
