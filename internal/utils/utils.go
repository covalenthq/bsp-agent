package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func NewRedisClient(config *config.RedisConfig) (*redis.Client, error) {

	u, err := url.Parse(config.Url)
	if err != nil {
		panic(err)
	}

	pass, _ := u.User.Password()
	dbString := strings.Replace(u.Path, "/", "", -1)
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
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

	config.Key = m["topic"][0]
	config.Group = u.Fragment

	_, err = redisClient.Ping().Result()
	return redisClient, err
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

func AckStreamSegment(config *config.Config, redisClient *redis.Client, streamIDs []string) error {

	if len(streamIDs) == int(config.GeneralConfig.SegmentLength) {
		for _, streamID := range streamIDs {
			redisClient.XAck(config.RedisConfig.Key, config.RedisConfig.Group, streamID)
		}
		return nil
	} else {
		return fmt.Errorf("failed to match streamIDs length to segment length config")
	}

}
