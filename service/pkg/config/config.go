package config

import (
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	SSLMode  string
}

var Config DBConfig

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	Config = DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Name:     getEnv("DB_NAME", "postgres"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		Port:     getEnv("DB_PORT", "5432"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	return nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
