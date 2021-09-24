package handler

import (
	"bytes"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

type resultHandler struct {
}

func NewResultHandler() Handler {
	return &resultHandler{}
}

func (h *resultHandler) Handle(config *config.Config, storage *storage.Client, ethSource *ethclient.Client, ethProof *ethclient.Client, e event.Event, hash string, datetime time.Time, data []byte, retry bool) (*event.SpecimenEvent, *event.ResultEvent, error) {

	//ctx := context.Background()
	replEvent, ok := e.(*event.ReplicationEvent)
	if !ok {
		return nil, nil, fmt.Errorf("incorrect event type")
	}

	replEvent.Hash = hash
	replEvent.DateTime = datetime

	result := &event.ResultEvent{
		ReplicationEvent: replEvent,
	}

	var decodedResult types.BlockResult
	err := rlp.Decode(bytes.NewReader(data), &decodedResult)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RLP bytes to block-result: %v", err)
	} else {
		result.Data = &decodedResult
	}

	// blockHash := common.HexToHash(result.ReplicationEvent.Hash)
	// block, err := ethSource.HeaderByHash(ctx, blockHash)
	// if err != nil {
	// 	log.Error("error in getting block: ", err.Error())
	// }

	// EncodeResultToAvro(result)

	// log.Info("Submitting block-result proof for: ", block.Number.Uint64())

	// proofTxHash := make(chan string, 1)

	// go proof.SendBlockResultProofTx(ctx, &config.EthConfig, ethProof, block.Number.Uint64(), *result, proofTxHash)

	// pTxHash := <-proofTxHash

	// if pTxHash != "" {

	// 	err = st.HandleObjectUploadToBucket(ctx, &config.GcpConfig, storage, string(replEvent.Type), replEvent.Hash, *result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Info("Uploaded block-result event: ", replEvent.Hash, " with proof tx hash: ", pTxHash)

	// } else {
	// 	log.Errorf("Failed to prove & upload block-result event")
	// }

	return nil, result, nil
}

func EncodeResultSegmentToAvro(blockResultSegment interface{}) {
	codec, err := goavro.NewCodec(`
	{
		"type":"record",
		"namespace":"com.covalenthq.brp.avro",
		"name":"BlockReplicationSegment",
		"fields":[
		   {
			  "name":"BlockResult",
			  "type":{
				 "type":"array",
				 "items":{
				 	"name":"BlockResult",
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
						  "name":"result",
						  "type":{
							 "name":"result",
							 "type":"record",
							 "fields":[
								{
								   "name":"Hash",
								   "type":"string"
								},
								{
								   "name":"TotalDifficulty",
								   "type":"int"
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
											"type":"int"
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
											"type":"int"
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
											   "type":"string"
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
								   "name":"Receipts",
								   "type":{
									  "type":"array",
									  "items":{
										 "name":"Receipts_record",
										 "type":"record",
										 "fields":[
											{
											   "name":"PostStateOrStatus",
											   "type":"string"
											},
											{
											   "name":"CumulativeGasUsed",
											   "type":"int"
											},
											{
											   "name":"TxHash",
											   "type":"string"
											},
											{
											   "name":"ContractAddress",
											   "type":"string"
											},
											{
											   "name":"Logs",
											   "type":{
												  "type":"array",
												  "items":{
													 "name":"Logs_record",
													 "type":"record",
													 "fields":[
														{
														   "name":"address",
														   "type":"string"
														},
														{
														   "name":"topics",
														   "type":{
															  "type":"array",
															  "items":"string"
														   }
														},
														{
														   "name":"data",
														   "type":"string"
														},
														{
														   "name":"blockNumber",
														   "type":"int"
														},
														{
														   "name":"transactionHash",
														   "type":"string"
														},
														{
														   "name":"transactionIndex",
														   "type":"int"
														},
														{
														   "name":"blockHash",
														   "type":"string"
														},
														{
														   "name":"logIndex",
														   "type":"int"
														},
														{
														   "name":"removed",
														   "type":"boolean"
														}
													 ]
												  }
											   }
											},
											{
											   "name":"GasUsed",
											   "type":"int"
											}
										 ]
									  }
								   }
								},
								{
								   "name":"Senders",
								   "type":{
									  "type":"array",
									  "items":"string"
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
		fmt.Println(err)
	}

	resultMap, err := utils.StructToMap(blockResultSegment)
	if err != nil {
		fmt.Println(err)
	}

	// Convert native Go map[string]interface{} to binary Avro data
	binary, err := codec.BinaryFromNative(nil, resultMap)
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
	//return binary
}
