package websocket

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ElrondNetwork/covalent-indexer-go/process/utility"
	"github.com/ElrondNetwork/covalent-indexer-go/schema"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	st "github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
)

func ConsumeWebsocketsEvents(config *config.EthConfig, websocketURL string, replicaCodec *goavro.Codec, ethClient *ethclient.Client, storageClient *storage.Client, binaryLocalPath, replicaBucket, proofChain string) (string, error) {
	ctx := context.Background()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urlReceiveData := url.URL{Scheme: "ws", Host: websocketURL, Path: "/block"}
	log.Printf("connecting to %s", urlReceiveData.String())
	connectionReceiveData, _, err := websocket.DefaultDialer.Dial(urlReceiveData.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connectionReceiveData.Close()

	urlAcknowledgeData := url.URL{Scheme: "ws", Host: websocketURL, Path: "/acknowledge"}
	log.Printf("connecting to %s", urlAcknowledgeData.String())
	connectionAcknowledgeData, _, err := websocket.DefaultDialer.Dial(urlAcknowledgeData.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connectionAcknowledgeData.Close()

	done := make(chan struct{})

	go func() (string, error) {
		defer close(done)
		for {
			_, message, err := connectionReceiveData.ReadMessage()
			if err != nil {
				log.Println("error read message:", err)
				return "", err
			}
			res := &schema.BlockResult{}

			errDecode := utility.Decode(res, message)
			if errDecode != nil {
				log.Println("could not decode block", errDecode)
			}
			log.Printf("received block hash: %v, nonce: %v", hex.EncodeToString(res.Block.Hash), res.Block.Nonce)
			log.Println("sending back acknowledged hash...")

			errAcknowledgeData := connectionAcknowledgeData.WriteMessage(websocket.BinaryMessage, res.Block.Hash)
			if errAcknowledgeData != nil {
				log.Println("could not send acknowledged hash :(", errAcknowledgeData)
			}
			segmentName := fmt.Sprint(res.Block.ShardID) + "-" + fmt.Sprint(res.Block.Nonce) + "-" + "segment"
			binary, _ := handler.EncodeReplicaSegmentToAvro(replicaCodec, res)
			proofTxHash := make(chan string, 1)
			go proof.SendBlockReplicaProofTx(ctx, config, proofChain, ethClient, uint64(res.Block.Nonce), 1, binary, proofTxHash)
			pTxHash := <-proofTxHash
			if pTxHash != "" {
				log.Info("Proof-chain tx hash: ", pTxHash, " for block-replica segment: ", segmentName)
				err := st.HandleObjectUploadToBucket(ctx, storageClient, binaryLocalPath, replicaBucket, segmentName, pTxHash, binary)
				if err != nil {
					return "", err
				}
			} else {
				return "", fmt.Errorf("failed to prove & upload block-replica segment event: %v", segmentName)
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return "", nil
		case <-ticker.C:
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := connectionReceiveData.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			_ = connectionAcknowledgeData.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return "", nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return "", nil
		}
	}
}
