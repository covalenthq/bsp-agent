package covenet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/covalenthq/bsp-agent/internal/config"
	app "github.com/covalenthq/covenet/app"
	params "github.com/covalenthq/covenet/app/params"
	covenet "github.com/covalenthq/covenet/x/covenet/types"

	ty "github.com/covalenthq/bsp-agent/internal/types"
)

var encCfg params.EncodingConfig

const (
	proofTxTimeout  uint64 = 480
	retryCountLimit int    = 1 // 1 retry for proofchain submission
)

func init() {
	encCfg = app.MakeEncodingConfig()
}

type CovenetInteractor struct {
	config     *config.AgentConfig
	grpcClient *grpc.ClientConn
	pubKey     cryptotypes.PubKey
	address    sdk.AccAddress
}

// NewCovenetInteract sets up a new interactor for covenet
func NewCovenetInteractor(config *config.AgentConfig) (*CovenetInteractor, error) {
	grpcConn, err := GetGRPCConnection(config)
	if err != nil {
		return nil, err
	}

	interactor := &CovenetInteractor{
		config:     config,
		grpcClient: grpcConn,
		pubKey:     &secp256k1.PubKey{},
		address:    sdk.AccAddress{},
	}

	return interactor, nil
}

func (interactor *CovenetInteractor) GetSystemInfo() (*covenet.SystemInfo, error) {
	// This creates a gRPC client to query the x/covenet service.
	covenetClient := covenet.NewQueryClient(interactor.grpcClient)
	params := &covenet.QueryGetSystemInfoRequest{}

	res, err := covenetClient.SystemInfo(context.Background(), params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("System Info: %s\n", res)

	return &res.SystemInfo, nil
}

func GetGRPCConnection(config *config.AgentConfig) (*grpc.ClientConn, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"covenet-node:9090", //config.CovenetConfig.GRPCURL, // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}

	log.Info("GRPC connection status:", grpcConn.GetState())

	return grpcConn, nil
}

func (interactor *CovenetInteractor) ProcessKey() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress, error) {
	// log.Info("this is the private key from env:", interactor.config.CovenetConfig.PrivateKey)
	// if len(interactor.config.CovenetConfig.PrivateKey) != 1 {
	// 	return nil, nil, nil, fmt.Errorf("expected 1 private keys, got %d", len(interactor.config.CovenetConfig.PrivateKey))
	// }

	// Set the bech32 prefix for your chain
	const (
		Bech32PrefixAccAddr = "cxtmos"
	)

	// Configure the address prefix
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccAddr+"pub")

	// Decode the hex string to bytes
	privKeyBytes, err := hex.DecodeString(interactor.config.CovenetConfig.PrivateKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error decoding private key: %v", err)
	}

	// Create a new PrivKey object
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}

	// Derive public key
	pubKey := privKey.PubKey()

	// Get the address bytes
	addrBytes := types.AccAddress(pubKey.Address())

	// Encode with the correct prefix
	bech32Addr, err := bech32.ConvertAndEncode(Bech32PrefixAccAddr, addrBytes)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error encoding bech32 address: %v", err)
	}

	// If you specifically need a Covenet address type
	covenetAddr, err := covenet.CovenetAccAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error converting to Covenet address: %v", err)
	}

	// set interactor values
	interactor.pubKey = pubKey
	interactor.address = covenetAddr

	return privKey, pubKey, covenetAddr, nil
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
		// If the account is not found, set sequence to 0
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
func (interactor *CovenetInteractor) SendCovenetBlockReplicaProofTx(ctx context.Context, chainHeight uint64, blockReplica *ty.BlockReplica, resultSegment []byte, replicaURL string, txHash chan string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(proofTxTimeout))
	defer cancel()

	jsonResult, err := json.Marshal(resultSegment)
	if err != nil {
		log.Error("error in JSON marshaling result segment: ", err.Error())
		txHash <- ""

		return
	}
	sha256Result := sha256.Sum256(jsonResult)

	err = interactor.CreateProofTx(ctx, blockReplica, txHash, chainHeight, replicaURL, sha256Result)
	if err != nil {
		log.Error("error in creating proof on covenet: ", err.Error())
		txHash <- ""

		return
	}
}

func (interactor *CovenetInteractor) CreateProofTx(ctx context.Context, blockReplica *ty.BlockReplica, txHash chan string, chainHeight uint64, replicaURL string, sha256Result [sha256.Size]byte) error {

	privK, pubK, address, err := interactor.ProcessKey()
	if err != nil {
		return fmt.Errorf("failed to process key %s: %v", interactor.address.String(), err)
	}

	acSeq, acNum, err := interactor.GetAccountInfo()
	if err != nil {
		return fmt.Errorf("failed to query account %s: %v", interactor.address.String(), err)
	}

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	proofMsg := covenet.NewMsgCreateProof(interactor.address.String(), int32(blockReplica.NetworkId), "specimen", chainHeight, blockReplica.Hash.String(), hex.EncodeToString(sha256Result[:]), replicaURL)

	err = txBuilder.SetMsgs(proofMsg)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(200000)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	// Assuming we have a single private key, account number, and sequence

	log.Info("account details: ", address.String(), " number: ", acNum, " nonce: ", acSeq)

	sigV2 := signing.SignatureV2{
		PubKey: pubK,
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: acSeq,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Second round: actual signing
	signerData := xauthsigning.SignerData{
		ChainID:       "covenet",
		AccountNumber: acNum,
		Sequence:      acSeq,
	}
	sigV2, err = clientTx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, privK, encCfg.TxConfig, acSeq)
	if err != nil {
		return err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())

	// Generate a JSON string.
	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}
	txJSON := string(txJSONBytes)

	log.Info("tx JSON:\n", txJSON)

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
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
	// defer interactor.grpcClient.Close()

	if grpcRes.TxResponse.Code != 0 {
		log.Error("proof tx failed/reverted, retrying proof tx for block hash: ", blockReplica.Hash.String())
		// executeWithRetry(ctx, interactor, proofChainContract, ethClient, opts, blockReplica, txHash, chainHeight, replicaURL, sha256Result, retryCount+1)
		// return
	}

	txHash <- grpcRes.TxResponse.TxHash
	return nil
}
