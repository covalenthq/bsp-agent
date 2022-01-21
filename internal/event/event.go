// Package event contains all the events that relate to a block replica
package event

import (
	"fmt"

	ty "github.com/covalenthq/mq-store-agent/internal/types"
)

// ReplicationSegment is block replication segment that is converted to AVRO encoding
type ReplicationSegment struct {
	BlockReplicaEvent []*BlockReplicaEvent `json:"replicaEvent"`
	StartBlock        uint64               `json:"startBlock"`
	EndBlock          uint64               `json:"endBlock"`
	Elements          uint64               `json:"elements"`
}

// BlockReplicaEvent is a replication event coming from redis stream
type BlockReplicaEvent struct {
	Hash string           `json:"hash"`
	Data *ty.BlockReplica `json:"data"`
}

// Event allows for accessing the hash of the object
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
