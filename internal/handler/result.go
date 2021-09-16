package handler

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	"github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(config *config.Config, e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {

	ctx := context.Background()
	Event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	Event.Hash = hash
	Event.DateTime = datetime

	result := &event.ResultEvent{
		ReplicationEvent: Event,
	}

	var decodedResult types.BlockResult
	err := rlp.Decode(bytes.NewReader(data), &decodedResult)
	if err != nil {
		return fmt.Errorf("error decoding RLP bytes to block-result: %v", err)
	} else {
		result.Data = &decodedResult
	}

	err = storage.HandleObjectUploadToBucket(config, string(Event.Type), Event.Hash, *result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Uploaded block-result event: %v \nhash: %v\n", Event.ID, Event.Hash)

	ethClient := proof.GetEthClient(config)
	blockHash := common.HexToHash(result.ReplicationEvent.Hash)

	block, err := ethClient.BlockByHash(ctx, blockHash)
	if err != nil {
		log.Fatal(err)
	}

	txHash, mined, err := proof.SubmitResultProofTx(config, ethClient, block.NumberU64(), 1, *result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Block-result proof hash: %v \nsubmitted: %v\n", txHash, mined)

	return nil
}
