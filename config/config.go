package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	OpenAIAPIKey string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Port:         getEnvOrDefault("PORT", "8080"),
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
