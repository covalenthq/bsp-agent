package proof

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"time"

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

var (
	proofTxTimeout uint64 = 60
	chainLen       uint64 = 1
)

func SendBlockSpecimenProofTx(ctx context.Context, config *config.Config, ethProof *ethclient.Client, chainHeight uint64, blockSpecimen event.SpecimenEvent, txHash chan string) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()
	//var onlyOnce sync.Once

	_, opts, chainId, err := GetTransactionOpts(config)
	if err != nil {
		log.Error("error in getting transaction ops: %v", err.Error())
	}

	contractAddress := common.HexToAddress(config.EthConfig.Contract)
	contract, err := NewProofChain(contractAddress, ethProof)

	if err != nil {
		log.Error("error in binding to deployed contract: %v", err.Error())
	}

	// onlyOnce.Do(func() {
	// 	WatchContractResultPublicationProof(ctx, contract)
	// })

	jsonSpecimen, err := json.Marshal(blockSpecimen)
	if err != nil {
		log.Error(err.Error())
	}
	sha256Specimen := sha256.Sum256(jsonSpecimen)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLen, uint64(len(jsonSpecimen)), sha256Specimen)
	if err != nil {
		log.Error("error in calling deployed contract: %v", err.Error())
	}

	receipt, err := bind.WaitMined(ctx, ethProof, tx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Block-specimen proof tx call: ", tx.Hash(), " to proof contract failed: ", err)
	}
	if err != nil {
		log.Error(err.Error())
	}

	txHash <- receipt.TxHash.String()

}

func SendBlockResultProofTx(ctx context.Context, config *config.Config, ethProof *ethclient.Client, chainHeight uint64, blockResult event.ResultEvent, txHash chan string) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()
	//var onlyOnce sync.Once

	_, opts, chainId, err := GetTransactionOpts(config)
	if err != nil {
		log.Error("error in getting transaction ops: %v", err.Error())
	}

	contractAddress := common.HexToAddress(config.EthConfig.Contract)
	contract, err := NewProofChain(contractAddress, ethProof)
	if err != nil {
		log.Error("error in binding to deployed contract: %v", err.Error())
	}

	// onlyOnce.Do(func() {
	// 	WatchContractSpecimenPublicationProof(ctx, contract)
	// })

	jsonResult, err := json.Marshal(blockResult)
	if err != nil {
		log.Error(err.Error())
	}
	sha256Result := sha256.Sum256(jsonResult)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLen, uint64(len(jsonResult)), sha256Result)
	if err != nil {
		log.Error("error in calling deployed contract: %v", err)
	}

	receipt, err := bind.WaitMined(ctx, ethProof, tx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Block-result proof tx call: ", tx.Hash(), " to proof contract failed: ", err.Error())
	}
	if err != nil {
		log.Error(err.Error())
	}

	txHash <- receipt.TxHash.String()
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

func WatchContractResultPublicationProof(ctx context.Context, contract *ProofChain) {

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	channel := make(chan *ProofChainBlockResultPublicationProofAppended)

	go func() {
		sub, err := contract.WatchBlockResultPublicationProofAppended(watchOpts, channel)
		if err != nil {
			log.Error("error in watching contract for result proof event: ", err.Error())
		}
		defer sub.Unsubscribe()
	}()

	event := <-channel
	log.Info("new result event emitted from prover contract: %v", event)

}

func WatchContractSpecimenPublicationProof(ctx context.Context, contract *ProofChain) {

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	channel := make(chan *ProofChainBlockSpecimenPublicationProofAppended)

	go func() {
		sub, err := contract.WatchBlockSpecimenPublicationProofAppended(watchOpts, channel)
		if err != nil {
			log.Error("error in watching contract for specimen proof event: ", err.Error())
		}
		defer sub.Unsubscribe()
	}()

	event := <-channel
	log.Info("new specimen event emitted from prover contract: %v", event)
}
