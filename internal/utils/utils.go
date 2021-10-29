package utils

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func NewRedisClient(redisConnection string) (*redis.Client, string, string, error) {

	u, err := url.Parse(redisConnection)
	if err != nil {
		log.Fatalf("we have an error here 1", err)
	}

	pass, _ := u.User.Password()
	dbString := strings.Replace(u.Path, "/", "", -1)
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatalf("we have an error here 2", err)
	}

	dbInt, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: pass,
		DB:       dbInt, // use default DB
	})

	streamKey := m["topic"][0]
	consumerGroup := u.Fragment

	_, err = redisClient.Ping().Result()
	return redisClient, streamKey, consumerGroup, err
}

func NewEthClient(address string) (*ethclient.Client, error) {

	ethClient, err := ethclient.Dial(address)
	if err != nil {
		log.Error("error in getting eth client: ", err.Error())
	}

	return ethClient, nil
}

func NewStorageClient(config *config.GcpConfig) (*storage.Client, error) {

	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(config.ServiceAccount))
	if err != nil {
		return nil, err
	}

	return storageClient, nil
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func AckStreamSegment(config *config.Config, redisClient *redis.Client, streamKey, consumerGroup string, streamIDs []string) error {

	if len(streamIDs) == int(config.GeneralConfig.SegmentLength) {
		for _, streamID := range streamIDs {
			redisClient.XAck(streamKey, consumerGroup, streamID)
		}
		return nil
	} else {
		return fmt.Errorf("failed to match streamIDs length to segment length config")
	}

}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func GetConfig(fs *flag.FlagSet) []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *flag.Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}
