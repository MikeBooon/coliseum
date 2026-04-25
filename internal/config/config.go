// Package config implements interface for getting config data
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConn string
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
