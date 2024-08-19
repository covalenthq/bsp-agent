package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

// Add this at the package level or in an init() function
var encCfg params.EncodingConfig

func init() {
	encCfg = simapp.MakeTestEncodingConfig()
}

// simd accounts
func main() {

	privKeyHexes := []string{
		"96891492094e9e3aea0dd33cb6e065dee8331f5736fa93b5629edb4b2073631a", //alice
		"ae87f7015182f1fe893c9b05e6c9412cb7c5531cb3451c033a13da4e5d3b6600", //bob
		"1d77b70a31797a9c11e73f50acfde9422608f1e37de63a8037ad5c6f1c9fac0b", //carol
		"8a81accad0711457e225115492e6d490ea4767e49fcb2df10640f1caabf1b06d", //david
	}

	// Get public, private and account addresses from hex keys
	privateKeys, publicKeys, addresses, err := processPrivateKeys(privKeyHexes)
	if err != nil {
		fmt.Printf("Error processing private keys: %v\n", err)
		return
	}

	// Get the GRPC node connection
	grpcConn, err := getGRPCConnection()
	if err != nil {
		fmt.Printf("Error processing grpc connection: %v\n", err)
		return
	}

	// Get account information like nonce and number
	sequences, numbers, err := queryAccountInfo(grpcConn, addresses)
	if err != nil {
		fmt.Printf("Error processing sequence queries: %v\n", err)
		return
	}

	fmt.Println("GRPC connection status:", grpcConn.GetState())

	for i := 0; i < 4; i++ {
		fmt.Printf("Key Set %d:\n", i+1)
		fmt.Printf("  Private Key: %x\n", privateKeys[i].Bytes())
		fmt.Printf("  Public Key: %x\n", publicKeys[i].Bytes())
		fmt.Printf("  Address: %s\n", addresses[i].String())
		fmt.Printf("  Sequence: %d\n", sequences[i])
		fmt.Printf("  Account Number: %d\n", numbers[i])
		fmt.Println()
	}

	// Send the proof transaction with account data
	sendProofTx(privateKeys, publicKeys, addresses, sequences, numbers)

	// Allow time for tx mine
	time.Sleep(5 * time.Second)

	// Get tx balances
	getStakeBalances(grpcConn, addresses)

	fmt.Println("process completed")

	// Close connection
	defer grpcConn.Close()
}

func getStakeBalances(grpcConn *grpc.ClientConn, addresses []sdk.AccAddress) error {
	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)

	for i, addr := range addresses {
		bankRes, err := bankClient.Balance(
			context.Background(),
			&banktypes.QueryBalanceRequest{Address: addr.String(), Denom: "stake"},
		)
		if err != nil {
			return fmt.Errorf("error querying balance for address %d: %v", i, err)
		}

		fmt.Printf("Address %d balance: %s\n", i, bankRes.GetBalance())
	}

	return nil
}

func getGRPCConnection() (*grpc.ClientConn, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"localhost:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}

	return grpcConn, nil
}

func processPrivateKeys(privKeyHexes []string) ([]cryptotypes.PrivKey, []cryptotypes.PubKey, []sdk.AccAddress, error) {
	if len(privKeyHexes) != 4 {
		return nil, nil, nil, fmt.Errorf("expected 4 private keys, got %d", len(privKeyHexes))
	}

	var privateKeys []cryptotypes.PrivKey
	var publicKeys []cryptotypes.PubKey
	var addresses []sdk.AccAddress

	for _, privKeyHex := range privKeyHexes {
		// Decode the hex string to bytes
		privKeyBytes, err := hex.DecodeString(privKeyHex)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error decoding private key: %v", err)
		}

		// Create a new PrivKey object
		privKey := &secp256k1.PrivKey{Key: privKeyBytes}

		// Get the public key from the private key
		pubKey := privKey.PubKey()

		// Generate Cosmos address
		addr := sdk.AccAddress(pubKey.Address())

		privateKeys = append(privateKeys, privKey)
		publicKeys = append(publicKeys, pubKey)
		addresses = append(addresses, addr)
	}

	return privateKeys, publicKeys, addresses, nil
}

func queryAccountInfo(grpcConn *grpc.ClientConn, addresses []sdk.AccAddress) ([]uint64, []uint64, error) {
	// Create a new auth query client
	authClient := authtypes.NewQueryClient(grpcConn)
	sequences := make([]uint64, len(addresses))
	numbers := make([]uint64, len(addresses))

	for i, addr := range addresses {
		// Query the account
		res, err := authClient.Account(
			context.Background(),
			&authtypes.QueryAccountRequest{Address: addr.String()},
		)
		if err != nil {
			// If the account is not found, set sequence to 0
			if err.Error() == "rpc error: code = NotFound desc = account cosmos1... not found: key not found" {
				sequences[i] = 0
				continue
			}
			return nil, nil, fmt.Errorf("failed to query account %s: %v", addr.String(), err)
		}

		// Create a new AccountI interface
		var account authtypes.AccountI

		// Unmarshal the account data
		err = encCfg.InterfaceRegistry.UnpackAny(res.Account, &account)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to unpack account %s: %v", addr.String(), err)
		}

		// Get the sequence number
		sequences[i] = account.GetSequence()
		numbers[i] = account.GetAccountNumber()
	}

	return sequences, numbers, nil
}

func sendProofTx(privateKeys []cryptotypes.PrivKey, publicKeys []cryptotypes.PubKey, addresses []sdk.AccAddress, sequences []uint64, numbers []uint64) error {
	// Choose your codec: Amino or Protobuf. Here, we use Protobuf, given by the
	// following function.
	encCfg := simapp.MakeTestEncodingConfig()
	_ = publicKeys

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	// Define two x/bank MsgSend messages:
	// - from addr1 to addr3,
	// - from addr2 to addr3.
	// This means that the transactions needs two signers: addr1 and addr2.
	msg1 := banktypes.NewMsgSend(addresses[0], addresses[1], types.NewCoins(types.NewInt64Coin("stake", 100))) //where to send the funds
	// msg2 := banktypes.NewMsgSend(addresses[1], addresses[3], types.NewCoins(types.NewInt64Coin("stake", 10)))

	err := txBuilder.SetMsgs(msg1)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(200000)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	privs := []cryptotypes.PrivKey{privateKeys[0]}
	accNums := numbers
	accSeqs := sequences

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       "learning-chain-1",
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := clientTx.SignWithPrivKey(
			encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
			txBuilder, priv, encCfg.TxConfig, accSeqs[i])
		if err != nil {
			return err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = txBuilder.SetSignatures(sigsV2...)
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

	fmt.Println("tx JSON:\n", txJSON)

	grpcConn, err := getGRPCConnection()
	if err != nil {
		return err
	}
	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := tx.NewServiceClient(grpcConn)
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

	fmt.Println("response code\n", grpcRes.TxResponse.String()) // Should be `0` if the tx is successful

	return nil
}
