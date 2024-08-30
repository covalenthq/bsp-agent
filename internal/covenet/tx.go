package covenet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/covalenthq/bsp-agent/internal/config"
	covapp "github.com/covalenthq/covenet/app"
	covparams "github.com/covalenthq/covenet/app/params"
	covtypes "github.com/covalenthq/covenet/x/covenet/types"

	bsptypes "github.com/covalenthq/bsp-agent/internal/types"
)

var encCfg covparams.EncodingConfig

const (
	proofTxTimeout  uint64 = 480
	retryCountLimit int    = 3 // 3 retry for covenet proofchain submission
)

func init() {
	encCfg = covapp.MakeEncodingConfig()
}

type CovenetInteractor struct {
	config         *config.AgentConfig
	grpcClient     *grpc.ClientConn
	pubKey         cryptotypes.PubKey
	address        sdk.AccAddress
	accountNumber  uint64
	sequenceNumber uint64
	sequenceMutex  sync.Mutex
}

// Set Private Key from Config Env Var Hex bytes
func (interactor *CovenetInteractor) getPrivateKey() (cryptotypes.PrivKey, error) {
	privKeyBytes, err := hex.DecodeString(interactor.config.CovenetConfig.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key: %v", err)
	}
	return &secp256k1.PrivKey{Key: privKeyBytes}, nil
}

// NewCovenetInteract sets up a new operator-interactor for covenet
func NewCovenetInteractor(config *config.AgentConfig) (*CovenetInteractor, error) {
	grpcConn, err := GetGRPCConnection(config)
	if err != nil {
		return nil, err
	}

	interactor := &CovenetInteractor{
		config:         config,
		grpcClient:     grpcConn,
		pubKey:         &secp256k1.PubKey{},
		address:        sdk.AccAddress{},
		accountNumber:  0,
		sequenceNumber: 0,
		sequenceMutex:  sync.Mutex{},
	}

	// Process the key and set up the interactor account
	err = interactor.ProcessKey()
	if err != nil {
		return nil, fmt.Errorf("failed to process key: %v", err)
	}

	// Cache initial account info
	sequenceNum, accountNum, err := interactor.GetAccountInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get initial account info: %v", err)
	}
	interactor.sequenceNumber = sequenceNum
	interactor.accountNumber = accountNum

	return interactor, nil
}

func (interactor *CovenetInteractor) GetSystemInfo() (*covtypes.SystemInfo, error) {
	// This creates a gRPC client to query the x/covenet service.
	covenetClient := covtypes.NewQueryClient(interactor.grpcClient)
	params := &covtypes.QueryGetSystemInfoRequest{}

	res, err := covenetClient.SystemInfo(context.Background(), params)
	if err != nil {
		return nil, err
	}
	log.Info("System Info: ", res)

	return &res.SystemInfo, nil
}

func GetGRPCConnection(config *config.AgentConfig) (*grpc.ClientConn, error) {
	// Create a connection to the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcConn, err := grpc.DialContext(
		ctx,
		"covenet-node:9090", //config.CovenetConfig.GRPCURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	log.Info("GRPC connection status: ", grpcConn.GetState())

	return grpcConn, nil
}

func (interactor *CovenetInteractor) ProcessKey() error {
	// Set the bech32 prefix for your chain
	const (
		Bech32PrefixAccAddr = "cxtmos"
	)

	// Configure the address
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccAddr+"pub")

	// Decode the hex string to bytes
	privKeyBytes, err := hex.DecodeString(interactor.config.CovenetConfig.PrivateKey)
	if err != nil {
		return fmt.Errorf("error decoding private key: %v", err)
	}

	// Create a new PrivKey object
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}

	// Derive public key
	pubKey := privKey.PubKey()

	// Get the address bytes
	addrBytes := sdk.AccAddress(pubKey.Address())

	// Encode with the correct prefix
	bech32Addr, err := bech32.ConvertAndEncode(Bech32PrefixAccAddr, addrBytes)
	if err != nil {
		return fmt.Errorf("error encoding bech32 address: %v", err)
	}

	// If you specifically need a Covenet address type
	covenetAddr, err := covtypes.CovenetAccAddressFromBech32(bech32Addr)
	if err != nil {
		return fmt.Errorf("error converting to Covenet address: %v", err)
	}

	// Set interactor key address values
	interactor.pubKey = pubKey
	interactor.address = covenetAddr

	return nil
}

func (interactor *CovenetInteractor) GetLatestSequence() uint64 {
	interactor.sequenceMutex.Lock()
	defer interactor.sequenceMutex.Unlock()
	return interactor.sequenceNumber
}

func (interactor *CovenetInteractor) IncrementSequence() {
	interactor.sequenceMutex.Lock()
	defer interactor.sequenceMutex.Unlock()
	interactor.sequenceNumber++
}

func (interactor *CovenetInteractor) GetAccountInfo() (uint64, uint64, error) {
	// Create a new auth query client
	authClient := authtypes.NewQueryClient(interactor.grpcClient)

	// Query the account
	res, err := authClient.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{Address: interactor.address.String()},
	)
	if err != nil {
		// If the account is not found return 0 values with error
		return 0, 0, fmt.Errorf("failed to query account %s: %v", interactor.address.String(), err)
	}

	// Create a new AccountI interface
	var account authtypes.AccountI

	// Unmarshal the account data
	err = encCfg.InterfaceRegistry.UnpackAny(res.Account, &account)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to unpack account %s: %v", interactor.address.String(), err)
	}

	// Get the sequence number
	accountSequence := account.GetSequence()
	accountNumber := account.GetAccountNumber()

	return accountSequence, accountNumber, nil
}

// SendBlockReplicaProofTx makes a proof-chain tx for the block-replica that has been processed
func (interactor *CovenetInteractor) SendCovenetBlockReplicaProofTx(ctx context.Context, chainHeight uint64, blockReplica *bsptypes.BlockReplica, resultSegment []byte, replicaURL string, txHash chan string) {
	// Empty error used for recursive retry tx call
	var emptyErr error
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()

	jsonResult, err := json.Marshal(resultSegment)
	if err != nil {
		log.Error("error in JSON marshaling result segment: ", err.Error())
		txHash <- ""

		return
	}
	sha256Result := sha256.Sum256(jsonResult)

	// Call Create Proof With Retry Count at 0 and Empty Err
	err = interactor.CreateProofTxWithRetry(ctx, blockReplica, txHash, chainHeight, replicaURL, sha256Result, 0, emptyErr)
	if err != nil {
		log.Error("covenet tx failed: ", err.Error())
	}
}

// CreateProofTx creates a proof transaction on covenet and return errors on failuer
func (interactor *CovenetInteractor) CreateProofTx(ctx context.Context, blockReplica *bsptypes.BlockReplica, txHash chan string, chainHeight uint64, replicaURL string, sha256Result [sha256.Size]byte) error {

	// Get account from private key
	privK, err := interactor.getPrivateKey()
	if err != nil {
		return fmt.Errorf("failed to get private key for address %s: %v", interactor.address.String(), err)
	}

	sequence := interactor.GetLatestSequence()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	// Create Msg from covenet types
	proofMsg := covtypes.NewMsgCreateProof(interactor.address.String(), int32(blockReplica.NetworkId), "specimen", chainHeight, blockReplica.Hash.String(), hex.EncodeToString(sha256Result[:]), replicaURL)

	err = txBuilder.SetMsgs(proofMsg)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(200000)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	// Assuming we have a single private key, account number, and sequence
	log.Info("account details: ", interactor.address.String(), " number: ", interactor.accountNumber, " nonce: ", sequence)

	sigV2 := signing.SignatureV2{
		PubKey: interactor.pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: sequence,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Second round: actual signing
	signerData := xauthsigning.SignerData{
		ChainID:       "covenet",
		AccountNumber: interactor.accountNumber,
		Sequence:      sequence,
	}
	sigV2, err = clientTx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, privK, encCfg.TxConfig, sequence)
	if err != nil {
		return err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())

	// Broadcast the tx via gRPC.
	// We create a new client for the Protobuf Tx service.
	txClient := tx.NewServiceClient(interactor.grpcClient)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return err
	}

	log.Info("response code\n", grpcRes.TxResponse.String()) // Should be `0` if the tx is successful

	if grpcRes.TxResponse.Code == 0 {
		interactor.IncrementSequence()
		txHash <- grpcRes.TxResponse.TxHash
		return nil
	}

	return fmt.Errorf("transaction failed with code %d: %s", grpcRes.TxResponse.Code, grpcRes.TxResponse.String())

}

func (interactor *CovenetInteractor) CreateProofTxWithRetry(ctx context.Context, blockReplica *bsptypes.BlockReplica, txHash chan string, chainHeight uint64, replicaURL string, sha256Result [sha256.Size]byte, retryCount int, lastError error) error {

	if retryCount >= retryCountLimit {
		errStr := lastError.Error()
		switch {
		case strings.Contains(errStr, "submitted tx creator is already session member"):
			log.Warn("skipping: covenet creator is already a session member")
			txHash <- "presubmitted hash"
		case strings.Contains(errStr, "proof session submitted out of acceptable live bounds"):
			log.Warn("skipping: covenet Proof session out of acceptable live bounds")
			txHash <- "out-of-bounds block"
		// Add additional cases that need skipping based on response from covenet
		// case strings.Contains(errStr, "the client connection is closing"):
		// 	log.Warn("skipping: covenet client connection is closing")
		// 	txHash <- "mine timeout"
		default:
			log.Error("too many errors in creating proof on covenet: ", errStr)
			txHash <- ""
		}
		return fmt.Errorf("exceeded retry limit of %d attempts, with response %s", retryCountLimit, lastError)
	}

	err := interactor.CreateProofTx(ctx, blockReplica, txHash, chainHeight, replicaURL, sha256Result)
	if err != nil {
		lastError = err
		log.Error(err)
	} else {
		return nil
	}

	// Exponential backoff
	backoffDuration := time.Duration(1<<uint(retryCount)) * time.Second
	log.Info("Retrying Create Proof tx in: ", backoffDuration)
	time.Sleep(backoffDuration)

	// Recursive call with now incremented retry count
	return interactor.CreateProofTxWithRetry(ctx, blockReplica, txHash, chainHeight, replicaURL, sha256Result, retryCount+1, lastError)
}
