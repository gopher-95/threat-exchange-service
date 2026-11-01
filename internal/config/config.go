package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	HTTPPort   string
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	pass := getEnv("DB_PASSWORD", "")
	if pass == "" {
		return nil, fmt.Errorf("empty password")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: pass,
		DBName:     getEnv("DB_NAME", "name"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		HTTPPort:   getEnv("HTTP_PORT", "8080"),
	}, nil
}

func getEnv(key, defaultvalue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultvalue
	}

	return value
}
