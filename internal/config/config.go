// Package config implements interface for getting config data
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConn string
	EncKey string
}

func Get(useDotEnv bool) *Config {
	if useDotEnv {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	c := Config{
		DBConn: getEnvOrThrow("DB_CONNECTION"),
		EncKey: getEnvOrThrow("ENC_KEY"),
	}
	return &c
}

func getEnvOrThrow(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("missing env var: " + key)
	}
	return v
}
