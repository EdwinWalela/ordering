package config

import (
	"log"
	"os"
)

type Config struct {
	Port   string
	DbURl  string
	ATKey  string
	ATuser string
	ATenv  string
}

func LoadConfig() *Config {
	config := &Config{}
	mustMapEnv(&config.Port, "PORT")
	mustMapEnv(&config.DbURl, "DB_URL")
	mustMapEnv(&config.ATKey, "AT_KEY")
	mustMapEnv(&config.ATuser, "AT_USER")
	mustMapEnv(&config.ATenv, "AT_ENV")
	return config
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Panicf("Environment variable %s not set", envKey)
	}
	*target = v
}
