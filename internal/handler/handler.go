package handler

import (
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
)

func HandlerFactory() func(t event.Type) Handler {
	return func(t event.Type) Handler {
		switch t {
		case event.SpecimenType:
			return NewSpecimenHandler()
		case event.ResultType:
			return NewResultHandler()
		default:
			return NewDefaultHandler()
		}
	}
}

type Handler interface {
	Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) (*event.SpecimenEvent, *event.ResultEvent, error)
}

type defaultHandler struct {
}

func NewDefaultHandler() Handler {
	return &defaultHandler{}
}

func (h *defaultHandler) Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) (*event.SpecimenEvent, *event.ResultEvent, error) {
	log.Printf("undefined event %+v\n", e)
	return &event.SpecimenEvent{}, &event.ResultEvent{}, nil
}
