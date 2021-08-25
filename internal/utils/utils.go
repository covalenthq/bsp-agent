package utils

import (
	"github.com/go-redis/redis/v7"
)

func NewRedisClient(Address, Password string, DB int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     Address,
		Password: Password,
		DB:       DB, // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
