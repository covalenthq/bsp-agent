package event

import (
	"encoding"
	"fmt"
	"time"
)

type Type string

const (
	SpecimenType Type = "block-specimen"
	ResultType   Type = "block-result"
)

type Base struct {
	ID       string
	Type     Type      `json:"type"`
	Hash     string    `json:"hash"`
	DateTime time.Time `json:"datetime"`
	Data     []byte    `json:"data"`
}

// Event ...
type Event interface {
	GetID() string
	GetType() Type
	GetDateTime() time.Time
	GetData() []byte
	GetHash() string
	SetID(id string)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}

	switch t {

	case SpecimenType:
		return &SpecimenEvent{
			Base: b,
		}, nil

	case ResultType:
		return &ResultEvent{
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

	return fmt.Sprintf("id:%s type:%s hash:%s", o.ID, o.Type, o.Hash)
}

func (o *Base) GetData() []byte {

	return o.Data
}

func (o *Base) GetHash() string {

	return o.Hash
}
