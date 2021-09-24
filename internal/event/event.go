package event

import (
	"fmt"
	"time"

	ty "github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/ethereum/go-ethereum/core/types"
)

type Type string

const (
	SpecimenType Type = "block-specimen"
	ResultType   Type = "block-result"
)

type ResultSegment struct {
	BlockResult []*ResultEvent `json:"BlockResult"`
	StartBlock  uint64         `json:"StartBlock"`
	EndBlock    uint64         `json:"EndBlock"`
	Elements    uint64         `json:"Elements"`
}

type SpecimenSegment struct {
	BlockSpecimen []*SpecimenEvent `json:"BlockResult"`
	StartBlock    uint64           `json:"StartBlock"`
	EndBlock      uint64           `json:"EndBlock"`
	Elements      uint64           `json:"Elements"`
}

type ResultEvent struct {
	ReplicationEvent *ReplicationEvent `json:"ReplicationEvent"`
	Data             *ty.BlockResult   `json:"result"`
}

type SpecimenEvent struct {
	ReplicationEvent *ReplicationEvent `json:"ReplicationEvent"`
	Data             *ty.BlockSpecimen `json:"specimen"`
	BlockHeader      *types.Header     `json:"header"`
}

type ReplicationEvent struct {
	*Base
}

type Base struct {
	ID       string    `json:"ID"`
	Type     Type      `json:"type"`
	Hash     string    `json:"hash"`
	DateTime time.Time `json:"datetime"`
}
type Event interface {
	GetID() string
	SetID(id string)
	GetType() Type
	GetHash() string
	GetDateTime() time.Time
	String() string
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

func (o *Base) GetHash() string {
	return o.Hash
}
