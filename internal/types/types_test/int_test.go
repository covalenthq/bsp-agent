package types_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/handler"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/linkedin/goavro/v2"
	"github.com/ubiq/go-ubiq/common"
	"gopkg.in/avro.v0"
)

type Structure struct {
	X *types.BigInt `json:"x"`
}

func TestGolangBindings(_ *testing.T) {
	var val *big.Int
	var success bool
	if val, success = new(big.Int).SetString("21810676825935641000", 10); !success {
		panic("fail")
	}
	s := &Structure{
		X: &types.BigInt{Int: val},
	}

	dbytes, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dbytes))

	mapData := make(map[string]interface{})
	err = json.Unmarshal(dbytes, &mapData)
	if err != nil {
		panic(err)
	}

	fmt.Println(mapData["x"])
}

func TestAvroConversion(t *testing.T) {
	bigIntV, _ := new(big.Int).SetString("21810676825935641000", 10)
	replicaCodec := setupReplicaCodec()
	from := common.StringToAddress("0x21d3b08e73ba157cf46832f2b81644aeea4b4aa4")
	whash := common.BigToHash((&big.Int{}).SetInt64(3423423))
	seg := event.ReplicationSegment{
		BlockReplicaEvent: []*event.BlockReplicaEvent{
			{
				Hash: "somethin",
				Data: &types.BlockReplica{
					Type:            "awesome",
					NetworkId:       1,
					Hash:            common.Hash{},
					TotalDifficulty: new(types.BigInt).SetUint64(211),
					Header: &types.Header{
						Difficulty:      new(types.BigInt).SetUint64(211),
						Number:          new(types.BigInt).SetUint64(211),
						BaseFee:         new(types.BigInt).SetUint64(211),
						Extra:           []byte{},
						WithdrawalsHash: &whash,
					},
					Transactions: []*types.Transaction{
						{
							AccessList:   make(types.AccessList, 0),
							Amount:       &types.BigInt{Int: bigIntV},
							AccountNonce: 33,
							ChainId:      new(types.BigInt).SetUint64(1),
							GasFeeCap:    new(types.BigInt).SetUint64(211321234),
							GasTipCap:    new(types.BigInt).SetUint64(2132322),
							Price:        new(types.BigInt).SetUint64(99),
							Payload:      []byte{},
							Sender:       &from,
						},
					},
					Receipts: []*types.Receipt{},
					Senders:  []common.Address{},
					State: &types.StateSpecimen{
						AccountRead:   []*types.AccountRead{},
						StorageRead:   []*types.StorageRead{},
						CodeRead:      []*types.CodeRead{},
						BlockhashRead: []*types.BlockhashRead{},
					},
					Withdrawals: []*types.Withdrawal{
						{
							Index:     1,
							Validator: 33,
							Address:   from,
							Amount:    123455,
						},
					},
				},
			},
		},
	}

	_, err := handler.EncodeReplicaSegmentToAvro(replicaCodec, seg)
	if err != nil {
		t.Fatal(err)
	}
}

func setupReplicaCodec() *goavro.Codec {
	replicaAvro, err := avro.ParseSchemaFile("../../../codec/block-ethereum.avsc")
	if err != nil {
		panic(err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		panic(err)
	}

	return replicaCodec
}
