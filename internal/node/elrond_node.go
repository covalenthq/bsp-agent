package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type elrondAgentNode struct {
	*agentNode
}

func newElrondAgentNode(anode *agentNode) *elrondAgentNode {
	return &elrondAgentNode{anode}
}

func (node *elrondAgentNode) NodeChainType() ChainType {
	return Elrond
}

func (node *elrondAgentNode) Start(_ context.Context) {
	websocketUrls := node.AgentConfig.ChainConfig.WebsocketURLs
	if websocketUrls != "" {
		urlArr := strings.Split(websocketUrls, " ")
		for _, url := range urlArr {
			go node.consumeWebsocketsEvents(url)
		}
	}
}

// consumeWebsocketsEvents is the primary consumer of websocket events from an websocket endpoint
func (node *elrondAgentNode) consumeWebsocketsEvents(websocketURL string) {
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

			replicaURL, ccid := node.StorageManager.GenerateLocation(segmentName, message)
			log.Info("elrond binary file should be available: ", replicaURL)

			proofTxHash := make(chan string, 1)
			go node.proofchi.SendBlockReplicaProofTx(ctx, uint64(res.Block.Nonce), &types.BlockReplica{}, message, replicaURL, proofTxHash)
			pTxHash := <-proofTxHash

			if pTxHash == "" {
				log.Error("failed to prove & upload block-replica segment from websocket event: ", segmentName)
			}

			log.Info("Proof-chain tx hash: ", pTxHash, " for block-replica segment: ", segmentName)
			filename := objectFileName(segmentName, pTxHash)
			err = node.StorageManager.Store(ccid, filename, message)

			if err != nil {
				log.Error("error in storing object: ", err)
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
