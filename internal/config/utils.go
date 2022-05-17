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
