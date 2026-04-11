package app

import (
	"log"
	"net/http"

	"url-shortener/internal/config"
	handlerhttp "url-shortener/internal/handler/http"
	"url-shortener/internal/repository/inmemory"
	"url-shortener/internal/service"
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

	urlRepo := inmemory.NewURLRepository()
	urlService := service.NewURLService(urlRepo)
	urlHandler := handlerhttp.NewURLHandler(urlService, a.cfg.BaseURL)

	mux.HandleFunc("/health", handlerhttp.Health)
	mux.HandleFunc("/shorten", urlHandler.Shorten)

	log.Printf("server running on :%s", a.cfg.Port)

	return http.ListenAndServe(":"+a.cfg.Port, mux)
}