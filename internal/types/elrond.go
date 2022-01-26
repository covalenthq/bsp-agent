// Package types contains all the types used across the repo
//nolint:stylecheck,revive
package types

import "github.com/elodina/go-avro"

type ElrondBlockReplica struct {
	Block        *Block
	Transactions []*ElrondTransaction
	SCResults    []*SCResult
	Receipts     []*ElrondReceipt
	Logs         []*Log
	StateChanges []*AccountBalanceUpdate
}

func NewBlockResult() *ElrondBlockReplica {
	return &ElrondBlockReplica{
		Block:        NewBlock(),
		Transactions: make([]*ElrondTransaction, 0),
		SCResults:    make([]*SCResult, 0),
		Receipts:     make([]*ElrondReceipt, 0),
		Logs:         make([]*Log, 0),
		StateChanges: make([]*AccountBalanceUpdate, 0),
	}
}

func (o *ElrondBlockReplica) Schema() avro.Schema {
	if _BlockResult_schema_err != nil {
		panic(_BlockResult_schema_err)
	}

	return _BlockResult_schema
}

type Block struct {
	Nonce                 int64
	Round                 int64
	Epoch                 int32
	Hash                  []byte
	MiniBlocks            []*MiniBlock
	NotarizedBlocksHashes [][]byte
	Proposer              int64
	Validators            []int64
	PubKeysBitmap         []byte
	Size                  int64
	Timestamp             int64
	StateRootHash         []byte
	PrevHash              []byte
	ShardID               int32
	TxCount               int32
	AccumulatedFees       []byte
	DeveloperFees         []byte
	EpochStartBlock       bool
	EpochStartInfo        *EpochStartInfo
}

func NewBlock() *Block {
	return &Block{
		Hash:            make([]byte, 32),
		Validators:      make([]int64, 0),
		PubKeysBitmap:   []byte{},
		StateRootHash:   make([]byte, 32),
		AccumulatedFees: []byte{},
		DeveloperFees:   []byte{},
	}
}

func (o *Block) Schema() avro.Schema {
	if _Block_schema_err != nil {
		panic(_Block_schema_err)
	}

	return _Block_schema
}

type MiniBlock struct {
	Hash            []byte
	SenderShardID   int32
	ReceiverShardID int32
	Type            int32
	Timestamp       int64
	TxHashes        [][]byte
}

func NewMiniBlock() *MiniBlock {
	return &MiniBlock{
		Hash:     make([]byte, 32),
		TxHashes: make([][]byte, 0),
	}
}

func (o *MiniBlock) Schema() avro.Schema {
	if _MiniBlock_schema_err != nil {
		panic(_MiniBlock_schema_err)
	}

	return _MiniBlock_schema
}

type EpochStartInfo struct {
	TotalSupply                      []byte
	TotalToDistribute                []byte
	TotalNewlyMinted                 []byte
	RewardsPerBlock                  []byte
	RewardsForProtocolSustainability []byte
	NodePrice                        []byte
	PrevEpochStartRound              int32
	PrevEpochStartHash               []byte
}

func NewEpochStartInfo() *EpochStartInfo {
	return &EpochStartInfo{
		TotalSupply:                      []byte{},
		TotalToDistribute:                []byte{},
		TotalNewlyMinted:                 []byte{},
		RewardsPerBlock:                  []byte{},
		RewardsForProtocolSustainability: []byte{},
		NodePrice:                        []byte{},
	}
}

func (o *EpochStartInfo) Schema() avro.Schema {
	if _EpochStartInfo_schema_err != nil {
		panic(_EpochStartInfo_schema_err)
	}

	return _EpochStartInfo_schema
}

type ElrondTransaction struct {
	Hash             []byte
	MiniBlockHash    []byte
	BlockHash        []byte
	Nonce            int64
	Round            int64
	Value            []byte
	Receiver         []byte
	Sender           []byte
	ReceiverShard    int32
	SenderShard      int32
	GasPrice         int64
	GasLimit         int64
	Data             []byte
	Signature        []byte
	Timestamp        int64
	SenderUserName   []byte
	ReceiverUserName []byte
}

func NewTransaction() *ElrondTransaction {
	return &ElrondTransaction{
		Hash:             make([]byte, 32),
		MiniBlockHash:    make([]byte, 32),
		BlockHash:        make([]byte, 32),
		Value:            []byte{},
		Receiver:         make([]byte, 62),
		Sender:           make([]byte, 62),
		Data:             []byte{},
		SenderUserName:   []byte{},
		ReceiverUserName: []byte{},
	}
}

func (o *ElrondTransaction) Schema() avro.Schema {
	if _Transaction_schema_err != nil {
		panic(_Transaction_schema_err)
	}

	return _Transaction_schema
}

type SCResult struct {
	Hash           []byte
	Nonce          int64
	GasLimit       int64
	GasPrice       int64
	Value          []byte
	Sender         []byte
	Receiver       []byte
	RelayerAddr    []byte
	RelayedValue   []byte
	Code           []byte
	Data           []byte
	PrevTxHash     []byte
	OriginalTxHash []byte
	CallType       int32
	CodeMetadata   []byte
	ReturnMessage  []byte
	Timestamp      int64
}

func NewSCResult() *SCResult {
	return &SCResult{
		Hash:           make([]byte, 32),
		Value:          []byte{},
		Sender:         make([]byte, 62),
		Receiver:       make([]byte, 62),
		RelayedValue:   []byte{},
		Code:           []byte{},
		Data:           []byte{},
		PrevTxHash:     make([]byte, 32),
		OriginalTxHash: make([]byte, 32),
		CodeMetadata:   []byte{},
		ReturnMessage:  []byte{},
	}
}

func (o *SCResult) Schema() avro.Schema {
	if _SCResult_schema_err != nil {
		panic(_SCResult_schema_err)
	}

	return _SCResult_schema
}

type ElrondReceipt struct {
	Hash      []byte
	Value     []byte
	Sender    []byte
	Data      []byte
	TxHash    []byte
	Timestamp int64
}

func NewReceipt() *ElrondReceipt {
	return &ElrondReceipt{
		Hash:   make([]byte, 32),
		Value:  []byte{},
		Sender: make([]byte, 62),
		Data:   []byte{},
		TxHash: make([]byte, 32),
	}
}

func (o *ElrondReceipt) Schema() avro.Schema {
	if _Receipt_schema_err != nil {
		panic(_Receipt_schema_err)
	}

	return _Receipt_schema
}

type Log struct {
	ID      []byte
	Address []byte
	Events  []*Event
}

func NewLog() *Log {
	return &Log{
		ID:     make([]byte, 32),
		Events: make([]*Event, 0),
	}
}

func (o *Log) Schema() avro.Schema {
	if _Log_schema_err != nil {
		panic(_Log_schema_err)
	}

	return _Log_schema
}

type Event struct {
	Address    []byte
	Identifier []byte
	Topics     [][]byte
	Data       []byte
}

func NewEvent() *Event {
	return &Event{
		Identifier: []byte{},
		Topics:     make([][]byte, 0),
		Data:       []byte{},
	}
}

func (o *Event) Schema() avro.Schema {
	if _Event_schema_err != nil {
		panic(_Event_schema_err)
	}

	return _Event_schema
}

type AccountBalanceUpdate struct {
	Address []byte
	Balance []byte
	Nonce   int64
}

func NewAccountBalanceUpdate() *AccountBalanceUpdate {
	return &AccountBalanceUpdate{
		Address: make([]byte, 62),
		Balance: []byte{},
	}
}

func (o *AccountBalanceUpdate) Schema() avro.Schema {
	if _AccountBalanceUpdate_schema_err != nil {
		panic(_AccountBalanceUpdate_schema_err)
	}

	return _AccountBalanceUpdate_schema
}

// Generated by codegen. Please do not modify.
var _BlockResult_schema, _BlockResult_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "com.covalenthq.block.schema",
    "name": "BlockResult",
    "fields": [
        {
            "name": "Block",
            "type": {
                "type": "record",
                "name": "Block",
                "fields": [
                    {
                        "name": "Nonce",
                        "type": "long"
                    },
                    {
                        "name": "Round",
                        "type": "long"
                    },
                    {
                        "name": "Epoch",
                        "type": "int"
                    },
                    {
                        "name": "Hash",
                        "type": {
                            "type": "fixed",
                            "size": 32,
                            "name": "hash"
                        }
                    },
                    {
                        "name": "MiniBlocks",
                        "default": null,
                        "type": [
                            "null",
                            {
                                "type": "array",
                                "items": {
                                    "type": "record",
                                    "name": "MiniBlock",
                                    "fields": [
                                        {
                                            "name": "Hash",
                                            "type": {
                                                "type": "fixed",
                                                "size": 32,
                                                "name": "hash"
                                            }
                                        },
                                        {
                                            "name": "SenderShardID",
                                            "type": "int"
                                        },
                                        {
                                            "name": "ReceiverShardID",
                                            "type": "int"
                                        },
                                        {
                                            "name": "Type",
                                            "type": "int"
                                        },
                                        {
                                            "name": "Timestamp",
                                            "type": "long"
                                        },
                                        {
                                            "name": "TxHashes",
                                            "type": {
                                                "type": "array",
                                                "items": "bytes"
                                            }
                                        }
                                    ]
                                }
                            }
                        ]
                    },
                    {
                        "name": "NotarizedBlocksHashes",
                        "default": null,
                        "type": [
                            "null",
                            {
                                "type": "array",
                                "items": {
                                    "type": "fixed",
                                    "size": 32,
                                    "name": "hash"
                                }
                            }
                        ]
                    },
                    {
                        "name": "Proposer",
                        "type": "long"
                    },
                    {
                        "name": "Validators",
                        "type": {
                            "type": "array",
                            "items": "long"
                        }
                    },
                    {
                        "name": "PubKeysBitmap",
                        "type": "bytes"
                    },
                    {
                        "name": "Size",
                        "type": "long"
                    },
                    {
                        "name": "Timestamp",
                        "type": "long"
                    },
                    {
                        "name": "StateRootHash",
                        "type": {
                            "type": "fixed",
                            "size": 32,
                            "name": "hash"
                        }
                    },
                    {
                        "name": "PrevHash",
                        "default": null,
                        "type": [
                            "null",
                            {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        ]
                    },
                    {
                        "name": "ShardID",
                        "type": "int"
                    },
                    {
                        "name": "TxCount",
                        "type": "int"
                    },
                    {
                        "name": "AccumulatedFees",
                        "type": "bytes"
                    },
                    {
                        "name": "DeveloperFees",
                        "type": "bytes"
                    },
                    {
                        "name": "EpochStartBlock",
                        "type": "boolean"
                    },
                    {
                        "name": "EpochStartInfo",
                        "default": null,
                        "type": [
                            "null",
                            {
                                "type": "record",
                                "name": "EpochStartInfo",
                                "fields": [
                                    {
                                        "name": "TotalSupply",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "TotalToDistribute",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "TotalNewlyMinted",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "RewardsPerBlock",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "RewardsForProtocolSustainability",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "NodePrice",
                                        "type": "bytes"
                                    },
                                    {
                                        "name": "PrevEpochStartRound",
                                        "type": "int"
                                    },
                                    {
                                        "name": "PrevEpochStartHash",
                                        "default": null,
                                        "type": [
                                            "null",
                                            {
                                                "type": "fixed",
                                                "size": 32,
                                                "name": "hash"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        },
        {
            "name": "Transactions",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "Transaction",
                    "fields": [
                        {
                            "name": "Hash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "MiniBlockHash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "BlockHash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "Nonce",
                            "type": "long"
                        },
                        {
                            "name": "Round",
                            "type": "long"
                        },
                        {
                            "name": "Value",
                            "type": "bytes"
                        },
                        {
                            "name": "Receiver",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "Sender",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "ReceiverShard",
                            "type": "int"
                        },
                        {
                            "name": "SenderShard",
                            "type": "int"
                        },
                        {
                            "name": "GasPrice",
                            "type": "long"
                        },
                        {
                            "name": "GasLimit",
                            "type": "long"
                        },
                        {
                            "name": "Data",
                            "type": "bytes"
                        },
                        {
                            "name": "Signature",
                            "default": null,
                            "type": [
                                "null",
                                {
                                    "type": "fixed",
                                    "size": 64,
                                    "name": "signature"
                                }
                            ]
                        },
                        {
                            "name": "Timestamp",
                            "type": "long"
                        },
                        {
                            "name": "SenderUserName",
                            "type": "bytes"
                        },
                        {
                            "name": "ReceiverUserName",
                            "type": "bytes"
                        }
                    ]
                }
            }
        },
        {
            "name": "SCResults",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "SCResult",
                    "fields": [
                        {
                            "name": "Hash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "Nonce",
                            "type": "long"
                        },
                        {
                            "name": "GasLimit",
                            "type": "long"
                        },
                        {
                            "name": "GasPrice",
                            "type": "long"
                        },
                        {
                            "name": "Value",
                            "type": "bytes"
                        },
                        {
                            "name": "Sender",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "Receiver",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "RelayerAddr",
                            "default": null,
                            "type": [
                                "null",
                                {
                                    "type": "fixed",
                                    "size": 62,
                                    "name": "address"
                                }
                            ]
                        },
                        {
                            "name": "RelayedValue",
                            "type": "bytes"
                        },
                        {
                            "name": "Code",
                            "type": "bytes"
                        },
                        {
                            "name": "Data",
                            "type": "bytes"
                        },
                        {
                            "name": "PrevTxHash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "OriginalTxHash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "CallType",
                            "type": "int"
                        },
                        {
                            "name": "CodeMetadata",
                            "type": "bytes"
                        },
                        {
                            "name": "ReturnMessage",
                            "type": "bytes"
                        },
                        {
                            "name": "Timestamp",
                            "type": "long"
                        }
                    ]
                }
            }
        },
        {
            "name": "Receipts",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "Receipt",
                    "fields": [
                        {
                            "name": "Hash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "Value",
                            "type": "bytes"
                        },
                        {
                            "name": "Sender",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "Data",
                            "type": "bytes"
                        },
                        {
                            "name": "TxHash",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "Timestamp",
                            "type": "long"
                        }
                    ]
                }
            }
        },
        {
            "name": "Logs",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "Log",
                    "fields": [
                        {
                            "name": "ID",
                            "type": {
                                "type": "fixed",
                                "size": 32,
                                "name": "hash"
                            }
                        },
                        {
                            "name": "Address",
                            "default": null,
                            "type": [
                                "null",
                                {
                                    "type": "fixed",
                                    "size": 62,
                                    "name": "address"
                                }
                            ]
                        },
                        {
                            "name": "Events",
                            "type": {
                                "type": "array",
                                "items": {
                                    "type": "record",
                                    "name": "Event",
                                    "fields": [
                                        {
                                            "name": "Address",
                                            "default": null,
                                            "type": [
                                                "null",
                                                {
                                                    "type": "fixed",
                                                    "size": 62,
                                                    "name": "address"
                                                }
                                            ]
                                        },
                                        {
                                            "name": "Identifier",
                                            "type": "bytes"
                                        },
                                        {
                                            "name": "Topics",
                                            "type": {
                                                "type": "array",
                                                "items": "bytes"
                                            }
                                        },
                                        {
                                            "name": "Data",
                                            "type": "bytes"
                                        }
                                    ]
                                }
                            }
                        }
                    ]
                }
            }
        },
        {
            "name": "StateChanges",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "AccountBalanceUpdate",
                    "fields": [
                        {
                            "name": "Address",
                            "type": {
                                "type": "fixed",
                                "size": 62,
                                "name": "address"
                            }
                        },
                        {
                            "name": "Balance",
                            "type": "bytes"
                        },
                        {
                            "name": "Nonce",
                            "type": "long"
                        }
                    ]
                }
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Block_schema, _Block_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Block",
    "fields": [
        {
            "name": "Nonce",
            "type": "long"
        },
        {
            "name": "Round",
            "type": "long"
        },
        {
            "name": "Epoch",
            "type": "int"
        },
        {
            "name": "Hash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "MiniBlocks",
            "default": null,
            "type": [
                "null",
                {
                    "type": "array",
                    "items": {
                        "type": "record",
                        "name": "MiniBlock",
                        "fields": [
                            {
                                "name": "Hash",
                                "type": {
                                    "type": "fixed",
                                    "size": 32,
                                    "name": "hash"
                                }
                            },
                            {
                                "name": "SenderShardID",
                                "type": "int"
                            },
                            {
                                "name": "ReceiverShardID",
                                "type": "int"
                            },
                            {
                                "name": "Type",
                                "type": "int"
                            },
                            {
                                "name": "Timestamp",
                                "type": "long"
                            },
                            {
                                "name": "TxHashes",
                                "type": {
                                    "type": "array",
                                    "items": "bytes"
                                }
                            }
                        ]
                    }
                }
            ]
        },
        {
            "name": "NotarizedBlocksHashes",
            "default": null,
            "type": [
                "null",
                {
                    "type": "array",
                    "items": {
                        "type": "fixed",
                        "size": 32,
                        "name": "hash"
                    }
                }
            ]
        },
        {
            "name": "Proposer",
            "type": "long"
        },
        {
            "name": "Validators",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "PubKeysBitmap",
            "type": "bytes"
        },
        {
            "name": "Size",
            "type": "long"
        },
        {
            "name": "Timestamp",
            "type": "long"
        },
        {
            "name": "StateRootHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "PrevHash",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 32,
                    "name": "hash"
                }
            ]
        },
        {
            "name": "ShardID",
            "type": "int"
        },
        {
            "name": "TxCount",
            "type": "int"
        },
        {
            "name": "AccumulatedFees",
            "type": "bytes"
        },
        {
            "name": "DeveloperFees",
            "type": "bytes"
        },
        {
            "name": "EpochStartBlock",
            "type": "boolean"
        },
        {
            "name": "EpochStartInfo",
            "default": null,
            "type": [
                "null",
                {
                    "type": "record",
                    "name": "EpochStartInfo",
                    "fields": [
                        {
                            "name": "TotalSupply",
                            "type": "bytes"
                        },
                        {
                            "name": "TotalToDistribute",
                            "type": "bytes"
                        },
                        {
                            "name": "TotalNewlyMinted",
                            "type": "bytes"
                        },
                        {
                            "name": "RewardsPerBlock",
                            "type": "bytes"
                        },
                        {
                            "name": "RewardsForProtocolSustainability",
                            "type": "bytes"
                        },
                        {
                            "name": "NodePrice",
                            "type": "bytes"
                        },
                        {
                            "name": "PrevEpochStartRound",
                            "type": "int"
                        },
                        {
                            "name": "PrevEpochStartHash",
                            "default": null,
                            "type": [
                                "null",
                                {
                                    "type": "fixed",
                                    "size": 32,
                                    "name": "hash"
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _MiniBlock_schema, _MiniBlock_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "MiniBlock",
    "fields": [
        {
            "name": "Hash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "SenderShardID",
            "type": "int"
        },
        {
            "name": "ReceiverShardID",
            "type": "int"
        },
        {
            "name": "Type",
            "type": "int"
        },
        {
            "name": "Timestamp",
            "type": "long"
        },
        {
            "name": "TxHashes",
            "type": {
                "type": "array",
                "items": "bytes"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _EpochStartInfo_schema, _EpochStartInfo_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "EpochStartInfo",
    "fields": [
        {
            "name": "TotalSupply",
            "type": "bytes"
        },
        {
            "name": "TotalToDistribute",
            "type": "bytes"
        },
        {
            "name": "TotalNewlyMinted",
            "type": "bytes"
        },
        {
            "name": "RewardsPerBlock",
            "type": "bytes"
        },
        {
            "name": "RewardsForProtocolSustainability",
            "type": "bytes"
        },
        {
            "name": "NodePrice",
            "type": "bytes"
        },
        {
            "name": "PrevEpochStartRound",
            "type": "int"
        },
        {
            "name": "PrevEpochStartHash",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 32,
                    "name": "hash"
                }
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Transaction_schema, _Transaction_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Transaction",
    "fields": [
        {
            "name": "Hash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "MiniBlockHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "BlockHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "Nonce",
            "type": "long"
        },
        {
            "name": "Round",
            "type": "long"
        },
        {
            "name": "Value",
            "type": "bytes"
        },
        {
            "name": "Receiver",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "Sender",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "ReceiverShard",
            "type": "int"
        },
        {
            "name": "SenderShard",
            "type": "int"
        },
        {
            "name": "GasPrice",
            "type": "long"
        },
        {
            "name": "GasLimit",
            "type": "long"
        },
        {
            "name": "Data",
            "type": "bytes"
        },
        {
            "name": "Signature",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 64,
                    "name": "signature"
                }
            ]
        },
        {
            "name": "Timestamp",
            "type": "long"
        },
        {
            "name": "SenderUserName",
            "type": "bytes"
        },
        {
            "name": "ReceiverUserName",
            "type": "bytes"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _SCResult_schema, _SCResult_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "SCResult",
    "fields": [
        {
            "name": "Hash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "Nonce",
            "type": "long"
        },
        {
            "name": "GasLimit",
            "type": "long"
        },
        {
            "name": "GasPrice",
            "type": "long"
        },
        {
            "name": "Value",
            "type": "bytes"
        },
        {
            "name": "Sender",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "Receiver",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "RelayerAddr",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 62,
                    "name": "address"
                }
            ]
        },
        {
            "name": "RelayedValue",
            "type": "bytes"
        },
        {
            "name": "Code",
            "type": "bytes"
        },
        {
            "name": "Data",
            "type": "bytes"
        },
        {
            "name": "PrevTxHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "OriginalTxHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "CallType",
            "type": "int"
        },
        {
            "name": "CodeMetadata",
            "type": "bytes"
        },
        {
            "name": "ReturnMessage",
            "type": "bytes"
        },
        {
            "name": "Timestamp",
            "type": "long"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Receipt_schema, _Receipt_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Receipt",
    "fields": [
        {
            "name": "Hash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "Value",
            "type": "bytes"
        },
        {
            "name": "Sender",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "Data",
            "type": "bytes"
        },
        {
            "name": "TxHash",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "Timestamp",
            "type": "long"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Log_schema, _Log_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Log",
    "fields": [
        {
            "name": "ID",
            "type": {
                "type": "fixed",
                "size": 32,
                "name": "hash"
            }
        },
        {
            "name": "Address",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 62,
                    "name": "address"
                }
            ]
        },
        {
            "name": "Events",
            "type": {
                "type": "array",
                "items": {
                    "type": "record",
                    "name": "Event",
                    "fields": [
                        {
                            "name": "Address",
                            "default": null,
                            "type": [
                                "null",
                                {
                                    "type": "fixed",
                                    "size": 62,
                                    "name": "address"
                                }
                            ]
                        },
                        {
                            "name": "Identifier",
                            "type": "bytes"
                        },
                        {
                            "name": "Topics",
                            "type": {
                                "type": "array",
                                "items": "bytes"
                            }
                        },
                        {
                            "name": "Data",
                            "type": "bytes"
                        }
                    ]
                }
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Event_schema, _Event_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Event",
    "fields": [
        {
            "name": "Address",
            "default": null,
            "type": [
                "null",
                {
                    "type": "fixed",
                    "size": 62,
                    "name": "address"
                }
            ]
        },
        {
            "name": "Identifier",
            "type": "bytes"
        },
        {
            "name": "Topics",
            "type": {
                "type": "array",
                "items": "bytes"
            }
        },
        {
            "name": "Data",
            "type": "bytes"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _AccountBalanceUpdate_schema, _AccountBalanceUpdate_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "AccountBalanceUpdate",
    "fields": [
        {
            "name": "Address",
            "type": {
                "type": "fixed",
                "size": 62,
                "name": "address"
            }
        },
        {
            "name": "Balance",
            "type": "bytes"
        },
        {
            "name": "Nonce",
            "type": "long"
        }
    ]
}`)
