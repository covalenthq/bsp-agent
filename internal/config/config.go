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

// EthConfig is set of all config that relates to ethereum
type EthConfig struct {
	CqtRPCURL    string `envconfig:"CQT_RPC_URL"`
	ExtRPCURL    string `envconfig:"EXT_RPC_URL"`
	PrivateKey   string `envconfig:"PRIVATE_KEY"`
	KeystorePath string `envconfig:"ETH_KEYSTORE_PATH"`
	KeyStorePwd  string `envconfig:"ETH_KEYSTORE_PWD"`
}

// Config is set of all config that relates to .envrc
type Config struct {
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
