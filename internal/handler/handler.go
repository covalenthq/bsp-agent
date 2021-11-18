package handler

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	st "github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

func Parse(e event.Event, hash string, data *types.BlockReplica) (*event.BlockReplicaEvent, error) {
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

func encodeReplicaSegmentToAvro(replicaAvro *goavro.Codec, blockReplicaSegment interface{}) ([]byte, error) {
	replicaMap, err := utils.StructToMap(blockReplicaSegment)
	if err != nil {
		return nil, err
	}
	// Convert native Go map[string]interface{} to binary Avro data
	binaryReplicaSegment, err := replicaAvro.BinaryFromNative(nil, replicaMap)
	if err != nil {
		log.Fatalf("failed to convert Go map to Avro binary data: %v", err)
	}

	return binaryReplicaSegment, nil
}

func EncodeProveAndUploadReplicaSegment(ctx context.Context, config *config.EthConfig, replicaAvro *goavro.Codec, replicaSegment *event.ReplicationSegment, replicaBucket, segmentName string, storage *storage.Client, ethClient *ethclient.Client, proofChain string) (string, error) {
	replicaSegmentAvro, err := encodeReplicaSegmentToAvro(replicaAvro, replicaSegment)
	if err != nil {
		return "", err
	}
	log.Info("Submitting block-replica segment proof for: ", segmentName)

	proofTxHash := make(chan string, 1)
	go proof.SendBlockReplicaProofTx(ctx, config, proofChain, ethClient, replicaSegment.EndBlock, replicaSegment.Elements, replicaSegmentAvro, proofTxHash)
	pTxHash := <-proofTxHash
	if pTxHash != "" {
		err := st.HandleObjectUploadToBucket(ctx, storage, replicaBucket, segmentName, replicaSegmentAvro)
		if err != nil {
			return "", err
		}
		log.Info("Uploaded block-replica segment event: ", segmentName, " with proof tx hash: ", pTxHash)
	} else {
		return "", fmt.Errorf("failed to prove & upload block-replica segment event: %v", segmentName)
	}

	return pTxHash, nil
}
