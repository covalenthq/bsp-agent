package proof

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

var (
	proofTxTimeout uint64 = 60
)

func SendBlockReplicaProofTx(ctx context.Context, config *config.EthConfig, proofChain string, ethClient *ethclient.Client, chainHeight uint64, chainLen uint64, resultSegment []byte, txHash chan string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()

	_, opts, chainId, err := getTransactionOpts(ctx, config, ethClient)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
		return
	}
	contractAddress := common.HexToAddress(proofChain)
	contract, err := NewProofChain(contractAddress, ethClient)
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
	tx, err := contract.ProveBlockReplicaProduced(opts, chainId, chainHeight, chainLen, uint64(len(jsonResult)), sha256Result)
	if err != nil {
		log.Error("error calling deployed contract: ", err)
		return
	}
	receipt, err := bind.WaitMined(ctx, ethClient, tx)
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

func getTransactionOpts(ctx context.Context, config *config.EthConfig, ethClient *ethclient.Client) (common.Address, *bind.TransactOpts, uint64, error) {
	sKey := config.PrivateKey
	chainId, err := ethClient.ChainID(ctx)
	if err != nil {
		log.Error(err.Error())
	}
	secretKey := crypto.ToECDSAUnsafe(common.FromHex(sKey))
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(secretKey, chainId)
	if err != nil {
		log.Fatalf("error getting new keyed transactor with chain id: %v", err.Error())
	}

	return addr, opts, chainId.Uint64(), err
}
