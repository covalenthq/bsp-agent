// Package handler contains fns for encoding redis stream messages to block-replica AVRO segments
package handler

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-redis/redis/v7"
	"github.com/golang/snappy"
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
	// Convert native <nil> to goavro Union type expected by codec
	replicaMod := utils.MapToAvroUnion(replicaMap)
	// Convert native Go map[string]interface{} to binary Avro data
	binaryReplicaSegment, err := replicaAvro.BinaryFromNative(nil, replicaMod)
	if err != nil {
		log.Error("failed to convert Go map to Avro binary data: ", err)
	}

	return binaryReplicaSegment, nil
}

// ParseEventToBlockReplica takes block-replica data and parses it to a block-replica event
func ParseEventToBlockReplica(e event.Event, hash string, data *types.BlockReplica) (*event.BlockReplicaEvent, error) {
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

// ParseMessageToBlockReplica decodes the redis message to a BlockReplicaEvent
func ParseMessageToBlockReplica(msg redis.XMessage) (*event.BlockReplicaEvent, error) {
	hash := msg.Values["hash"].(string)
	decodedData, err := snappy.Decode(nil, []byte(msg.Values["data"].(string)))
	if err != nil {
		log.Info("Failed to snappy decode: ", err.Error())

		return nil, fmt.Errorf("%w", err)
	}

	var blockReplica types.BlockReplica

	err = rlp.DecodeBytes(decodedData, &blockReplica)
	if err != nil {
		log.Error("error decoding RLP bytes to block-replica: ", err)

		return nil, fmt.Errorf("%w", err)
	}

	newEvent, _ := event.NewBlockReplicaEvent()
	replica, err := ParseEventToBlockReplica(newEvent, hash, &blockReplica)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return replica, nil
}
