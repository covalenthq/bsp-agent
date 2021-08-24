package handler

import (
	"fmt"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/gcp"
)

type resultHandler struct {
}

//NewResultHandler ...
func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	event, ok := e.(*event.ResultEvent)

	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	event.Hash = hash
	event.Data = data
	event.DateTime = datetime

	gcp.HandleResultUploadToBucket(*event, event.Hash)

	fmt.Printf("completed uploading block-result event %v hash: %v\n", event.ID, event.Hash)

	return nil
}
