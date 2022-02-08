//nolint:wrapcheck
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/TylerBrock/colorjson"
	"github.com/covalenthq/mq-store-agent/internal/utils"
	"github.com/linkedin/goavro/v2"
	"gopkg.in/avro.v0"
)

var (
	binaryFilePathFlag string
	avroCodecPathFlag  string
	indentJSONFlag     int
)

func main() {
	flag.StringVar(&binaryFilePathFlag, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&avroCodecPathFlag, "codec-path", utils.LookupEnvOrString("CodecPath", avroCodecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.IntVar(&indentJSONFlag, "indent-json", utils.LookupEnvOrInt("IndentJson", indentJSONFlag), "allows for an indented view of the AVRO decoded JSON object")

	flag.Parse()
	fmt.Println("bsp-extractor command line config: ", utils.GetConfig(flag.CommandLine))

	avroCodecPathFlag = utils.LookupEnvOrString("CodecPath", avroCodecPathFlag)
	binaryFilePathFlag = utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag)
	indentJSONFlag = utils.LookupEnvOrInt("BinaryFilePath", indentJSONFlag)

	colorJSON := colorjson.NewFormatter()
	colorJSON.Indent = indentJSONFlag
	files := getBinFiles(binaryFilePathFlag)
	codec := getAvroCodec(avroCodecPathFlag)

	for _, fileInf := range files {
		var fileMap map[string]interface{}
		filename := fileInf.Name()
		fileBuff, size, err := copyFileToMemory(binaryFilePathFlag, filename)
		if err != nil {
			log.Error("unable to read binary file to memory: ", err)
		}
		native, _, err := codec.NativeFromBinary(fileBuff) // convert binary avro data back to native Go form
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
		colorJSONMap, _ := colorJSON.Marshal(fileMap)

		fmt.Println("\nfile: ", filepath.Join(binaryFilePathFlag, filepath.Base(filename)), "bytes: ", size, "\n", string(colorJSONMap))
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

func getBinFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("unable to read files from directory: ", err)
	}

	return files
}

func copyFileToMemory(binaryFilePathFlag, filename string) ([]byte, int, error) {
	file, err := os.Open(filepath.Join(filepath.Clean(binaryFilePathFlag), filepath.Base(filepath.Clean(filename))))
	if err != nil {
		return nil, 0, fmt.Errorf("error opening file %s: %w", filename, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error("Error closing file: ", err)
		}
	}()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, 0, fmt.Errorf("error in file info structure: %w", err)
	}
	size := stats.Size()
	bytes := make([]byte, size)
	buffr := bufio.NewReader(file)
	sizeBytes, err := buffr.Read(bytes)

	return bytes, sizeBytes, err
}
