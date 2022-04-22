package node

import (
	"bytes"

	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/handler"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/go-redis/redis/v7"
	"github.com/golang/snappy"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"
)

func ReplicaFromRedisMessage(msg redis.XMessage) (*event.BlockReplicaEvent, string, error) {
	hash := msg.Values["hash"].(string)
	decodedData, err := snappy.Decode(nil, []byte(msg.Values["data"].(string)))
	if err != nil {
		log.Info("Failed to snappy decode: ", err.Error())
	}

	var blockReplica types.BlockReplica
	err = rlp.Decode(bytes.NewReader(decodedData), &blockReplica)
	if err != nil {
		log.Error("error decoding RLP bytes to block-replica: ", err)
	}

	newEvent, _ := event.NewBlockReplicaEvent()
	replica, err := handler.ParseStreamToEvent(newEvent, hash, &blockReplica)
	objectType := replica.Type()
	return replica, objectType, err
}
