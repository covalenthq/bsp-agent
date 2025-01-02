package config

import (
	"os"
	"testing"
)

func TestLookupEnvOrString(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		defaultVal string
		envValue   string
		want       string
	}{
		{
			name:       "returns default when env not set",
			key:        "TEST_STRING_KEY",
			defaultVal: "default",
			envValue:   "",
			want:       "default",
		},
		{
			name:       "returns env value when set",
			key:        "TEST_STRING_KEY",
			defaultVal: "default",
			envValue:   "env_value",
			want:       "env_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			if got := LookupEnvOrString(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("LookupEnvOrString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupEnvOrInt(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		defaultVal int
		envValue   string
		want       int
		wantPanic  bool
	}{
		{
			name:       "returns default when env not set",
			key:        "TEST_INT_KEY",
			defaultVal: 42,
			envValue:   "",
			want:       42,
		},
		{
			name:       "returns env value when set",
			key:        "TEST_INT_KEY",
			defaultVal: 42,
			envValue:   "123",
			want:       123,
		},
		{
			name:       "panics on invalid integer",
			key:        "TEST_INT_KEY",
			defaultVal: 42,
			envValue:   "not_an_int",
			wantPanic:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("LookupEnvOrInt() should have panicked")
					}
				}()
			}

			if got := LookupEnvOrInt(tt.key, tt.defaultVal); !tt.wantPanic && got != tt.want {
				t.Errorf("LookupEnvOrInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupEnvOrInt64(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		defaultVal int64
		envValue   string
		want       int64
		wantPanic  bool
	}{
		{
			name:       "returns default when env not set",
			key:        "TEST_INT64_KEY",
			defaultVal: 42,
			envValue:   "",
			want:       42,
		},
		{
			name:       "returns env value when set",
			key:        "TEST_INT64_KEY",
			defaultVal: 42,
			envValue:   "9223372036854775807", // max int64
			want:       9223372036854775807,
		},
		{
			name:       "panics on invalid integer",
			key:        "TEST_INT64_KEY",
			defaultVal: 42,
			envValue:   "not_an_int",
			wantPanic:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("LookupEnvOrInt64() should have panicked")
					}
				}()
			}

			if got := LookupEnvOrInt64(tt.key, tt.defaultVal); !tt.wantPanic && got != tt.want {
				t.Errorf("LookupEnvOrInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
