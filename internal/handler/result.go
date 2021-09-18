package handler

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	st "github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {

	ctx := context.Background()
	replEvent, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	replEvent.Hash = hash
	replEvent.DateTime = datetime

	result := &event.ResultEvent{
		ReplicationEvent: replEvent,
	}

	var decodedResult types.BlockResult
	err := rlp.Decode(bytes.NewReader(data), &decodedResult)
	if err != nil {
		return fmt.Errorf("error decoding RLP bytes to block-result: %v", err)
	} else {
		result.Data = &decodedResult
	}

	blockHash := common.HexToHash(result.ReplicationEvent.Hash)
	block, err := ethSource.HeaderByHash(ctx, blockHash)

	if err != nil {
		log.Error("error in getting block: ", err.Error())
	}

	log.Info("Submitting block-result proof for: ", block.Number.Uint64())

	proofTxHash := make(chan string, 1)

	go proof.SendBlockResultProofTx(ctx, config, ethProof, block.Number.Uint64(), *result, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {

		err = st.HandleObjectUploadToBucket(ctx, config, storage, string(replEvent.Type), replEvent.Hash, *result)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Uploaded block-result event: ", replEvent.Hash, " with proof tx hash: ", pTxHash)

	} else {
		log.Fatal(err)
	}

	return nil
}
