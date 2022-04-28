// Package proof contains all the functions to make a proof on the proofchain about a block replica
package proof

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"strings"
	"time"

	"github.com/covalenthq/bsp-agent/internal/config"
	ty "github.com/covalenthq/bsp-agent/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const (
	proofTxTimeout  uint64 = 301
	retryCountLimit int    = 1 // 1 retry for proofchain submission
)

// SendBlockReplicaProofTx calls the proof-chain contract to make a transaction for the block-replica that it is processing
func SendBlockReplicaProofTx(ctx context.Context, config *config.EthConfig, proofChain string, ethClient *ethclient.Client, chainHeight uint64, chainLen uint64, resultSegment []byte, replicaURL string, blockReplica *ty.BlockReplica, txHash chan string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()

	_, opts, _, err := getTransactionOpts(ctx, config, ethClient)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
		txHash <- ""

		return
	}

	contractAddress := common.HexToAddress(proofChain)
	proofChainContract, err := NewProofChain(contractAddress, ethClient)
	if err != nil {
		log.Error("error binding to deployed contract: ", err.Error())
		txHash <- ""

		return
	}

	jsonResult, err := json.Marshal(resultSegment)
	if err != nil {
		log.Error("error in JSON marshaling result segment: ", err.Error())
		txHash <- ""

		return
	}
	sha256Result := sha256.Sum256(jsonResult)

	executeWithRetry(ctx, proofChainContract, ethClient, opts, blockReplica, txHash, chainHeight, replicaURL, sha256Result, 0)
}

func executeWithRetry(ctx context.Context, proofChainContract *ProofChain, ethClient *ethclient.Client, opts *bind.TransactOpts, blockReplica *ty.BlockReplica, txHash chan string, chainHeight uint64, replicaURL string, sha256Result [sha256.Size]byte, retryCount int) {
	transaction, err := proofChainContract.SubmitBlockSpecimenProof(opts, blockReplica.NetworkId, chainHeight, blockReplica.Hash, sha256Result, replicaURL)

	if err != nil {
		if strings.Contains(err.Error(), "Session submissions have closed") {
			log.Error("skip creating proof-chain session: ", err)
			txHash <- "session closed"

			return
		}
		if strings.Contains(err.Error(), "Operator already submitted for the provided block hash") {
			log.Error("skip creating proof-chain session: ", err)
			txHash <- "presubmitted hash"

			return
		}
		log.Error("error sending tx to deployed contract: ", err)
		txHash <- ""

		return
	}

	receipt, err := bind.WaitMined(ctx, ethClient, transaction)
	if err != nil {
		log.Error("proof tx wait on mine timeout in seconds: ", proofTxTimeout, " with err: ", err.Error())
		txHash <- "mine timeout"

		return
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		if retryCount >= retryCountLimit {
			log.Error("proof tx failed/revereted on tx retry, skipping: ", transaction.Hash())
			txHash <- transaction.Hash().String()

			return
		}
		log.Error("proof tx failed/revereted, retrying proof tx for block hash: ", blockReplica.Hash)
		executeWithRetry(ctx, proofChainContract, ethClient, opts, blockReplica, txHash, chainHeight, replicaURL, sha256Result, retryCount+1)

		return
	}

	txHash <- receipt.TxHash.String()
}

func getTransactionOpts(ctx context.Context, config *config.EthConfig, ethClient *ethclient.Client) (common.Address, *bind.TransactOpts, uint64, error) {
	sKey := config.PrivateKey
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		log.Error("error in getting transaction options: ", err.Error())
	}
	secretKey := crypto.ToECDSAUnsafe(common.FromHex(sKey))
	addr := crypto.PubkeyToAddress(secretKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(secretKey, chainID)
	if err != nil {
		log.Fatalf("error getting new keyed transactor with chain id: %v", err)
	}

	return addr, opts, chainID.Uint64(), err
}
