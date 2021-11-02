package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type EthConfig struct {
	PrivateKey   string `envconfig:"ETH_PRIVATE_KEY"`
	KeystorePath string `envconfig:"ETH_KEYSTORE_PATH"`
	KeyStorePwd  string `envconfig:"ETH_KEYSTORE_PWD"`
}
type Config struct {
	EthConfig EthConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
