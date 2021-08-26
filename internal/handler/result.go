package handler

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/storage"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	event.Hash = hash
	event.Data = data
	event.DateTime = datetime

	var s big.Int
	err := rlp.Decode(bytes.NewReader(data), &s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Decoded RLP value: %v\n", s)
	}

	err = storage.HandleResultUploadToBucket(*event, event.Hash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Uploaded block-result event: %v \nhash: %v\n", event.ID, event.Hash)

	return nil
}
