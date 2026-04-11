package config

import "os"

type Config struct{
	Port string 
	BaseURL string 
	DatabaseURL string 
}

func Load() Config{
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL== ""{
		baseURL = "http://localhost:" + port
	}

	databaseURL := os.Getenv("DATABASE_URL")

	return Config{
		Port:        port,
		BaseURL:     baseURL,
		DatabaseURL: databaseURL,
	}
}