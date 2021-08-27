package handler

import (
	"bytes"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
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
		return fmt.Errorf("error decoding RLP bytes to block result: %v", err)
	} else {
		result.Data = &decodedResult
	}

	err = storage.HandleResultUploadToBucket(*result, Event.Hash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Uploaded block-result event: %v \nhash: %v\n", Event.ID, Event.Hash)

	return nil
}
