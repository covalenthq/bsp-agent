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

type specimenHandler struct {
}

func NewSpecimenHandler() Handler {
	return &specimenHandler{}
}

func (h *specimenHandler) Handle(e event.Event, hash string, data []byte) (*event.SpecimenEvent, *event.ResultEvent, error) {

	replEvent, ok := e.(*event.ReplicationMsg)
	if !ok {
		return nil, nil, fmt.Errorf("incorrect event type: %v", replEvent.Type)
	}

	replEvent.Hash = hash

	specimen := &event.SpecimenEvent{
		Msg: replEvent,
	}

	var decodedSpecimen types.BlockSpecimen
	err := rlp.Decode(bytes.NewReader(data), &decodedSpecimen)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RLP bytes to block-specimen: %w", err)
	} else {
		specimen.Data = &decodedSpecimen
	}

	return specimen, nil, nil
}

func encodeSpecimenSegmentToAvro(specimenAvro *goavro.Codec, blockSpecimenSegment interface{}) ([]byte, error) {

	specimenMap, err := utils.StructToMap(blockSpecimenSegment)
	if err != nil {
		return nil, err
	}

	// Convert native Go form to binary Avro data
	binarySpecimenSegment, err := specimenAvro.BinaryFromNative(nil, specimenMap)
	if err != nil {
		log.Fatalf("failed to convert Go map to Avro binary data: %v", err)
	}

	return binarySpecimenSegment, nil
}

func EncodeProveAndUploadSpecimenSegment(ctx context.Context, config *config.EthConfig, specimenAvro *goavro.Codec, specimenSegment *event.ReplicationSegment, specimenBucket, segmentName string, storage *storage.Client, ethClient *ethclient.Client, proofChain string) (string, error) {

	specimenSegmentAvro, err := encodeSpecimenSegmentToAvro(specimenAvro, specimenSegment)
	if err != nil {
		return "", err
	}
	log.Info("Submitting block-specimen segment proof for: ", segmentName)

	proofTxHash := make(chan string, 1)

	go proof.SendBlockSpecimenProofTx(ctx, config, proofChain, ethClient, specimenSegment.EndBlock, specimenSegment.Elements, specimenSegmentAvro, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {
		err := st.HandleObjectUploadToBucket(ctx, storage, "block-specimen", specimenBucket, segmentName, specimenSegmentAvro)
		if err != nil {
			return "", err
		}
		log.Info("Uploaded block-specimen segment event: ", segmentName, " with proof tx hash: ", pTxHash)

	} else {
		return "", fmt.Errorf("failed to prove & upload block-specimen segment event: %v", segmentName)
	}

	return pTxHash, nil

}
