package utils

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/storage"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	"github.com/ubiq/go-ubiq/log"
	"google.golang.org/api/option"
)

func NewRedisClient(config *config.RedisConfig) (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB, // use default DB
	})

	_, err := redisClient.Ping().Result()
	return redisClient, err
}

func NewEthClient(address string) (*ethclient.Client, error) {

	ethClient, err := ethclient.Dial(address)
	if err != nil {
		log.Error("error in getting eth client: ", err.Error())
	}

	return ethClient, nil
}

func NewStorageCliemt(config *config.GcpConfig) (*storage.Client, error) {

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
