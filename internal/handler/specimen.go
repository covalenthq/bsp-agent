package handler

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/common"
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

func (h *specimenHandler) Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) error {

	ctx := context.Background()
	replEvent, ok := e.(*event.ReplicationEvent)
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	replEvent.Hash = hash
	replEvent.DateTime = datetime

	specimen := &event.SpecimenEvent{
		ReplicationEvent: replEvent,
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

	EncodeSpecimenToAvro(specimen)

	log.Info("Submitting block-specimen proof for: ", block.Number.Uint64())

	proofTxHash := make(chan string, 1)

	go proof.SendBlockSpecimenProofTx(ctx, &config.EthConfig, ethProof, block.Number.Uint64(), *specimen, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {

		err = st.HandleObjectUploadToBucket(ctx, &config.GcpConfig, storage, string(replEvent.Type), replEvent.Hash, *specimen)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Uploaded block-specimen event: ", replEvent.Hash, " with proof tx hash: ", pTxHash)

	} else {
		log.Fatal(err)
	}

	return nil
}

func EncodeSpecimenToAvro(blockSpecimen interface{}) {
	codec, err := goavro.NewCodec(`
	{
		"type": "record",
		"name": "BlockSpecimen",
		"namespace": "com.covalenthq.blockspecimen.avro",
		"fields": [
		  {
			"name": "ReplicationEvent",
			"type": {
			  "name": "ReplicationEvent",
			  "type": "record",
			  "fields": [
				{
				  "name": "ID",
				  "type": "string"
				},
				{
				  "name": "type",
				  "type": "string"
				},
				{
				  "name": "hash",
				  "type": "string"
				},
				{
				  "name": "datetime",
				  "type": "string"
				}
			  ]
			}
		  },
		  {
			"name": "specimen",
			"type": {
			  "name": "specimen",
			  "type": "record",
			  "fields": [
				{
				  "name": "AccountRead",
				  "type": {
					"type": "array",
					"items": {
					  "name": "AccountRead_record",
					  "type": "record",
					  "fields": [
						{
						  "name": "Address",
						  "type": "string"
						},
						{
						  "name": "Nonce",
						  "type": "int"
						},
						{
						  "name": "Balance",
						  "type": "double"
						},
						{
						  "name": "CodeHash",
						  "type": "string"
						}
					  ]
					}
				  }
				},
				{
				  "name": "StorageRead",
				  "type": {
					"type": "array",
					"items": {
					  "name": "StorageRead_record",
					  "type": "record",
					  "fields": [
						{
						  "name": "Account",
						  "type": "string"
						},
						{
						  "name": "SlotKey",
						  "type": "string"
						},
						{
						  "name": "Value",
						  "type": "string"
						}
					  ]
					}
				  }
				},
				{
				  "name": "CodeRead",
				  "type": {
					"type": "array",
					"items": {
					  "name": "CodeRead_record",
					  "type": "record",
					  "fields": [
						{
						  "name": "Hash",
						  "type": "string"
						},
						{
						  "name": "Code",
						  "type": "string"
						}
					  ]
					}
				  }
				}
			  ]
			}
		  }
		]
	  }`)
	if err != nil {
		fmt.Println(err)
	}

	specimenMap, err := utils.StructToMap(blockSpecimen)
	if err != nil {
		fmt.Println(err)
	}

	// Convert native Go form to binary Avro data
	binary, err := codec.BinaryFromNative(nil, specimenMap)
	if err != nil {
		log.Fatalf("Failed to convert Go map to Avro binary data: %v", err)
	}
	_ = binary

	//Convert binary Avro data back to native Go form
	native, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		fmt.Println(err)
	}

	//Convert native Go form to textual Avro data
	textual, err := codec.TextualFromNative(nil, native)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(textual))
}
