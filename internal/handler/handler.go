package handler

import (
	"fmt"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/event"
)

//HandlerFactory ...
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
	Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error
}

type defaultHandler struct {
}

//NewViewHandler ...
func NewDefaultHandler() Handler {
	return &defaultHandler{}
}

func (h *defaultHandler) Handle(e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {
	fmt.Printf("undefined event %+v\n", e)
	return nil
}
