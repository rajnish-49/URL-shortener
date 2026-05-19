package service

import (
	"context"
	"errors"
	"net/url"
	"time"

	"url-shortener/internal/domain"
	"url-shortener/internal/repository"
	"url-shortener/internal/shortener"
)

var ErrInvalidURL = errors.New("invalid url")

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) Create(ctx context.Context, longURL string) (domain.URL, error) {
	parsedURL, err := url.ParseRequestURI(longURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return domain.URL{}, ErrInvalidURL
	}

	for {
		code, err := shortener.Generate(6)
		if err != nil {
			return domain.URL{}, err
		}

		exists, err := s.repo.Exists(ctx, code)
		if err != nil {
			return domain.URL{}, err
		}

		if exists {
			continue
		}

		urlModel := domain.URL{
			ShortCode:  code,
			LongURL:    longURL,
			CreatedAt:  time.Now(),
			ClickCount: 0,
		}

		if err := s.repo.Create(ctx, urlModel); err != nil {
			return domain.URL{}, err
		}

		return urlModel, nil
	}
}

func (s *URLService) GetByCode(ctx context.Context, code string) (domain.URL, error) {
	return s.repo.GetByCode(ctx, code)
}

func (s *URLService) Resolve(ctx context.Context, code string) (domain.URL, error) {
	url, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return domain.URL{}, err
	}

	if err := s.repo.IncrementClickCount(ctx, code); err != nil {
		return domain.URL{}, err
	}

	return url, nil
}

