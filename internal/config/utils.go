package config

import (
	"log"
	"os"
	"strconv"
)

// LookupEnvOrString looks up a flag env that is a string
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

// LookupEnvOrInt looks up a flag env that is an integer
func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("unable to lookupEnvOrInt[%s]: %v", key, err)
		}

		return v
	}

	return defaultVal
}

// LookupEnvOrInt64 looks up a flag env that is an integer
func LookupEnvOrInt64(key string, defaultVal int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatalf("unable to lookupEnvOrInt[%s]: %v", key, err)
		}

		return v
	}

	return defaultVal
}
