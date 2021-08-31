package handler

import (
	"bytes"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
)

type specimenHandler struct {
}

func NewSpecimenHandler() Handler {
	return &specimenHandler{}
}

func (h *specimenHandler) Handle(config *config.Config, e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	Event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	Event.Hash = hash
	Event.DateTime = datetime

	specimen := &event.SpecimenEvent{
		ReplicationEvent: Event,
	}

	var decodedSpecimen types.BlockSpecimen
	err := rlp.Decode(bytes.NewReader(data), &decodedSpecimen)
	if err != nil {
		return fmt.Errorf("error decoding RLP bytes to block-specimen: %v", err)
	} else {
		specimen.Data = &decodedSpecimen
	}

	err = storage.HandleObjectUploadToBucket(config, string(Event.Type), Event.Hash, *specimen)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Uploaded block-specimen event: %v \nhash: %v\n", Event.ID, Event.Hash)

	return nil
}
