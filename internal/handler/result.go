package handler

import (
	"fmt"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/gcp"
	log "github.com/sirupsen/logrus"
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

	err := gcp.HandleResultUploadToBucket(*event, event.Hash)
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("completed uploading block-result event %v hash: %v\n", event.ID, event.Hash)

	return nil
}
