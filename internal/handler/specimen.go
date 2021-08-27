package handler

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/storage"
)

type specimenHandler struct {
}

func NewSpecimenHandler() Handler {
	return &specimenHandler{}
}

func (h *specimenHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	Event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	Event.Hash = hash
	Event.DateTime = datetime

	specimen := &event.SpecimenEvent{
		ReplicationEvent: Event,
	}

	specimen.Data = &data

	err := storage.HandleSpecimenUploadToBucket(*specimen, Event.Hash)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Uploaded block-specimen event: %v \nhash: %v\n", Event.ID, Event.Hash)

	return nil
}
