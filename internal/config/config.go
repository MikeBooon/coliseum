// Package config implements interface for getting config data
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Domain             string
	DBConn             string
	EncKey             string
	CasdoorDomain      string
	CasdoorAppClientID string
	CasdoorAppSecret   string
	CasdoorAppCert     string
	CasdoorAppName     string
	CasdoorAdminOrg    string
}

func GetEnv(useDotEnv bool) *EnvConfig {
	if useDotEnv {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	c := EnvConfig{
		Domain: getEnvVarOrThrow("DOMAIN"),
		DBConn: getEnvVarOrThrow("DB_CONNECTION"),
		EncKey: getEnvVarOrThrow("ENC_KEY"),
	}
	return &c
}

func getEnvVarOrThrow(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing env var: %s", key)
	}
	return v
}

func getFileContentsOrThrow(path string) string {
	certBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load file '%s' because: %v", path, err)
	}
	return string(certBytes)
}
