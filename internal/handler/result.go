package handler

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/storage"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(e event.Event, hash string, datetime time.Time, data []interface{}, retry bool) error {
	event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	event.Hash = hash
	event.Data = data
	event.DateTime = datetime

	err := storage.HandleResultUploadToBucket(*event, event.Hash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Uploaded block-result event: %v \nhash: %v\n", event.ID, event.Hash)

	return nil
}
