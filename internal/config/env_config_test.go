package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvConfig(t *testing.T) {
	t.Run("successful config loading with defaults", func(t *testing.T) {
		// Clear any existing env vars that might interfere
		os.Clearenv()

		cfg, err := loadEnvConfig()
		assert.NoError(t, err)
		assert.NotNil(t, cfg)

		// Check default values
		assert.Equal(t, "", cfg.RedisConfig.Password)
		assert.Equal(t, "http://127.0.0.1:3000", cfg.IpfsConfig.IpfsPinnerServer)
	})

	t.Run("successful config loading with custom values", func(t *testing.T) {
		// Clear any existing env vars
		os.Clearenv()

		// Set test environment variables
		envVars := map[string]string{
			"REDIS_PWD":          "testpassword",
			"MB_RPC_URL":         "https://test.rpc.url",
			"MB_PRIVATE_KEY":     "testkey123",
			"MB_KEYSTORE_PATH":   "/test/path",
			"MB_KEYSTORE_PWD":    "keystorepass",
			"IFPS_PINNER_SERVER": "https://custom.ipfs.server",
		}

		for k, v := range envVars {
			os.Setenv(k, v)
		}

		cfg, err := loadEnvConfig()
		assert.NoError(t, err)
		assert.NotNil(t, cfg)

		// Verify all values were loaded correctly
		assert.Equal(t, "testpassword", cfg.RedisConfig.Password)
		assert.Equal(t, "https://test.rpc.url", cfg.EthConfig.RPCURL)
		assert.Equal(t, "testkey123", cfg.EthConfig.PrivateKey)
		assert.Equal(t, "/test/path", cfg.EthConfig.KeystorePath)
		assert.Equal(t, "keystorepass", cfg.EthConfig.KeyStorePwd)
		assert.Equal(t, "https://custom.ipfs.server", cfg.IpfsConfig.IpfsPinnerServer)
	})

	t.Run("invalid environment variable type", func(t *testing.T) {
		// Clear any existing env vars
		os.Clearenv()

		// Set an invalid environment variable (if we had any integer fields)
		// This test is included as an example, but with the current config
		// structure it's hard to trigger a real parsing error since all fields
		// are strings

		cfg, err := loadEnvConfig()
		assert.NoError(t, err) // Should still succeed as we have no non-string fields
		assert.NotNil(t, cfg)
	})
}
