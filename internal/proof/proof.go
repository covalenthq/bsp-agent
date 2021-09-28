package proof

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

var (
	proofTxTimeout uint64 = 60
)

func SendBlockSpecimenProofTx(ctx context.Context, config *config.EthConfig, ethProof *ethclient.Client, chainHeight uint64, chainLen uint64, specimenSegment []byte, txHash chan string) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()
	//var onlyOnce sync.Once

	_, opts, chainId, err := getTransactionOpts(config)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
		return
	}

	contractAddress := common.HexToAddress(config.Contract)
	contract, err := NewProofChain(contractAddress, ethProof)

	if err != nil {
		log.Error("error binding to deployed contract: ", err.Error())
		return
	}

	// onlyOnce.Do(func() {
	// 	WatchContractResultPublicationProof(ctx, contract)
	// })

	jsonSpecimen, err := json.Marshal(specimenSegment)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sha256Specimen := sha256.Sum256(jsonSpecimen)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLen, uint64(len(jsonSpecimen)), sha256Specimen)
	if err != nil {
		log.Error("error calling contract function: ", err.Error())
		return
	}

	receipt, err := bind.WaitMined(ctx, ethProof, tx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("block-specimen proof tx call: ", tx.Hash(), " to proof contract failed: ", err)
		return
	}
	if err != nil {
		log.Error(err.Error())
		return
	}

	txHash <- receipt.TxHash.String()

}

func SendBlockResultProofTx(ctx context.Context, config *config.EthConfig, ethProof *ethclient.Client, chainHeight uint64, chainLen uint64, resultSegment []byte, txHash chan string) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()
	//var onlyOnce sync.Once

	_, opts, chainId, err := getTransactionOpts(config)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
		return
	}

	contractAddress := common.HexToAddress(config.Contract)
	contract, err := NewProofChain(contractAddress, ethProof)
	if err != nil {
		log.Error("error binding to deployed contract: ", err.Error())
		return
	}

	// onlyOnce.Do(func() {
	// 	WatchContractSpecimenPublicationProof(ctx, contract)
	// })

	jsonResult, err := json.Marshal(resultSegment)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sha256Result := sha256.Sum256(jsonResult)

	tx, err := contract.ProveBlockSpecimenProduced(opts, uint64(chainId), chainHeight, chainLen, uint64(len(jsonResult)), sha256Result)
	if err != nil {
		log.Error("error calling deployed contract: ", err)
		return
	}

	receipt, err := bind.WaitMined(ctx, ethProof, tx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("block-result proof tx call: ", tx.Hash(), " to proof contract failed: ", err.Error())
		return
	}
	if err != nil {
		log.Error(err.Error())
		return
	}

	txHash <- receipt.TxHash.String()
}

func getTransactionOpts(config *config.EthConfig) (common.Address, *bind.TransactOpts, uint64, error) {

	sKey := config.Key
	chainId := config.ChainId

	secretKey := crypto.ToECDSAUnsafe(common.FromHex(sKey))
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(secretKey, new(big.Int).SetUint64(chainId))
	if err != nil {
		log.Fatalf("error getting new keyed transactor with chain id: ", err.Error())
	}

	return addr, opts, chainId, err
}

func getKeyStore(config *config.Config) (*bind.TransactOpts, error) {

	chainId := config.EthConfig.ChainId
	ks := keystore.NewKeyStore(config.EthConfig.Keystore, keystore.StandardScryptN, keystore.StandardScryptP)

	accs := ks.Accounts()
	ks.Unlock(accs[0], config.EthConfig.Password)

	ksOpts, err := bind.NewKeyStoreTransactorWithChainID(ks, accs[0], new(big.Int).SetUint64(chainId))
	if err != nil {
		log.Error("error getting new key store transactor with chain id: ", err.Error())
		return nil, err
	}

	return ksOpts, err
}

func watchContractResultPublicationProof(ctx context.Context, contract *ProofChain) {

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	channel := make(chan *ProofChainBlockResultPublicationProofAppended)

	go func() {
		sub, err := contract.WatchBlockResultPublicationProofAppended(watchOpts, channel)
		if err != nil {
			log.Error("error watching contract for result proof event: ", err.Error())
		}
		defer sub.Unsubscribe()
	}()

	event := <-channel
	log.Info("New result event emitted from prover contract: ", event)

}

func watchContractSpecimenPublicationProof(ctx context.Context, contract *ProofChain) {

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	channel := make(chan *ProofChainBlockSpecimenPublicationProofAppended)

	go func() {
		sub, err := contract.WatchBlockSpecimenPublicationProofAppended(watchOpts, channel)
		if err != nil {
			log.Error("error watching contract for specimen proof event: ", err.Error())
		}
		defer sub.Unsubscribe()
	}()

	event := <-channel
	log.Info("New specimen event emitted from prover contract: ", event)
}
