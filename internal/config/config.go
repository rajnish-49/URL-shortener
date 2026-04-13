package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	BaseURL     string
	DatabaseURL string
}

func Load() Config {
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:" + port
	}

	databaseURL := os.Getenv("DATABASE_URL")

	return Config{
		Port:        port,
		BaseURL:     baseURL,
		DatabaseURL: databaseURL,
	}
}
