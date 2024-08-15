package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	// "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"google.golang.org/grpc"
)

func sendProofTx() error {
	// Choose your codec: Amino or Protobuf. Here, we use Protobuf, given by the
	// following function.
	encCfg := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	// create test keys
	priv1, _, addr1 := testdata.KeyTestPubAddr()
	priv2, _, addr2 := testdata.KeyTestPubAddr()
	_, _, addr3 := testdata.KeyTestPubAddr()

	fmt.Println(priv1, priv2, "Test Keys")

	fmt.Println(addr1, addr2, " 1 & 2 Addresses")

	// Define two x/bank MsgSend messages:
	// - from addr1 to addr3,
	// - from addr2 to addr3.
	// This means that the transactions needs two signers: addr1 and addr2.
	msg1 := banktypes.NewMsgSend(addr1, addr3, types.NewCoins(types.NewInt64Coin("stake", 12)))
	msg2 := banktypes.NewMsgSend(addr2, addr3, types.NewCoins(types.NewInt64Coin("stake", 34)))

	err := txBuilder.SetMsgs(msg1, msg2)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(200000)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	privs := []cryptotypes.PrivKey{priv1, priv2}
	accNums := []uint64{0, 1, 2} // The accounts' account numbers
	accSeqs := []uint64{0, 1, 2} // The accounts' sequence numbers

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

	fmt.Println(txJSON, "txJson")

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

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

	fmt.Println(grpcRes.TxResponse.Code, "response code") // Should be `0` if the tx is successful

	return nil
}

// func simulateTx() error {
//     // --snip--

//     // Simulate the tx via gRPC. We create a new client for the Protobuf Tx
//     // service.
//     txClient := tx.NewServiceClient(grpcConn)
//     txBytes := /* Fill in with your signed transaction bytes. */

//     // We then call the Simulate method on this client.
//     grpcRes, err := txClient.Simulate(
//         context.Background(),
//         &tx.SimulateRequest{
//             TxBytes: txBytes,
//         },
//     )
//     if err != nil {
//         return err
//     }

//     fmt.Println(grpcRes.GasInfo) // Prints estimated gas used.

//     return nil
// }

func main() {

	// Assuming you have a private key as a hex string
	privKeyHex := "96891492094e9e3aea0dd33cb6e065dee8331f5736fa93b5629edb4b2073631a"

	// Decode the hex string to bytes
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		fmt.Printf("Error decoding private key: %v\n", err)
		return
	}

	// Create a new PrivKey object
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}
	// Get the public key from the private key
	pubKey := privKey.PubKey().(*secp256k1.PubKey)

	// Convert to the API type
	apiPrivKey := &secp256k1.PrivKey{
		Key: privKey.Key,
	}
	apiPubKey := &secp256k1.PubKey{
		Key: pubKey.Key,
	}

	// Get the address from the public key
	addr := types.AccAddress(pubKey.Address())

	fmt.Printf("Private Key: %x\n", apiPrivKey.Key)
	fmt.Printf("Public Key: %x\n", apiPubKey.Key)
	fmt.Printf("Address: %s\n", addr)

	sendProofTx()

}
