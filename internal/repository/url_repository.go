package repository

import (
	"context"
	
	"url-shortener/internal/domain"
)

type URLRepository interface {
	Create(ctx context.Context, url domain.URL) error
	GetByCode(ctx context.Context, code string) (domain.URL, error)
	Exists(ctx context.Context, code string) (bool, error)
	IncrementClickCount(ctx context.Context , code string ) error
}
