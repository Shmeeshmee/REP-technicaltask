package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DriverName string
	User       string
	Password   string
	Host       string
	Port       string
	Name       string
}

type Config struct {
	Port       string
	DebugLevel string
	DB         DBConfig
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	return &Config{
		Port:       os.Getenv("PORT"),
		DebugLevel: os.Getenv("DEBUG_LEVEL"),
		DB: DBConfig{
			DriverName: os.Getenv("DB_DRIVER_NAME"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			Name:       os.Getenv("DB_NAME"),
		},
	}
}
