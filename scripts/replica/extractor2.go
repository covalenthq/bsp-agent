//nolint:wrapcheck
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"

	"github.com/TylerBrock/colorjson"
	"github.com/linkedin/goavro/v2"
	"gopkg.in/avro.v0"
)

// generate specimen json and result json given avro encoded block replicas from bsp-agent
// go run extractor.go --binary-file-path="../../../Documents/bspec2/" --codec-path "../codec/block-ethereum.avsc" --indent-json 0 --end-block-number 15084335 --start-block-number 15084335  --chain-id "1"
// use https://bafybeid4ymbdmwos2lvjnkyqoeojtnr44lfqr5mpdwvtlsihvbtoagbiom.ipfs.dweb.link/ to download sample block replicas (avro codec version 0.2)

var (
	binaryFilePathFlag   string
	avroCodecPathFlag    string
	indentJSONFlag       int
	startBlockNumberFlag int64
	endBlockNumberFlag   int64
	chainIDFlag          string
	outputFilePathFlag   string
)

func generateTestForOneBlock() {
	replicaSegmentFiles := filterReplicaSegmentFiles(binaryFilePathFlag, startBlockNumberFlag, endBlockNumberFlag, chainIDFlag)
	codec := getAvroCodec(avroCodecPathFlag)

	// colorjson formatter can correctly format the block replicas big.Int
	// but it gives some parsing error due to introduction of escape sequences of color
	// so disabling here
	formatter := colorjson.NewFormatter()
	formatter.DisabledColor = true
	formatter.Indent = indentJSONFlag
	color.NoColor = true

	for _, replicaSegmentFile := range replicaSegmentFiles {
		filename := replicaSegmentFile.Name()
		fileNameSplit := strings.Split(filename, "-")
		directory := fileNameSplit[1] // block number is the directory
		if err := os.MkdirAll(outputFilePathFlag, os.ModePerm); err != nil {
			panic(err)
		}
		fileBuff, _, err := readReplicaFile(binaryFilePathFlag, filename)

		if err != nil {
			panic(err)
		}

		replicaSegment := decodeReplicaSegment(fileBuff, codec, *formatter)
		var replicaSegmentJSON []byte
		if replicaSegmentJSON, err = json.MarshalIndent(replicaSegment, "", " "); err != nil {
			panic(err)
		}

		writeFile(outputFilePathFlag, directory+".segment.json", replicaSegmentJSON)

		for _, component := range getComponents(replicaSegment) {
			filename := component.specimen.Header.Number.String()
			var componentJSON []byte
			if componentJSON, err = json.MarshalIndent(component.specimen, "", " "); err != nil {
				panic(err)
			}

			writeFile(outputFilePathFlag, filename+".specimen.json", componentJSON)

			if componentJSON, err = json.MarshalIndent(component.result, "", " "); err != nil {
				panic(err)
			}

			writeFile(outputFilePathFlag, filename+".result.json", componentJSON)
		}
	}
}

func writeFile(outputDir string, filename string, data []byte) {
	filepath := path.Join(outputDir, filename)
	if err := ioutil.WriteFile(filepath, data, 0600); err != nil {
		panic(err)
	}
}

func getAvroCodec(path string) *goavro.Codec {
	replicaAvro, err := avro.ParseSchemaFile(path)
	if err != nil {
		log.Error("unable to parse avro schema for specimen: ", err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		log.Error("unable to gen avro codec for specimen: ", err)
	}

	return replicaCodec
}

// picks out the replica files in `path` which fall within start-end range and given chainID
func filterReplicaSegmentFiles(path string, start int64, end int64, chainID string) []fs.FileInfo {
	allFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("unable to read files from directory: ", err)
	}
	sort.Slice(allFiles, func(i, j int) bool {
		return allFiles[i].Name() < allFiles[j].Name()
	})
	var filteredFiles []fs.FileInfo
	for _, fileInfo := range allFiles {
		fileName := fileInfo.Name()
		fmt.Println("filename: ", fileName)
		fileNameSplit := strings.Split(fileName, "-") // chainId-startBlocknumber-replica-....
		fBlockNumber := fileNameSplit[1]              // the block number
		fChainID := fileNameSplit[0]
		if chainID == fChainID {
			fileNameInt, err := strconv.ParseInt(fBlockNumber, 10, 0)
			if err != nil {
				panic(err)
			}
			// NOTE: the 2nd condition (fileNameInt <= end) works is segment length is 1 i.e. the block replica segment has 1 replica only (which is the case now)
			if fileNameInt >= start && fileNameInt <= end {
				filteredFiles = append(filteredFiles, fileInfo)
			}
		}
	}

	return filteredFiles
}

func readReplicaFile(directory, filename string) ([]byte, int, error) {
	file, err := os.Open(filepath.Join(filepath.Clean(directory), filepath.Base(filepath.Clean(filename))))
	if err != nil {
		return nil, 0, fmt.Errorf("error opening file %s: %w", filename, err)
	}
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, 0, fmt.Errorf("error in file info structure: %w", err)
	}
	size := stats.Size()
	bytes := make([]byte, size)
	buffr := bufio.NewReader(file)
	sizeBytes, err := buffr.Read(bytes)
	if err := file.Close(); err != nil {
		log.Error("Error closing file: ", err)
	}

	return bytes, sizeBytes, err
}

func decodeReplicaSegment(segmentBytes []byte, codec *goavro.Codec, formatter colorjson.Formatter) *event.ReplicationSegment {
	var fileMap map[string]interface{}
	native, _, err := codec.NativeFromBinary(segmentBytes) // convert binary avro data back to native Go form
	if err != nil {
		log.Error("unable to convert avro binary file to native Go from avro codec: ", err)
	}
	textAvro, err := codec.TextualFromNative(nil, native) // convert native Go form to textual avro data
	if err != nil {
		log.Error("unable to convert from native Go to textual avro: ", err)
	}
	decodedAvro := string(textAvro)
	err = json.Unmarshal([]byte(decodedAvro), &fileMap)
	if err != nil {
		log.Error("unable to unmarshal decoded AVRO binary: ", err)
	}

	fileMap = utils.UnwrapAvroUnion(fileMap)
	jsonMap, _ := formatter.Marshal(fileMap)
	// ReplicationSegment is a struct to store blockspecimen
	segment := event.ReplicationSegment{}

	if err = json.Unmarshal([]byte(string(jsonMap)), &segment); err != nil {
		panic(err)
	}

	return &segment
}

func getComponents(segment *event.ReplicationSegment) []*blockPair {
	var pairs []*blockPair
	for _, replica := range segment.BlockReplicaEvent {
		specimen := types.BlockReplica{
			Type:            "block-specimen",
			NetworkId:       replica.Data.NetworkId,
			Hash:            replica.Data.Hash,
			TotalDifficulty: replica.Data.TotalDifficulty,
			Header:          replica.Data.Header,
			Transactions:    replica.Data.Transactions,
			Uncles:          replica.Data.Uncles,
			Receipts:        []*types.Receipt{},
			Senders:         replica.Data.Senders,
			State:           replica.Data.State,
			Withdrawals:     replica.Data.Withdrawals,
			BlobTxSidecars: replica.Data.BlobTxSidecars,
		}

		result := types.BlockReplica{
			Type:            "block-result",
			NetworkId:       replica.Data.NetworkId,
			Hash:            replica.Data.Hash,
			TotalDifficulty: replica.Data.TotalDifficulty,
			Header:          replica.Data.Header,
			Transactions:    replica.Data.Transactions,
			Uncles:          replica.Data.Uncles,
			Receipts:        replica.Data.Receipts,
			Senders:         replica.Data.Senders,
			State:           &types.StateSpecimen{},
			Withdrawals:     replica.Data.Withdrawals,
			BlobTxSidecars: replica.Data.BlobTxSidecars,
		}
		pairs = append(pairs, &blockPair{
			specimen: &specimen,
			result:   &result,
		})
	}

	return pairs
}

type blockPair struct {
	specimen *types.BlockReplica
	result   *types.BlockReplica
}

func main() {
	flag.StringVar(&binaryFilePathFlag, "binary-file-path", config.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&avroCodecPathFlag, "codec-path", config.LookupEnvOrString("CodecPath", avroCodecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.IntVar(&indentJSONFlag, "indent-json", config.LookupEnvOrInt("IndentJson", indentJSONFlag), "allows for an indented view of the AVRO decoded JSON object")
	flag.Int64Var(&startBlockNumberFlag, "start-block-number", config.LookupEnvOrInt64("StartBlockNumber", 0), "block number range's start (default 0)")
	flag.Int64Var(&endBlockNumberFlag, "end-block-number", config.LookupEnvOrInt64("EndBlockNumber", math.MaxInt64), "block number range's end (default MAX_INT64)")
	flag.StringVar(&chainIDFlag, "chain-id", config.LookupEnvOrString("chainID", "1"), "chain id (default 1)")
	flag.StringVar(&outputFilePathFlag, "output-file-path", config.LookupEnvOrString("OutputFilePath", "./"), "local path to output files for the specimen/result.json (default: ./)")
	flag.Parse()

	if startBlockNumberFlag > endBlockNumberFlag {
		fmt.Println(errors.New("starting block should be smaller than or equal to the ending block"))
	}
	generateTestForOneBlock()
}
