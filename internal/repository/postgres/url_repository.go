package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"url-shortener/internal/domain"
)

var ErrNotFound = errors.New("url not found")

type URLRepository struct {
	pool *pgxpool.Pool
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepository {
	return &URLRepository{
		pool: pool,
	}
}

func (r *URLRepository) Create(ctx context.Context, url domain.URL) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO urls (short_code, long_url, created_at, expires_at, click_count)
		VALUES ($1, $2, $3, $4, $5)
	`, url.ShortCode, url.LongURL, url.CreatedAt, url.ExpiresAt, url.ClickCount)

	return err
}

func (r *URLRepository) GetByCode(ctx context.Context, code string) (domain.URL, error) {
	var url domain.URL

	err := r.pool.QueryRow(ctx, `
		SELECT id, short_code, long_url, created_at, expires_at, click_count
		FROM urls
		WHERE short_code = $1
	`, code).Scan(
		&url.ID,&url.ShortCode,&url.LongURL,&url.CreatedAt,&url.ExpiresAt, &url.ClickCount,
	)
	if err != nil {
		return domain.URL{}, ErrNotFound
	}

	return url, nil
}

func (r *URLRepository) Exists(ctx context.Context, code string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM urls
			WHERE short_code = $1
		)
	`, code).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
