package app

import (
	"context"
	"log"
	"net/http"

	"url-shortener/internal/config"
	"url-shortener/internal/database"
	handlerhttp "url-shortener/internal/handler/http"
	"url-shortener/internal/repository/postgres"
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
	ctx := context.Background()

	pool, err := database.NewPostgresPool(ctx, a.cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	mux := http.NewServeMux()

	urlRepo := postgres.NewURLRepository(pool)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handlerhttp.NewURLHandler(urlService, a.cfg.BaseURL)

	mux.HandleFunc("/health", handlerhttp.Health)
	mux.HandleFunc("/shorten", urlHandler.Shorten)
	mux.HandleFunc("/", urlHandler.Redirect)

	log.Printf("server running on :%s", a.cfg.Port)

	return http.ListenAndServe(":"+a.cfg.Port, mux)
}
