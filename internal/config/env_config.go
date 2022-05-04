// Package config contains all the config functions that cannot be used in the cli interface
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// RedisEnvConfig is set of all config that relates to redis
type RedisEnvConfig struct {
	Password string `envconfig:"REDIS_PWD" default:""`
}

// EthEnvConfig is set of all config that relates to ethereum / ethereum like (EVM) networks
type EthEnvConfig struct {
	RPCURL       string `envconfig:"MB_RPC_URL"`
	PrivateKey   string `envconfig:"MB_PRIVATE_KEY"`
	KeystorePath string `envconfig:"MB_KEYSTORE_PATH"`
	KeyStorePwd  string `envconfig:"MB_KEYSTORE_PWD"`
}

// IpfsEnvConfig is set of all configs that relates to IPFS pinning (passed via env)
type IpfsEnvConfig struct {
	ServiceToken string `envconfig:"IPFS_SERVICE_TOKEN"`
}

// EnvConfig is set of all EnvConfig that relates to .envrc
type EnvConfig struct {
	IpfsConfig  IpfsEnvConfig
	RedisConfig RedisEnvConfig
	EthConfig   EthEnvConfig
}

func loadEnvConfig() (*EnvConfig, error) {
	cfg := &EnvConfig{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
