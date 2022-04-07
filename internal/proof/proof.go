// Package proof contains all the functions to make a proof on the proofchain about a block replica
//nolint:wrapcheck
package proof

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
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

var (
	proofTxTimeout uint64 = 60
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
	contract, err := NewProofChain(contractAddress, ethClient)
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

	transaction, err := contract.SubmitBlockSpecimenProof(opts, blockReplica.NetworkId, chainHeight, blockReplica.Hash, sha256Result, replicaURL)

	if err != nil {
		log.Error("error calling deployed contract: ", err)
		txHash <- ""

		return
	}
	receipt, err := bind.WaitMined(ctx, ethClient, transaction)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("block-result proof tx call: ", transaction.Hash(), " to proof contract failed: ", err.Error())
		txHash <- ""

		return
	}
	if err != nil {
		log.Error("error in waiting for tx to be mined on the blockchain: ", err.Error())
		txHash <- ""

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

// SendPackerProofTx sends a batched tx proof to Ext network
func SendPackerProofTx(ctx context.Context, config *config.EthConfig, ethClient *ethclient.Client, packerTxBatch []string) (string, error) {
	packerTxAddress := common.HexToAddress("0xABEFF04FC17983135Ad91b9f8B03da7d7Ba9C83e")

	var data []byte
	data = append(data, ':')
	for _, tx := range packerTxBatch {
		data = append(data, tx...)
		data = append(data, ',')
	}
	txs := string(data)

	address, _, chainID, err := getTransactionOpts(ctx, config, ethClient)
	if err != nil {
		log.Error("error getting transaction ops: ", err.Error())
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), packerTxAddress)
	if err != nil {
		log.Error("error in getting transaction nonce: ", err.Error())
	}

	senderPrivKey, _ := crypto.HexToECDSA(config.PrivateKey)

	// Construct the tx
	packTx := types.NewTransaction(
		nonce,           // nonce
		packerTxAddress, // to
		new(big.Int),    // amount
		uint64(21000),   // gas limit
		new(big.Int),    // gas price
		[]byte(txs),     // data
	)

	// Sign the tx
	signer := types.NewEIP155Signer(big.NewInt(int64(chainID)))
	signedTx, _ := types.SignTx(packTx, signer, senderPrivKey)
	fmt.Println("tx payload:", packTx, "from address", address)

	// var buff bytes.Buffer
	// signedTx.EncodeRLP(&buff)
	// fmt.Printf("0x%x\n", buff.Bytes())

	// Send the transaction to the network
	txErr := ethClient.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		log.Error("send tx error:", txErr)
	}

	fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())

	return signedTx.Hash().String(), txErr
}
