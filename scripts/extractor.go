package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/TylerBrock/colorjson"
	"github.com/covalenthq/mq-store-agent/internal/utils"
	"github.com/linkedin/goavro/v2"
	"gopkg.in/avro.v0"
)

var (
	binaryFilePathFlag string
	codecPathFlag      string
	indentJSONFlag     int
)

func main() {
	flag.StringVar(&binaryFilePathFlag, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&codecPathFlag, "codec-path", utils.LookupEnvOrString("CodecPath", codecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.IntVar(&indentJSONFlag, "indent-json", utils.LookupEnvOrInt("IndentJson", indentJSONFlag), "allows for an indented view of the AVRO decoded JSON object")

	flag.Parse()
	fmt.Println("Agent command line config: ", utils.GetConfig(flag.CommandLine))

	codecPathFlag = utils.LookupEnvOrString("CodecPath", codecPathFlag)
	binaryFilePathFlag = utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag)
	indentJSONFlag = utils.LookupEnvOrInt("BinaryFilePath", indentJSONFlag)

	colorJSON := colorjson.NewFormatter()
	colorJSON.Indent = indentJSONFlag
	files := getBinFiles(binaryFilePathFlag)
	codec := getAvroCodec(codecPathFlag)

	for _, fileInf := range files {
		var fileMap map[string]interface{}
		filename := fileInf.Name()
		fileBuff, size, err := copyFileToMemory(binaryFilePathFlag, filename)
		if err != nil {
			log.Fatalf("unable to read binary file to memory: %v", err)
		}
		native, _, err := codec.NativeFromBinary(fileBuff) // convert binary avro data back to native Go form
		if err != nil {
			log.Fatalf("unable to convert avro binary file to native Go from avro codec: %v", err)
		}
		textAvro, err := codec.TextualFromNative(nil, native) // convert native Go form to textual avro data
		if err != nil {
			log.Fatalf("unable to convert from native Go to textual avro: %v", err)
		}
		decodedAvro := string(textAvro)
		err = json.Unmarshal([]byte(decodedAvro), &fileMap)
		if err != nil {
			log.Fatalf("unable to unmarshal decoded AVRO binary: %v", err)
		}
		colorJSONMap, _ := colorJSON.Marshal(fileMap)

		fmt.Println("\nfile: ", filepath.Join(binaryFilePathFlag, filepath.Base(filename)), "bytes: ", size, "\n", string(colorJSONMap))
	}
}

func getAvroCodec(path string) *goavro.Codec {
	replicaAvro, err := avro.ParseSchemaFile(path)
	if err != nil {
		log.Fatalf("unable to parse avro schema for specimen: %v", err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		log.Fatalf("unable to gen avro codec for specimen: %v", err)
	}

	return replicaCodec
}

func getBinFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("unable to read files from directory: %v", err)
	}

	return files
}

func copyFileToMemory(binaryFilePathFlag, filename string) ([]byte, int, error) {
	file, err := os.Open(filepath.Join(binaryFilePathFlag, filepath.Base(filename)))
	if err != nil {
		return nil, 0, fmt.Errorf("error opening file %s: %s", filename, err)
	}
	defer func() {
		if ferr := file.Close(); ferr != nil && err == nil {
			err = ferr
		}
	}()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, 0, fmt.Errorf("error in file info structure: %s", err)
	}
	size := stats.Size()
	bytes := make([]byte, size)
	buffr := bufio.NewReader(file)
	sizeBytes, err := buffr.Read(bytes)

	return bytes, sizeBytes, err
}
