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
	BinaryFilePath string
	CodecPath      string
	IndentJson     int
)

func main() {

	flag.StringVar(&BinaryFilePath, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", BinaryFilePath), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&CodecPath, "codec-path", utils.LookupEnvOrString("CodecPath", CodecPath), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.IntVar(&IndentJson, "indent-json", utils.LookupEnvOrInt("IndentJson", IndentJson), "allows for an indented view of the AVRO decoded JSON object")

	flag.Parse()
	fmt.Println("Agent command line config: ", utils.GetConfig(flag.CommandLine))

	CodecPath = utils.LookupEnvOrString("CodecPath", CodecPath)
	BinaryFilePath = utils.LookupEnvOrString("BinaryFilePath", BinaryFilePath)
	IndentJson = utils.LookupEnvOrInt("BinaryFilePath", IndentJson)

	colorJson := colorjson.NewFormatter()
	colorJson.Indent = IndentJson
	files := readAllFiles(BinaryFilePath)
	codec := getAvroCodec(CodecPath)

	for _, file := range files {
		var obj map[string]interface{}
		filename := file.Name()

		f, err := os.Open(filepath.Join(BinaryFilePath, filepath.Base(filename)))
		if err != nil {
			fmt.Printf("error opening %s: %s", filename, err)
			return
		}
		defer f.Close()

		buf, err := retrieveROM(f.Name())
		if err != nil {
			log.Fatalf("unable to read file %v", err)
		}

		// Convert binary Avro data back to native Go form
		native, _, err := codec.NativeFromBinary(buf)
		if err != nil {
			fmt.Println(err)
		}
		// Convert native Go form to textual Avro data
		textual, err := codec.TextualFromNative(nil, native)
		if err != nil {
			fmt.Println(err)
		}

		decodedString := string(textual)
		json.Unmarshal([]byte(decodedString), &obj)
		s, _ := colorJson.Marshal(obj)
		fmt.Println("\nfilename:", filename, "\n", string(s))
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

func readAllFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("unable to read files from directory: %v", err)
	}
	return files
}

func retrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
