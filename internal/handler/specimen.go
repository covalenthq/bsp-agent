package handler

import (
	"fmt"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/event"
)

type specimenHandler struct {
}

//NewSpecimenHandler ...
func NewSpecimenHandler() Handler {
	return &specimenHandler{}
}

func (h *specimenHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	event, ok := e.(*event.SpecimenEvent)

	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	event.Hash = hash
	event.Data = data
	event.DateTime = datetime

	// fmt.Printf("processed event %+v UserID: %v Comment:%v \n", event, event.UserID, event.Comment)

	fmt.Printf("completed block-specimen event %v hash: %v\n", event.ID, event.Hash)

	return nil
}
