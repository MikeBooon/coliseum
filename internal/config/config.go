// Package config implements interface for getting config data
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Domain        string
	DBConn        string
	EncKey        string
	ZitadelDomain string
	ZitadelClient string
}

func GetEnv(useDotEnv bool) *EnvConfig {
	if useDotEnv {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	c := EnvConfig{
		Domain:        getEnvVarOrThrow("DOMAIN"),
		DBConn:        getEnvVarOrThrow("DB_CONNECTION"),
		EncKey:        getEnvVarOrThrow("ENC_KEY"),
		ZitadelDomain: getEnvVarOrThrow("ZITADEL_DOMAIN"),
		ZitadelClient: getEnvVarOrThrow("ZITADEL_CLIENT"),
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
