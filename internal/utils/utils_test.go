package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	type testStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
		wantErr  bool
	}{
		{
			name: "valid struct",
			input: testStruct{
				Name:  "test",
				Value: 42,
			},
			expected: map[string]interface{}{
				"name":  "test",
				"value": float64(42), // JSON numbers are float64
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := utils.StructToMap(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)

				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestExpandPath(t *testing.T) {
	// Save original HOME env
	originalHome := os.Getenv("HOME")
	err := os.Setenv("HOME", originalHome)
	assert.NoError(t, err)
	defer func() {
		err := os.Setenv("HOME", originalHome)
		assert.NoError(t, err)
	}()

	// Set test HOME
	testHome := "/test/home"
	err = os.Setenv("HOME", testHome)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "tilde expansion",
			path:     "~/test",
			expected: filepath.Join(testHome, "test"),
		},
		{
			name:     "env var expansion",
			path:     "$HOME/test",
			expected: filepath.Join(testHome, "test"),
		},
		{
			name:     "clean path",
			path:     "/a/b/../c",
			expected: "/a/c",
		},
		{
			name:     "no expansion needed",
			path:     "/absolute/path",
			expected: "/absolute/path",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := utils.ExpandPath(testCase.path)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestGetLogLocationURL(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-log-*")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	tests := []struct {
		name    string
		logPath string
		wantErr bool
		setup   func() error
		cleanup func()
	}{
		{
			name:    "valid path",
			logPath: tempDir,
			wantErr: false,
		},
		{
			name:    "non-existent path",
			logPath: filepath.Join(tempDir, "new-dir"),
			wantErr: false,
		},
		{
			name:    "invalid path",
			logPath: "/invalid/path/with/special/chars/\x00",
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setup != nil {
				err := testCase.setup()
				assert.NoError(t, err)
			}
			if testCase.cleanup != nil {
				defer testCase.cleanup()
			}

			url, err := utils.GetLogLocationURL(testCase.logPath)
			if testCase.wantErr {
				assert.Error(t, err)

				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, url)
			assert.Equal(t, testCase.logPath, url.Path)
		})
	}
}

func TestWritable(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-writable-*")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	// Create a file in the temp directory
	testFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test"), 0600)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "writable directory",
			path:     tempDir,
			expected: true,
		},
		{
			name:     "writable file",
			path:     testFile,
			expected: true,
		},
		{
			name:     "non-existent path",
			path:     filepath.Join(tempDir, "non-existent"),
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := utils.Writable(testCase.path)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestVersion(t *testing.T) {
	// This is a simple test to ensure Version() doesn't panic
	assert.NotPanics(t, func() {
		utils.Version()
	})
}

func TestAckTrimStreamSegment(t *testing.T) {
	// Create a mock Redis client
	mockClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	tests := []struct {
		name          string
		redisClient   *redis.Client
		segmentLength int
		streamKey     string
		consumerGroup string
		streamIDs     []string
		wantErr       bool
	}{
		{
			name:          "mismatched lengths",
			redisClient:   mockClient,
			segmentLength: 2,
			streamKey:     "test-stream",
			consumerGroup: "test-group",
			streamIDs:     []string{"1"},
			wantErr:       true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			length, err := utils.AckTrimStreamSegment(testCase.redisClient, testCase.segmentLength, testCase.streamKey, testCase.consumerGroup, testCase.streamIDs)
			if testCase.wantErr {
				assert.Error(t, err)

				return
			}
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, length, int64(0))
		})
	}
}
