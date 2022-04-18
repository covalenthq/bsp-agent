// Package config contains all the config functions that cannot be used in the cli interface
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// RedisConfig is set of all config that relates to redis
type RedisConfig struct {
	Password string `envconfig:"REDIS_PWD" default:""`
}

// EthConfig is set of all config that relates to ethereum / ethereum like (EVM) networks
type EthConfig struct {
	RPCURL       string `envconfig:"MB_RPC_URL"`
	PrivateKey   string `envconfig:"MB_PRIVATE_KEY"`
	KeystorePath string `envconfig:"MB_KEYSTORE_PATH"`
	KeyStorePwd  string `envconfig:"MB_KEYSTORE_PWD"`
}

// IPFSConfig is set of all configs that relates to IPFS pinning
type IPFSConfig struct {
	ServiceToken string `envconfig:"IPFS_SERVICE_TOKEN"`
}

// Config is set of all config that relates to .envrc
type Config struct {
	IPFSConfig  IPFSConfig
	RedisConfig RedisConfig
	EthConfig   EthConfig
}

// LoadConfig loads the config from .envrc file
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
