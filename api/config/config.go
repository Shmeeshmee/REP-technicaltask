package config

import (
	"os"
)

type RedisConfig struct {
	Driver string
	Port   string
}

type Config struct {
	Port       string
	DebugLevel string
	Redis      RedisConfig
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("PORT"),
		DebugLevel: os.Getenv("DEBUG_LEVEL"),
		Redis: RedisConfig{
			Port:   os.Getenv("REDIS_PORT"),
			Driver: os.Getenv("REDIS_DRIVER"),
		},
	}
}
