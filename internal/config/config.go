package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GcpConfig struct {
	ServiceAccount string `envconfig:"GCP_SERVICE_ACCOUNT"`
	ProjectId      string `envconfig:"GCP_PROJECT_ID"`
	ResultBucket   string `envconfig:"GCP_RESULT_BUCKET"`
	SpecimenBucket string `envconfig:"GCP_SPECIMEN_BUCKET"`
}

type RedisConfig struct {
	Address  string `envconfig:"REDIS_ADDRESS"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
	Key      string `envconfig:"REDIS_STREAM_KEY" default:"replication"`
	Group    string `envconfig:"REDIS_CONSUMER_GROUP" default:"replicate"`
}

type EthConfig struct {
	ProofClient  string `envconfig:"ETH_PROOF_CLIENT"`
	SourceClient string `envconfig:"ETH_SOURCE_CLIENT"`
	Key          string `envconfig:"ETH_PRIVATE_KEY"`
	Contract     string `envconfig:"ETH_PROOF_CONTRACT"`
	ChainId      uint64 `envconfig:"ETH_CHAIN_ID" default:"5"`
	Keystore     string `envconfig:"ETH_KEYSTORE_PATH"`
	Password     string `envconfig:"ETH_KEYSTORE_PWD"`
}

type GeneralConfig struct {
	ConsumeEvents int64 `envconfig:"CONSUME_EVENTS" default:"10"`
}

type Config struct {
	GcpConfig     GcpConfig
	RedisConfig   RedisConfig
	GeneralConfig GeneralConfig
	EthConfig     EthConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
