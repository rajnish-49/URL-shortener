package inmemory

import (
	"context"
	"errors"
	"sync"

	"url-shortener/internal/domain"
)

var ErrNotFound = errors.New("url not found")

type URLRepository struct {
	mu   sync.RWMutex
	urls map[string]domain.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		urls: make(map[string]domain.URL),
	}
}

func (r *URLRepository) Create(ctx context.Context, url domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.urls[url.ShortCode] = url
	return nil
}

func (r *URLRepository) GetByCode(ctx context.Context, code string) (domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.urls[code]
	if !ok {
		return domain.URL{}, ErrNotFound
	}

	return url, nil
}

func (r *URLRepository) Exists(ctx context.Context, code string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.urls[code]
	return ok, nil
}
