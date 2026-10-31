package config

import (
	"errors"
	"os"
)

type Config struct {
	HTTPPort    string
	DatabaseURL string
}

func Load() (*Config, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("не указана строка подключения к бд")
	}

	return &Config{
		HTTPPort:    getEnv("HTTP_PORT", "8080"),
		DatabaseURL: dbUrl,
	}, nil
}

func getEnv(key, defaultvalue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultvalue
	}

	return value
}
