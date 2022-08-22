// Package event contains events/interfaces for raw and processsed block-replica objects
package event

import (
	"fmt"

	ty "github.com/covalenthq/bsp-agent/internal/types"
)

// ReplicaSegmentWrapped wraps a ReplicationSegment with additional information
type ReplicaSegmentWrapped struct {
	ReplicationSegment
	IDBatch     []string
	SkipIDBatch []string
	SegmentName string
}

// ReplicationSegment is block replication segment that is converted to AVRO encoding
type ReplicationSegment struct {
	BlockReplicaEvent []*BlockReplicaEvent `json:"replicaEvent"`
	StartBlock        uint64               `json:"startBlock"`
	EndBlock          uint64               `json:"endBlock"`
	Elements          uint64               `json:"elements"`
}

// BlockReplicaEvent is a replication event coming from a redis stream
type BlockReplicaEvent struct {
	Hash string           `json:"hash"`
	Data *ty.BlockReplica `json:"data"`
}

// Event allows for accessing the hash of block-replica event object
type Event interface {
	GetBlockReplicaHash() string
	GetBlockReplicaString() string
}

// NewBlockReplicaEvent creates a new block-replica event
func NewBlockReplicaEvent() (Event, error) {
	return &BlockReplicaEvent{}, nil
}

// GetBlockReplicaString gets the block-replica string
func (o *BlockReplicaEvent) GetBlockReplicaString() string {
	return fmt.Sprintf("hash: %s", o.Hash)
}

// GetBlockReplicaHash gets the block-replica hash
func (o *BlockReplicaEvent) GetBlockReplicaHash() string {
	return o.Hash
}

// Type gets the block-replica event types - block-specimen, block-result, or combined(block-replica)
func (o *BlockReplicaEvent) Type() string {
	// would return -replica, -specimen or -result
	return o.Data.Type[5:]
}
