package handler

import (
	"bytes"
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	st "github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(e event.Event, hash string, data []byte) (*event.SpecimenEvent, *event.ResultEvent, error) {

	replEvent, ok := e.(*event.ReplicationMsg)
	if !ok {
		return nil, nil, fmt.Errorf("incorrect event type: %v", replEvent.Type)
	}

	replEvent.Hash = hash

	result := &event.ResultEvent{
		Msg: replEvent,
	}

	var decodedResult types.BlockResult
	err := rlp.Decode(bytes.NewReader(data), &decodedResult)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RLP bytes to block-result: %w", err)
	} else {
		result.Data = &decodedResult
	}

	return nil, result, nil
}

func encodeResultSegmentToAvro(resultAvro *goavro.Codec, blockResultSegment interface{}) ([]byte, error) {

	resultMap, err := utils.StructToMap(blockResultSegment)
	if err != nil {
		return nil, err
	}

	// Convert native Go map[string]interface{} to binary Avro data
	binaryResultSegment, err := resultAvro.BinaryFromNative(nil, resultMap)
	if err != nil {
		log.Fatalf("failed to convert Go map to Avro binary data: %v", err)
	}

	return binaryResultSegment, nil
}

func EncodeProveAndUploadResultSegment(ctx context.Context, config *config.EthConfig, resultAvro *goavro.Codec, resultSegment *event.ReplicationSegment, resultBucket, segmentName string, storage *storage.Client, ethClient *ethclient.Client, proofChain string) (string, error) {

	resultSegmentAvro, err := encodeResultSegmentToAvro(resultAvro, resultSegment)
	if err != nil {
		return "", err
	}

	log.Info("Submitting block-result segment proof for: ", segmentName)

	proofTxHash := make(chan string, 1)

	go proof.SendBlockResultProofTx(ctx, config, proofChain, ethClient, resultSegment.EndBlock, resultSegment.Elements, resultSegmentAvro, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {
		err := st.HandleObjectUploadToBucket(ctx, storage, "block-result", resultBucket, segmentName, resultSegmentAvro)
		if err != nil {
			return "", err
		}
		log.Info("Uploaded block-result segment event: ", segmentName, " with proof tx hash: ", pTxHash)
	} else {
		return "", fmt.Errorf("failed to prove & upload block-result segment event: %v", segmentName)
	}

	return pTxHash, nil
}
