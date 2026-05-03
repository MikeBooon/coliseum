// Package config implements interface for getting config data
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Domain           string
	DBConn           string
	EncKey           string
	CasdoorDomain    string
	CasdoorAppID     string
	CasdoorAppSecret string
}

func GetEnv(useDotEnv bool) *EnvConfig {
	if useDotEnv {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	c := EnvConfig{
		Domain:           getEnvVarOrThrow("DOMAIN"),
		DBConn:           getEnvVarOrThrow("DB_CONNECTION"),
		EncKey:           getEnvVarOrThrow("ENC_KEY"),
		CasdoorDomain:    getEnvVarOrThrow("CASDOOR_DOMAIN"),
		CasdoorAppID:     getEnvVarOrThrow("CASDOOR_APP_ID"),
		CasdoorAppSecret: getEnvVarOrThrow("CASDOOR_APP_SECRET"),
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
