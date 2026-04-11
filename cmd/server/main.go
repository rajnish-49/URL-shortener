package main

import (
	"log"

	"url-shortener/internal/app"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.Load()
	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
