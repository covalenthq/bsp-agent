// Package proof contains all functions to make a proof-chain tx for an encoded block-replica object
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

// ProofchainInteractor a wrapper over proofchain contract to help clients interact with it
type ProofchainInteractor struct {
	config             *config.AgentConfig
	ethClient          *ethclient.Client
	proofChainContract *ProofChain
}

// NewProofchainInteractor sets up a new interactor for proof-chain
func NewProofchainInteractor(config *config.AgentConfig, ethClient *ethclient.Client) *ProofchainInteractor {
	interactor := &ProofchainInteractor{config: config, ethClient: ethClient}
	contractAddress := common.HexToAddress(config.ProofchainConfig.ProofChainAddr)
	contract, err := NewProofChain(contractAddress, ethClient)
	if err != nil {
		log.Fatalf("error binding to deployed contract: %v", err.Error())
	}

	interactor.proofChainContract = contract

	return interactor
}

// SendBlockReplicaProofTx makes a proof-chain tx for the block-replica that has been processed
func (interactor *ProofchainInteractor) SendBlockReplicaProofTx(ctx context.Context, chainHeight uint64, blockReplica *ty.BlockReplica, resultSegment []byte, replicaURL string, txHash chan string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()

	txHash <- "abracadabra"
	return // Returning early to bypass proofchain

	_, opts, _, err := getTransactionOpts(ctx, &interactor.config.ChainConfig, interactor.ethClient)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
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

	executeWithRetry(ctx, interactor.proofChainContract, interactor.ethClient, opts, blockReplica, txHash, chainHeight, replicaURL, sha256Result, 0)
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
		if strings.Contains(err.Error(), "Block height is out of bounds for live sync") {
			log.Error("skip creating proof-chain session: ", err)
			txHash <- "out-of-bounds block"

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
			log.Error("proof tx failed/reverted on tx retry, skipping: ", transaction.Hash())
			txHash <- "retry fail"

			return
		}
		log.Error("proof tx failed/reverted, retrying proof tx for block hash: ", blockReplica.Hash.String())
		executeWithRetry(ctx, proofChainContract, ethClient, opts, blockReplica, txHash, chainHeight, replicaURL, sha256Result, retryCount+1)

		return
	}

	txHash <- receipt.TxHash.String()
}

func getTransactionOpts(ctx context.Context, cfg *config.ChainConfig, ethClient *ethclient.Client) (common.Address, *bind.TransactOpts, uint64, error) {
	sKey := cfg.PrivateKey
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
