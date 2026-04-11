package app

import (
	"log"
	"net/http"

	"url-shortener/internal/config"
	handlerhttp "url-shortener/internal/handler/http"
)

type App struct {
	cfg config.Config
}

func New(cfg config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlerhttp.Health)

	log.Printf("server running on :%s", a.cfg.Port)

	return http.ListenAndServe(":"+a.cfg.Port, mux)
}
