// Package utils contains all the utilites used across the repo
//
//nolint:wrapcheck
package utils

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/elodina/go-avro"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const (
	clientIdentifier = "bspagent" // Client identifier to advertise over the network
)

const (
	// BspAgentVersionMajor is Major version component of the current release
	BspAgentVersionMajor = 1
	// BspAgentVersionMinor is Minor version component of the current release
	BspAgentVersionMinor = 9
	// BspAgentVersionPatch is Patch version component of the current release
	BspAgentVersionPatch = 1
)

// BspAgentVersion holds the textual version string.
var BspAgentVersion = func() string {
	return fmt.Sprintf("%d.%d.%d", BspAgentVersionMajor, BspAgentVersionMinor, BspAgentVersionPatch)
}()

// NewRedisClient provides a new redis client using a redis config
func NewRedisClient(rconfig *config.RedisConfig) (*redis.Client, string, string, error) {
	var pwd string
	redisURL, err := url.Parse(rconfig.RedisURL)
	if err != nil {
		log.Fatalf("unable to parse redis connection string: %v", err)
	}

	pass, _ := redisURL.User.Password()
	if pass != "" {
		log.Fatal("remove password from connection string cli flag and add it in .envrc as `REDIS_PWD`")
	}
	pwd = rconfig.Password

	dbString := strings.ReplaceAll(redisURL.Path, "/", "")
	urlMap, err := url.ParseQuery(redisURL.RawQuery)
	if err != nil {
		log.Fatalf("unable to parse redis connection string query string: %v", err)
	}

	dbInt, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL.Host,
		Password: pwd,
		DB:       dbInt, // use default DB
	})
	streamKey := urlMap["topic"][0]
	consumerGroup := redisURL.Fragment
	_, err = redisClient.Ping().Result()

	return redisClient, streamKey, consumerGroup, err
}

// NewEthClient initializes a new ethereum client using an address string
func NewEthClient(address string) (*ethclient.Client, error) {
	ethClient, err := ethclient.Dial(address)
	if err != nil {
		log.Fatalf("error in getting eth client: %v", err)
	}

	return ethClient, nil
}

// StructToMap converts a struct to go map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error in converting struct to map: %w", err)
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshaling to JSON: %w", err)
	}

	return mapData, nil
}

// AckTrimStreamSegment acknowledges a stream segment from the redis stream
func AckTrimStreamSegment(redisClient *redis.Client, segmentLength int, streamKey, consumerGroup string, streamIDs []string) (int64, error) {
	if len(streamIDs) == segmentLength {
		redisClient.XAck(streamKey, consumerGroup, streamIDs...)
		redisClient.XDel(streamKey, streamIDs...)
		xlen := redisClient.XLen(streamKey)
		length, err := xlen.Result()
		if err != nil {
			log.Error("failed to extract length of stream key: ", streamKey, "with error: ", err)
		}

		return length, nil
	}

	return 0, fmt.Errorf("failed to match streamIDs length to segment length config")
}

// GetConfig retrieves the config from the config packages
func GetConfig(fs *flag.FlagSet) []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *flag.Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}

// EncodeAvro returns a byte slice representing the binary encoding of the input avro record
func EncodeAvro(record avro.AvroRecord) ([]byte, error) {
	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(record.Schema())

	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	err := writer.Write(record, encoder)
	if err != nil {
		return nil, fmt.Errorf("error in encoding to AVRO: %w", err)
	}

	return buffer.Bytes(), nil
}

// DecodeAvro tries to decode a data buffer, read it and store it on the input record. If successfully, the record is filled with data from the buffer, otherwise an error might be returned
func DecodeAvro(record avro.AvroRecord, buffer []byte) error {
	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(record.Schema())
	decoder := avro.NewBinaryDecoder(buffer)

	return reader.Read(record, decoder)
}

// ExpandPath expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func ExpandPath(fsPath string) string {
	if strings.HasPrefix(fsPath, "~/") || strings.HasPrefix(fsPath, "~\\") {
		if home := HomeDir(); home != "" {
			fsPath = home + fsPath[1:]
		}
	}

	return path.Clean(os.ExpandEnv(fsPath))
}

// HomeDir returns full path of home directory for current user
func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}

	return ""
}

// GetLogLocationURL gets full path of log directory for current user or creates it
func GetLogLocationURL(logPath string) (*url.URL, error) {
	logLocation := ExpandPath(logPath)
	locationURL, err := url.Parse(logLocation)
	if err == nil {
		if _, existErr := os.Stat(locationURL.Path); os.IsNotExist(existErr) {
			// directory doesn't exist, create
			createErr := os.Mkdir(locationURL.Path, os.ModePerm)
			if createErr != nil {
				return nil, fmt.Errorf("error creating the directory: %w", createErr)
			}
		}

		if !Writable(locationURL.Path) {
			return nil, fmt.Errorf("write access not present for given log location")
		}

		return locationURL, nil
	}

	return locationURL, fmt.Errorf("log-folder: %w", err)
}

// Writable informs if path is writable to or not
func Writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

// UnwrapAvroUnion unwraps avro wrapped maps
func UnwrapAvroUnion(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	// v, r, s
	unwrapType(data, vLens, "bytes")
	unwrapType(data, rLens, "bytes")
	unwrapType(data, sLens, "bytes")

	// to, from
	unwrapType(data, toLens, "string")
	unwrapType(data, fromLens, "string")

	// withdrawalsRoot
	unwrapType(data, withdrawalsRootLens, "string")

	// withdrawals, uncles
	unwrapType(data, withdrawalsLens, "array")
	unwrapType(data, uncleLens, "array")

	// blobs
	unwrapType(data, blobGasUsedLens, "int")
	unwrapType(data, excessBlobGasLens, "int")
	unwrapType(data, parentBeaconRootLens, "string")
	unwrapType(data, blobFeeCapLens, "bytes")
	unwrapType(data, blobHashesLens, "array")
	unwrapType(data, blobGasLens, "int")

	unwrapType(data, blobTxSidecarsLens, "array")
	unwrapType(data, blobsLens, "string")
	unwrapType(data, commitmentsLens, "string")
	unwrapType(data, proofsLens, "string")

	// EIP-7685 (EL-CL tx)
	unwrapType(data, requestsHashLens, "string")

	// EIP-7702 (set EOA)
	unwrapType(data, dataTxLens, "bytes")
	unwrapType(data, authListLens, "array")
	unwrapType(data, chainIdLens, "int")
	unwrapType(data, addressLens, "string")
	unwrapType(data, nonceLens, "long")
	unwrapType(data, yParityLens, "bytes")
	unwrapType(data, rTxLens, "bytes")
	unwrapType(data, sTxLens, "bytes")

	return data
}

// MapToAvroUnion converts the several field in the replica map to an Avro Union type allowing <nil>
func MapToAvroUnion(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	// v,r,s
	wrapType(data, vLens, "bytes")
	wrapType(data, rLens, "bytes")
	wrapType(data, sLens, "bytes")

	// to, from
	wrapType(data, toLens, "string")
	wrapType(data, fromLens, "string")

	// withdrawalsRoot
	wrapType(data, withdrawalsRootLens, "string")

	// withdrawals, uncles
	wrapType(data, withdrawalsLens, "array")
	wrapType(data, uncleLens, "array")

	// blobs
	wrapType(data, blobGasUsedLens, "int")
	wrapType(data, excessBlobGasLens, "int")
	wrapType(data, parentBeaconRootLens, "string")
	wrapType(data, blobFeeCapLens, "bytes")
	wrapType(data, blobHashesLens, "array")
	wrapType(data, blobGasLens, "int")

	wrapType(data, blobTxSidecarsLens, "array")
	wrapType(data, blobsLens, "string")
	wrapType(data, commitmentsLens, "string")
	wrapType(data, proofsLens, "string")

	// EIP-7685 (EL-CL tx)
	wrapType(data, requestsHashLens, "string")

	// EIP-7702 (set EOA)
	wrapType(data, dataTxLens, "bytes")
	wrapType(data, authListLens, "array")
	wrapType(data, chainIdLens, "int")
	wrapType(data, addressLens, "string")
	wrapType(data, nonceLens, "long")
	wrapType(data, yParityLens, "bytes")
	wrapType(data, rTxLens, "bytes")
	wrapType(data, sTxLens, "bytes")

	return data
}

// Version Provides version info on bsp agent binary
func Version() {
	fmt.Println(clientIdentifier)
	fmt.Println("Bsp Agent Version:", BspAgentVersion)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
}
