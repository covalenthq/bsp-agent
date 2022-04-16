// Package websocket provides a websocket interface as a data ingestion mechanism
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
	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/proof"
	st "github.com/covalenthq/bsp-agent/internal/storage"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
)

// ConsumeWebsocketsEvents is the primary consumer of websocket events from an websocket endpoint
func ConsumeWebsocketsEvents(config *config.EthConfig, websocketURL string, replicaCodec *goavro.Codec, ethClient *ethclient.Client, gcpStorageClient *storage.Client, binaryLocalPath, replicaBucket, proofChain string) {
	var replicaURL string
	ctx := context.Background()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urlReceiveData := url.URL{Scheme: "ws", Host: websocketURL, Path: "/block"}
	log.Info("Connecting to websocket: ", urlReceiveData.String())
	connectionReceiveData, _, err := websocket.DefaultDialer.Dial(urlReceiveData.String(), nil)
	if err != nil {
		log.Error("error in connecting to websocket /block dial: ", err)
	}
	defer func() {
		if cerr := connectionReceiveData.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	urlAcknowledgeData := url.URL{Scheme: "ws", Host: websocketURL, Path: "/acknowledge"}
	log.Info("Connecting to websocket: ", urlAcknowledgeData.String())
	connectionAcknowledgeData, _, err := websocket.DefaultDialer.Dial(urlAcknowledgeData.String(), nil)
	if err != nil {
		log.Error("error in connecting to websocket /ack dial: ", err)
	}
	defer func() {
		if cerr := connectionAcknowledgeData.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := connectionReceiveData.ReadMessage()
			if err != nil {
				log.Error("error in websocket received message: ", err)
			}
			res := &types.ElrondBlockReplica{}
			errDecode := utils.DecodeAvro(res, message)
			if errDecode != nil {
				log.Error("error in decoding block from avro: ", errDecode)
			}

			segmentName := fmt.Sprint(res.Block.ShardID) + "-" + fmt.Sprint(res.Block.Nonce) + "-" + "segment"
			fmt.Printf("\n---> Processing %v <---\n", segmentName)
			log.Info("Sending acknowledged hash: ", hex.EncodeToString(res.Block.Hash), " nonce: ", res.Block.Nonce)

			errAcknowledgeData := connectionAcknowledgeData.WriteMessage(websocket.BinaryMessage, res.Block.Hash)
			if errAcknowledgeData != nil {
				log.Error("error in sending acknowledged hash: ", errAcknowledgeData)
			}

			proofTxHash := make(chan string, 1)
			// Only google storage is supported for now
			if gcpStorageClient != nil {
				replicaURL = "https://storage.cloud.google.com/" + replicaBucket + "/" + segmentName
			} else {
				replicaURL = "only local ./bin/"
			}
			go proof.SendBlockReplicaProofTx(ctx, config, proofChain, ethClient, uint64(res.Block.Nonce), 1, message, replicaURL, &types.BlockReplica{}, proofTxHash)
			pTxHash := <-proofTxHash
			if pTxHash != "" {
				log.Info("Proof-chain tx hash: ", pTxHash, " for block-replica segment: ", segmentName)
				err := st.HandleObjectUploadToBucket(ctx, gcpStorageClient, binaryLocalPath, replicaBucket, segmentName, pTxHash, message)
				if err != nil {
					log.Error("error in handling object upload and storage: ", err)
				}
			} else {
				log.Error("failed to prove & upload block-replica segment from websocket event: ", segmentName)
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:

			return
		case <-ticker.C:
		case <-interrupt:
			log.Info("interrupt recevied")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := connectionReceiveData.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			_ = connectionAcknowledgeData.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Error("error in closing websocket message: ", err)

				return
			}

			return
		}
	}
}
