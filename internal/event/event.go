package event

import (
	"fmt"
	"time"
)

type Type string

const (
	SpecimenType Type = "block-specimen"
	ResultType   Type = "block-result"
)

type ReplicationEvent struct {
	*Base
}

type Base struct {
	ID       string        `json:"ID"`
	Type     Type          `json:"type"`
	Hash     string        `json:"hash"`
	DateTime time.Time     `json:"datetime"`
	Data     []interface{} `json:"data"`
}
type Event interface {
	GetID() string
	GetType() Type
	GetDateTime() time.Time
	GetData() []interface{}
	GetHash() string
	SetID(id string)
}

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}

	switch t {

	case SpecimenType:
		return &ReplicationEvent{
			Base: b,
		}, nil

	case ResultType:
		return &ReplicationEvent{
			Base: b,
		}, nil

	}

	return nil, fmt.Errorf("type %v not supported", t)
}

func (o *Base) GetID() string {
	return o.ID
}

func (o *Base) SetID(id string) {
	o.ID = id
}

func (o *Base) GetType() Type {
	return o.Type
}

func (o *Base) GetDateTime() time.Time {
	return o.DateTime
}

func (o *Base) String() string {
	return fmt.Sprintf("id: %s type: %s hash: %s", o.ID, o.Type, o.Hash)
}

func (o *Base) GetData() []interface{} {
	return o.Data
}

func (o *Base) GetHash() string {
	return o.Hash
}
