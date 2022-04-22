// Package handler contains all the encoding to avro handler functions
package handler

import (
	"fmt"

	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
)

// EncodeReplicaSegmentToAvro encodes replica segment into AVRO binary encoding
func EncodeReplicaSegmentToAvro(replicaAvro *goavro.Codec, blockReplicaSegment interface{}) ([]byte, error) {
	replicaMap, err := utils.StructToMap(blockReplicaSegment)
	if err != nil {
		return nil, fmt.Errorf("error in converting struct to map: %w", err)
	}
	// Convert native Go map[string]interface{} to binary Avro data
	binaryReplicaSegment, err := replicaAvro.BinaryFromNative(nil, replicaMap)
	if err != nil {
		log.Error("failed to convert Go map to Avro binary data: ", err)
	}

	return binaryReplicaSegment, nil
}

// ParseStreamToEvent takes the stream message and parses it to a block replica event
func ParseStreamToEvent(e event.Event, hash string, data *types.BlockReplica) (*event.BlockReplicaEvent, error) {
	replEvent, ok := e.(*event.BlockReplicaEvent)
	if !ok {
		return nil, fmt.Errorf("incorrect event type: %v", replEvent)
	}
	replicaEvent := &event.BlockReplicaEvent{
		Data: data,
		Hash: hash,
	}

	return replicaEvent, nil
}
