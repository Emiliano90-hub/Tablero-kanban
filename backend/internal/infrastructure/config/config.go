package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PostgreConfig PostgreConfig
}

type PostgreConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

func GetConfig() Config {
	loadEnv()

	return Config{
		PostgreConfig: PostgreConfig{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
			DB_PORT:     os.Getenv("DB_PORT"),
		},
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %v", err))
	}
}
