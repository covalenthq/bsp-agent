package types

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

func TestGolangBindings(t *testing.T) {
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
	//assert.Equal(t, string(dbytes), "{\"x\":[2,1,8,1,0,6,7,6,8,2,5,9,3,5,6,4,1,0,0,0]}", "marshalled bytes not same")

	mapData := make(map[string]interface{})
	err = json.Unmarshal(dbytes, &mapData)
	if err != nil {
		panic(err)
	}

	fmt.Println(mapData["x"])
}

func TestAnother(t *testing.T) {
	v, _ := new(big.Int).SetString("21810676825935641000", 10)
	replicaCodec := setupReplicaCodec()
	from := common.StringToAddress("0x21d3b08e73ba157cf46832f2b81644aeea4b4aa4")
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
						Difficulty: new(types.BigInt).SetUint64(211),
						Number:     new(types.BigInt).SetUint64(211),
						BaseFee:    new(types.BigInt).SetUint64(211),
						Extra:      []byte{},
					},
					Transactions: []*types.Transaction{
						{
							AccessList:   make(types.AccessList, 0),
							Amount:       &types.BigInt{Int: v},
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
