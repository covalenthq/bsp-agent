package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/TylerBrock/colorjson"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/fatih/color"

	//tt "github.com/ethereum/go-ethereum/core/types"
	"github.com/linkedin/goavro/v2"
	"gopkg.in/avro.v0"
)

func TestRlpCoding(t *testing.T) {
	type InputStructure struct {
		T1 *big.Int //`rlp:"nilList"`
		T2 *big.Int
	}

	type OutputStructure struct {
		T1 *types.BigInt //`rlp:"nilList"`
		T2 *types.BigInt
	}

	inp := InputStructure{T1: new(big.Int).SetUint64(12), T2: new(big.Int).SetUint64(8)}
	//vv := Sss{T1: nil, T2: new(big.Int).SetUint64(8)}
	out := OutputStructure{}
	bytes, err := rlp.EncodeToBytes(inp)
	if err != nil {
		t.Errorf("error is not nil: %v", err)
	}
	fmt.Println("size is:", len(bytes))
	fmt.Println("bytes is:", bytes)

	err = rlp.DecodeBytes(bytes, &out)
	if err != nil {
		t.Errorf("dec error is not nil: %v", err)
		return
	}

	if out.T1.Int.Cmp(inp.T1) != 0 {
		t.Errorf("encoded/decoded values don't match: %v %v", out.T1.Int, inp.T1)
		return
	}

	if out.T2.Int.Cmp(inp.T2) != 0 {
		t.Errorf("encoded/decoded values don't match: %v %v", out.T2.Int, inp.T2)
		return
	}

	fmt.Println(out.T1)
	fmt.Println(out.T2)
}

func TestRlpCodingWithNils(t *testing.T) {
	type InputStructure struct {
		T1 *big.Int `rlp:"nil"`
		T2 *big.Int
	}

	type OutputStructure struct {
		T1 *types.BigInt `rlp:"nilString"`
		T2 *types.BigInt
	}

	inp := InputStructure{T1: nil, T2: new(big.Int).SetUint64(8)}
	//vv := Sss{T1: nil, T2: new(big.Int).SetUint64(8)}
	out := OutputStructure{}
	bytes, err := rlp.EncodeToBytes(inp)
	if err != nil {
		t.Errorf("error is not nil: %v", err)
	}
	fmt.Println("size is:", len(bytes))
	fmt.Println("bytes is:", bytes)

	err = rlp.DecodeBytes(bytes, &out)
	if err != nil {
		t.Errorf("dec error is not nil: %v", err)
		return
	}

	if !(out.T1 == nil && inp.T1 == nil) {
		t.Errorf("encoded/decoded values don't match: %v %v", out.T1, inp.T1)
		return
	}

	if out.T2.Int.Cmp(inp.T2) != 0 {
		t.Errorf("encoded/decoded values don't match: %v %v", out.T2.Int, inp.T2)
		return
	}

	fmt.Println(out.T1)
	fmt.Println(out.T2)
}

func TestAvroEncodingForLogicalTypesAndBigInt(t *testing.T) {
	type StructT struct {
		V *types.BigInt `json:"v"`
		R string        `json:"r"`
	}

	replicaAvro, err := avro.ParseSchema("{\"type\":\"record\",\"name\":\"StructT\",\"fields\":[{\"name\":\"v\",\"type\":{\"type\":\"bytes\",\"logicalType\":\"decimal\",\"precision\":1000}},{\"name\":\"r\",\"type\":\"string\",\"default\":\"ok\"}]}")
	if err != nil {
		t.Fatal(err)
	}
	codec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		panic(err)
	}

	inp := StructT{V: new(types.BigInt).SetUint64(123)}
	replicaMap, err := utils.StructToMap(inp)
	if err != nil {
		t.Fatalf("error in converting struct to map: %v", err)
	}

	binaryReplicaSegment, err := codec.BinaryFromNative(nil, replicaMap)
	if err != nil {
		t.Fatal("failed to convert Go map to Avro binary data: ", err)
	}

	var fileMap map[string]interface{}
	native, _, err := codec.NativeFromBinary(binaryReplicaSegment) // convert binary avro data back to native Go form
	if err != nil {
		t.Fatal("unable to convert avro binary file to native Go from avro codec: ", err)
	}
	textAvro, err := codec.TextualFromNative(nil, native) // convert native Go form to textual avro data
	if err != nil {
		t.Fatal("unable to convert from native Go to textual avro: ", err)
	}
	decodedAvro := string(textAvro)
	err = json.Unmarshal([]byte(decodedAvro), &fileMap)
	if err != nil {
		t.Fatal("unable to unmarshal decoded AVRO binary: ", err)
	}

	formatter := colorjson.NewFormatter()
	formatter.DisabledColor = true
	color.NoColor = true
	jsonMap, err := formatter.Marshal(fileMap)
	if err != nil {
		t.Fatal("error marshalling", err)
	}

	out := &StructT{}
	if err = json.Unmarshal([]byte(string(jsonMap)), &out); err != nil {
		t.Fatal("unmarshalling failed: ", err)
	}

	if inp.V.Int.Cmp(out.V.Int) != 0 {
		t.Fatal("comparison failed", inp.V, out.V)
	}
}

func TestAvroEncodingForUnionSchemaAndNilBigInt(t *testing.T) {
	type StructT struct {
		V *types.BigInt `json:"v"`
		R string        `json:"r"`
	}

	// field v is a union of nil and logicalType
	replicaAvro, err := avro.ParseSchema("{\"type\":\"record\",\"name\":\"StructT\",\"fields\":[{\"name\":\"v\",\"type\":[\"null\",{\"type\":\"bytes\",\"logicalType\":\"decimal\",\"precision\":1000}],\"default\":\"null\"},{\"name\":\"r\",\"type\":\"string\",\"default\":\"ok\"}]}")
	if err != nil {
		t.Fatal(err)
	}
	codec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		panic(err)
	}

	//inp := StructT{V: new(types.BigInt).SetUint64(123)}
	inp := StructT{}
	replicaMap, err := utils.StructToMap(inp)
	if err != nil {
		t.Fatalf("error in converting struct to map: %v", err)
	}

	replicaMap = wrapAvroUnion(replicaMap)

	binaryReplicaSegment, err := codec.BinaryFromNative(nil, replicaMap)
	if err != nil {
		t.Fatal("failed to convert Go map to Avro binary data: ", err)
	}

	var fileMap map[string]interface{}
	native, _, err := codec.NativeFromBinary(binaryReplicaSegment) // convert binary avro data back to native Go form
	if err != nil {
		t.Fatal("unable to convert avro binary file to native Go from avro codec: ", err)
	}
	textAvro, err := codec.TextualFromNative(nil, native) // convert native Go form to textual avro data
	if err != nil {
		t.Fatal("unable to convert from native Go to textual avro: ", err)
	}
	decodedAvro := string(textAvro)
	err = json.Unmarshal([]byte(decodedAvro), &fileMap)
	if err != nil {
		t.Fatal("unable to unmarshal decoded AVRO binary: ", err)
	}

	formatter := colorjson.NewFormatter()
	formatter.DisabledColor = true
	color.NoColor = true
	jsonMap, err := formatter.Marshal(fileMap)
	if err != nil {
		t.Fatal("error marshalling", err)
	}

	out := &StructT{}
	if err = json.Unmarshal([]byte(string(jsonMap)), &out); err != nil {
		t.Fatal("unmarshalling failed: ", err)
	}

	if inp.V != out.V || inp.R != out.R || inp.V != nil {
		t.Fatal("comparison failed", inp, out)
	}
}

func TestAvroEncodingForUnionSchemaAndBigInt(t *testing.T) {
	type StructT struct {
		V *types.BigInt `json:"v"`
		R string        `json:"r"`
	}

	// field v is a union of nil and logicalType
	replicaAvro, err := avro.ParseSchema("{\"type\":\"record\",\"name\":\"StructT\",\"fields\":[{\"name\":\"v\",\"type\":[\"null\",{\"type\":\"bytes\",\"logicalType\":\"decimal\",\"precision\":1000}],\"default\":\"null\"},{\"name\":\"r\",\"type\":\"string\",\"default\":\"ok\"}]}")
	if err != nil {
		t.Fatal(err)
	}
	codec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		panic(err)
	}

	inp := StructT{V: new(types.BigInt).SetUint64(123)}
	//ss := StructT{}
	replicaMap, err := utils.StructToMap(inp)
	if err != nil {
		t.Fatalf("error in converting struct to map: %v", err)
	}

	replicaMap = wrapAvroUnion(replicaMap)

	binaryReplicaSegment, err := codec.BinaryFromNative(nil, replicaMap)
	if err != nil {
		t.Fatal("failed to convert Go map to Avro binary data: ", err)
	}

	var fileMap map[string]interface{}
	native, _, err := codec.NativeFromBinary(binaryReplicaSegment) // convert binary avro data back to native Go form
	if err != nil {
		t.Fatal("unable to convert avro binary file to native Go from avro codec: ", err)
	}
	textAvro, err := codec.TextualFromNative(nil, native) // convert native Go form to textual avro data
	if err != nil {
		t.Fatal("unable to convert from native Go to textual avro: ", err)
	}
	decodedAvro := string(textAvro)
	err = json.Unmarshal([]byte(decodedAvro), &fileMap)
	if err != nil {
		t.Fatal("unable to unmarshal decoded AVRO binary: ", err)
	}

	fileMap = unwrapAvroUnion(fileMap)

	formatter := colorjson.NewFormatter()
	formatter.DisabledColor = true
	color.NoColor = true
	jsonMap, err := formatter.Marshal(fileMap)
	if err != nil {
		t.Fatal("error marshalling", err)
	}

	out := &StructT{}
	if err = json.Unmarshal([]byte(string(jsonMap)), &out); err != nil {
		t.Fatal("unmarshalling failed: ", err)
	}

	if inp.V.Int.Cmp(out.V.Int) != 0 || inp.R != out.R {
		t.Fatal("comparison failed", inp, out)
	}
}

func wrapAvroUnion(data map[string]interface{}) map[string]interface{} {
	if data["v"] == nil {
		data["v"] = goavro.Union("null", nil)
	} else {
		data["v"] = goavro.Union("bytes", data["v"])
	}

	return data
}

func unwrapAvroUnion(data map[string]interface{}) map[string]interface{} {
	if data["v"] != nil {
		mp := data["v"].(map[string]interface{})
		data["v"] = mp["bytes"]
	}

	return data
}
