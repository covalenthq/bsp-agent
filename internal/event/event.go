package event

import (
	"fmt"

	ty "github.com/covalenthq/mq-store-agent/internal/types"
)

type ReplicationSegment struct {
	BlockReplicaEvent []*BlockReplicaEvent `json:"replicaEvent"`
	StartBlock        uint64               `json:"startBlock"`
	EndBlock          uint64               `json:"endBlock"`
	Elements          uint64               `json:"elements"`
}
type BlockReplicaEvent struct {
	Hash string           `json:"hash"`
	Data *ty.BlockReplica `json:"data"`
	// Need to add Elrond specifications here I guess
}

type Event interface {
	GetHash() string
	String() string
}

func New() (Event, error) {
	return &BlockReplicaEvent{}, nil
}

func (o *BlockReplicaEvent) String() string {
	return fmt.Sprintf("hash: %s", o.Hash)
}

func (o *BlockReplicaEvent) GetHash() string {
	return o.Hash
}
