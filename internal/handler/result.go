package handler

import (
	"fmt"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/event"
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

	// fmt.Printf("processed event %+v UserID: %v Comment:%v \n", event, event.UserID, event.Comment)

	fmt.Printf("completed block-result event %v hash: %v\n", event.ID, event.Hash)

	return nil
}
