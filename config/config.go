package config

import (
	"log"
	"os"
)

type Config struct {
	Port  string
	DbURl string
}

func LoadConfig() *Config {
	config := &Config{}
	mustMapEnv(&config.Port, "PORT")
	mustMapEnv(&config.DbURl, "DB_URL")
	return config
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Panicf("Environment variable %s not set", envKey)
	}
	*target = v
}
