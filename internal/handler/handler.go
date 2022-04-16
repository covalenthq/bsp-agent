// Package handler contains all the encoding to avro handler functions
package handler

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-cid"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/proof"
	st "github.com/covalenthq/bsp-agent/internal/storage"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	pinapi "github.com/covalenthq/ipfs-pinner"
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

// EncodeProveAndUploadReplicaSegment atomically encodes the event into an AVRO binary, proves the replica on proof-chain and upload and stores the binary file
func EncodeProveAndUploadReplicaSegment(ctx context.Context, config *config.EthConfig, pinnode pinapi.PinnerNode, replicaAvro *goavro.Codec, replicaSegment *event.ReplicationSegment, blockReplica *types.BlockReplica, gcpStorageClient *storage.Client, ethClient *ethclient.Client, binaryLocalPath, replicaBucket, segmentName, proofChain string) (string, error) {
	replicaSegmentAvro, err := EncodeReplicaSegmentToAvro(replicaAvro, replicaSegment)
	if err != nil {
		return "", err
	}
	fmt.Printf("\n---> Processing %v <---\n", segmentName)
	log.Info("Submitting block-replica segment proof for: ", segmentName)

	proofTxHash := make(chan string, 1)
	var replicaURL string
	var ccid cid.Cid
	switch {
	case gcpStorageClient != nil:
		replicaURL = "https://storage.cloud.google.com/" + replicaBucket + "/" + segmentName
	case pinnode != nil:
		ccid, err = st.GenerateCidFor(ctx, pinnode, replicaSegmentAvro)
		if err != nil {
			log.Errorf("error generating cid for %s. Error: %s", binaryLocalPath, err)
			replicaURL = "only local ./bin/"
		} else {
			replicaURL = "ipfs://" + ccid.String()
		}

	default:
		replicaURL = "only local ./bin/"
	}

	log.Info("binary file should be available: ", replicaURL)

	go proof.SendBlockReplicaProofTx(ctx, config, proofChain, ethClient, replicaSegment.EndBlock, replicaSegment.Elements, replicaSegmentAvro, replicaURL, blockReplica, proofTxHash)
	pTxHash := <-proofTxHash
	if pTxHash != "" {
		// clean this up. From the earlier switch, only one replicaURL is allowed.
		// But the object can actually be written to gcloud, local bin and ipfs.
		log.Info("Proof-chain tx hash: ", pTxHash, " for block-replica segment: ", segmentName)
		err := st.HandleObjectUploadToBucket(ctx, gcpStorageClient, binaryLocalPath, replicaBucket, segmentName, pTxHash, replicaSegmentAvro)
		_ = st.HandleObjectUploadToIPFS(ctx, pinnode, ccid, binaryLocalPath, segmentName, pTxHash)

		if err != nil {
			return "", fmt.Errorf("error in uploading object to bucket: %w", err)
		}
	} else {
		return "", fmt.Errorf("failed to prove & upload block-replica segment event: %v", segmentName)
	}

	return pTxHash, nil
}
