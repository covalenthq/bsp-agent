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
	BinaryFilePathFlag string
	CodecPathFlag      string
	IndentJsonFlag     int
)

func main() {
	flag.StringVar(&BinaryFilePathFlag, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", BinaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&CodecPathFlag, "codec-path", utils.LookupEnvOrString("CodecPath", CodecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.IntVar(&IndentJsonFlag, "indent-json", utils.LookupEnvOrInt("IndentJson", IndentJsonFlag), "allows for an indented view of the AVRO decoded JSON object")

	flag.Parse()
	fmt.Println("Agent command line config: ", utils.GetConfig(flag.CommandLine))

	CodecPathFlag = utils.LookupEnvOrString("CodecPath", CodecPathFlag)
	BinaryFilePathFlag = utils.LookupEnvOrString("BinaryFilePath", BinaryFilePathFlag)
	IndentJsonFlag = utils.LookupEnvOrInt("BinaryFilePath", IndentJsonFlag)

	colorJson := colorjson.NewFormatter()
	colorJson.Indent = IndentJsonFlag
	files := getBinFiles(BinaryFilePathFlag)
	codec := getAvroCodec(CodecPathFlag)

	for _, fileInf := range files {
		var fileMap map[string]interface{}
		filename := fileInf.Name()
		fileBuff, size, err := copyFileToMemory(BinaryFilePathFlag, filename)
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
		colorJsonMap, _ := colorJson.Marshal(fileMap)

		fmt.Println("\nfile: ", filepath.Join(BinaryFilePathFlag, filepath.Base(filename)), "bytes: ", size, "\n", string(colorJsonMap))
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

func copyFileToMemory(BinaryFilePathFlag, filename string) ([]byte, int, error) {
	file, err := os.Open(filepath.Join(BinaryFilePathFlag, filepath.Base(filename)))
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
