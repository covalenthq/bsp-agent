package handler

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/proof"
	st "github.com/covalenthq/mq-store-agent/internal/storage"
	"github.com/covalenthq/mq-store-agent/internal/types"
)

type specimenHandler struct {
}

func NewSpecimenHandler() Handler {
	return &specimenHandler{}
}

func (h *specimenHandler) Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {

	ctx := context.Background()
	Event, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	Event.Hash = hash
	Event.DateTime = datetime

	specimen := &event.SpecimenEvent{
		ReplicationEvent: Event,
	}

	var decodedSpecimen types.BlockSpecimen
	err := rlp.Decode(bytes.NewReader(data), &decodedSpecimen)
	if err != nil {
		return fmt.Errorf("error decoding RLP bytes to block-specimen: %v", err)
	} else {
		specimen.Data = &decodedSpecimen
	}

	blockHash := common.HexToHash(specimen.ReplicationEvent.Hash)

	block, err := ethSource.HeaderByHash(ctx, blockHash)
	if err != nil {
		log.Error("error in getting block: ", err.Error())
	}

	log.Info("Submitting block-specimen proof for: ", block.Number.Uint64())

	proofTxHash := make(chan string, 1)

	go proof.SendBlockSpecimenProofTx(ctx, config, ethProof, block.Number.Uint64(), *specimen, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {

		err = st.HandleObjectUploadToBucket(ctx, config, storage, string(Event.Type), Event.Hash, *specimen)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Uploaded block-specimen event: ", Event.Hash, " with proof tx hash: ", pTxHash)

	} else {
		log.Fatal(err)
	}

	return nil
}
