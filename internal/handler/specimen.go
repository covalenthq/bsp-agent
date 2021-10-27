package handler

import (
	"bytes"
	"context"
	"fmt"
	"time"

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

func (h *specimenHandler) Handle(config *config.Config, storage *storage.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) (*event.SpecimenEvent, *event.ResultEvent, error) {

	replEvent, ok := e.(*event.ReplicationEvent)
	if !ok {
		return nil, nil, fmt.Errorf("incorrect event type: %v", replEvent.Type)
	}

	replEvent.Hash = hash
	replEvent.DateTime = datetime

	specimen := &event.SpecimenEvent{
		ReplicationEvent: replEvent,
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

func encodeSpecimenSegmentToAvro(blockSpecimenSegment interface{}) ([]byte, error) {
	codec, err := goavro.NewCodec(`
	{
		"type":"record",
		"namespace":"com.covalenthq.bsp.avro",
		"name":"BlockReplicationSegment",
		"fields":[
		   {
			  "name":"BlockSpecimen",
			  "type":{
				 "type":"array",
				 "items":{
					"name":"BlockSpecimen",
					"type":"record",
					"fields":[
					   {
						  "name":"ReplicationEvent",
						  "type":{
							 "name":"ReplicationEvent",
							 "type":"record",
							 "fields":[
								{
								   "name":"ID",
								   "type":"string"
								},
								{
								   "name":"type",
								   "type":"string"
								},
								{
								   "name":"hash",
								   "type":"string"
								},
								{
								   "name":"datetime",
								   "type":"string"
								}
							 ]
						  }
					   },
					   {
						  "name":"specimen",
						  "type":{
							 "name":"specimen",
							 "type":"record",
							 "fields":[
								{
								   "name":"Hash",
								   "type":"string"
								},
								{
								   "name":"Header",
								   "type":{
									  "name":"Header",
									  "type":"record",
									  "fields":[
										 {
											"name":"parentHash",
											"type":"string"
										 },
										 {
											"name":"sha3Uncles",
											"type":"string"
										 },
										 {
											"name":"miner",
											"type":"string"
										 },
										 {
											"name":"stateRoot",
											"type":"string"
										 },
										 {
											"name":"transactionsRoot",
											"type":"string"
										 },
										 {
											"name":"receiptsRoot",
											"type":"string"
										 },
										 {
											"name":"logsBloom",
											"type":{
											   "type":"array",
											   "items":"int"
											}
										 },
										 {
											"name":"difficulty",
											"type":"double"
										 },
										 {
											"name":"number",
											"type":"int"
										 },
										 {
											"name":"gasLimit",
											"type":"int"
										 },
										 {
											"name":"gasUsed",
											"type":"int"
										 },
										 {
											"name":"timestamp",
											"type":"int"
										 },
										 {
											"name":"extraData",
											"type":"string"
										 },
										 {
											"name":"mixHash",
											"type":"string"
										 },
										 {
											"name":"nonce",
											"type":{
											   "type":"array",
											   "items":"int"
											}
										 },
										 {
											"name":"baseFeePerGas",
											"type":["null", "int"],
											"default":null
										 }
									  ]
								   }
								},
								{
								   "name":"Transactions",
								   "type":{
									  "type":"array",
									  "items":{
										 "name":"Transactions_record",
										 "type":"record",
										 "fields":[
											{
											   "name":"nonce",
											   "type":"int"
											},
											{
											   "name":"gasPrice",
											   "type":"long"
											},
											{
											   "name":"gas",
											   "type":"int"
											},
											{
											   "name":"from",
											   "type":"string"
											},
											{
											   "name":"to",
											   "type":"string",
											   "default":"\u00ff"
											},
											{
											   "name":"value",
											   "type":"double"
											},
											{
											   "name":"input",
											   "type":"string"
											}
										 ]
									  }
								   }
								},
								{
								   "name":"uncles",
								   "type":[
									  "null",
									  {
										 "type":"array",
										 "items":"Header"
									  }
								   ],
								   "default":null
								},
								{
								   "name":"State",
								   "type":{
									  "name":"State",
									  "type":"record",
									  "fields":[
										 {
											"name":"AccountRead",
											"type":{
											   "type":"array",
											   "items":{
												  "name":"AccountRead_record",
												  "type":"record",
												  "fields":[
													 {
														"name":"Address",
														"type":"string"
													 },
													 {
														"name":"Nonce",
														"type":"int"
													 },
													 {
														"name":"Balance",
														"type":"double"
													 },
													 {
														"name":"CodeHash",
														"type":"string"
													 }
												  ]
											   }
											}
										 },
										 {
											"name":"StorageRead",
											"type":{
											   "type":"array",
											   "items":{
												  "name":"StorageRead_record",
												  "type":"record",
												  "fields":[
													 {
														"name":"Account",
														"type":"string"
													 },
													 {
														"name":"SlotKey",
														"type":"string"
													 },
													 {
														"name":"Value",
														"type":"string"
													 }
												  ]
											   }
											}
										 },
										 {
											"name":"CodeRead",
											"type":{
											   "type":"array",
											   "items":{
												  "name":"CodeRead_record",
												  "type":"record",
												  "fields":[
													 {
														"name":"Hash",
														"type":"string"
													 },
													 {
														"name":"Code",
														"type":"string"
													 }
												  ]
											   }
											}
										 }
									  ]
								   }
								}
							 ]
						  }
					   }
					]
				 }
			  }
		   },
		   {
			  "name":"StartBlock",
			  "type":"long"
		   },
		   {
			  "name":"EndBlock",
			  "type":"long"
		   },
		   {
			  "name":"Elements",
			  "type":"long"
		   }
		]
	 }`)
	if err != nil {
		return nil, err
	}

	specimenMap, err := utils.StructToMap(blockSpecimenSegment)
	if err != nil {
		return nil, err
	}

	// Convert native Go form to binary Avro data
	binarySpecimenSegment, err := codec.BinaryFromNative(nil, specimenMap)
	if err != nil {
		log.Fatalf("failed to convert Go map to Avro binary data: %v", err)
	}

	return binarySpecimenSegment, nil
}

func EncodeProveAndUploadSpecimenSegment(ctx context.Context, config *config.Config, specimenSegment *event.SpecimenSegment, segmentName string, storage *storage.Client, ethProof *ethclient.Client) (string, error) {

	specimenSegmentAvro, err := encodeSpecimenSegmentToAvro(specimenSegment)
	if err != nil {
		return "", err
	}
	log.Info("Submitting block-specimen segment proof for: ", segmentName)

	proofTxHash := make(chan string, 1)

	go proof.SendBlockSpecimenProofTx(ctx, &config.EthConfig, ethProof, specimenSegment.EndBlock, specimenSegment.Elements, specimenSegmentAvro, proofTxHash)

	pTxHash := <-proofTxHash

	if pTxHash != "" {
		err := st.HandleObjectUploadToBucket(ctx, &config.GcpConfig, storage, "block-specimen", segmentName, specimenSegmentAvro)
		if err != nil {
			return "", err
		}
		log.Info("Uploaded block-specimen segment event: ", segmentName, " with proof tx hash: ", pTxHash)

	} else {
		return "", fmt.Errorf("failed to prove & upload block-specimen segment event: %v", segmentName)
	}

	return pTxHash, nil

}
