package handler

import (
	"time"

	log "github.com/sirupsen/logrus"

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
	Handle(e event.Event, hash string, datetime time.Time, data []interface{}, retry bool) error
}

type defaultHandler struct {
}

func NewDefaultHandler() Handler {
	return &defaultHandler{}
}

func (h *defaultHandler) Handle(e event.Event, hash string, datetime time.Time, data []interface{}, retry bool) error {
	log.Printf("undefined event %+v\n", e)
	return nil
}
