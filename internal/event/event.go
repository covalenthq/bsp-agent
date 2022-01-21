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
	GetBlockReplicaHash() string
	GetBlockReplicaString() string
}

// NewBlockReplicaEvent creates a new block replica event
func NewBlockReplicaEvent() (Event, error) {
	return &BlockReplicaEvent{}, nil
}

// GetBlockReplicaString gets the block replica string
func (o *BlockReplicaEvent) GetBlockReplicaString() string {
	return fmt.Sprintf("hash: %s", o.Hash)
}

// GetBlockReplicaHash gets the block replica hash
func (o *BlockReplicaEvent) GetBlockReplicaHash() string {
	return o.Hash
}
