package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GeneralConfig struct {
	GcpServiceAccount string `envconfig:"GCP_SERVICE_ACCOUNT"`
	GcpProjectId      string `envconfig:"GCP_PROJECT_ID"`
	GcpBucketName     string `envconfig:"GCP_BUCKET_NAME"`
}

type Config struct {
	GeneralConfig GeneralConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
