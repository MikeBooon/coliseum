// Package config implements interface for getting config data
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBConn string
	EncKey string
}

func GetEnv(useDotEnv bool) *EnvConfig {
	if useDotEnv {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	c := EnvConfig{
		DBConn: getEnvVarOrThrow("DB_CONNECTION"),
		EncKey: getEnvVarOrThrow("ENC_KEY"),
	}
	return &c
}

func getEnvVarOrThrow(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("missing env var: " + key)
	}
	return v
}
