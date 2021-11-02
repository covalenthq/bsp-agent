package event

import (
	"fmt"

	ty "github.com/covalenthq/mq-store-agent/internal/types"
)

type Type string

const (
	SpecimenType Type = "block-specimen"
	ResultType   Type = "block-result"
)

type ReplicationSegment struct {
	BlockReplicationData *BlockReplicationData `json:"ReplicationEvent"`
	StartBlock           uint64                `json:"StartBlock"`
	EndBlock             uint64                `json:"EndBlock"`
	Elements             uint64                `json:"Elements"`
}

type BlockReplicationData struct {
	Result   []*ResultEvent   `json:"result,omitempty"`
	Specimen []*SpecimenEvent `json:"specimen,omitempty"`
}

type ResultEvent struct {
	Data *ty.BlockResult `json:"data"`
	Msg  *ReplicationMsg `json:"msg"`
}

type SpecimenEvent struct {
	Data *ty.BlockSpecimen `json:"data"`
	Msg  *ReplicationMsg   `json:"msg"`
}

type ReplicationMsg struct {
	*Base
}

type Base struct {
	ID   string `json:"ID"`
	Type Type   `json:"type"`
	Hash string `json:"hash"`
}
type Event interface {
	GetID() string
	SetID(id string)
	GetType() Type
	GetHash() string
	String() string
}

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}

	switch t {
	case SpecimenType:
		return &ReplicationMsg{
			Base: b,
		}, nil
	case ResultType:
		return &ReplicationMsg{
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

func (o *Base) String() string {
	return fmt.Sprintf("id: %s type: %s hash: %s", o.ID, o.Type, o.Hash)
}

func (o *Base) GetHash() string {
	return o.Hash
}
