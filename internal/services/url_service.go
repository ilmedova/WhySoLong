package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ilmedova/go-url-shortener/internal/models"
	"github.com/ilmedova/go-url-shortener/internal/repositories"
)

type URLService struct {
	Repo  *repositories.URLRepository
	Cache *repositories.CacheRepository
}

func NewURLService(repo *repositories.URLRepository, cache *repositories.CacheRepository) *URLService {
	return &URLService{Repo: repo, Cache: cache}
}

func (s *URLService) ShortenURL(ctx context.Context, original string) (string, error) {
	short := uuid.New().String()[:8] // Generate a unique 8-char string
	url := models.URL{ID: uuid.New().String(), Original: original, Shortened: short}

	err := s.Repo.SaveURL(ctx, url)
	if err != nil {
		return "", err
	}
	_ = s.Cache.SaveShortURL(ctx, short, original)
	return short, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, short string) (string, error) {
	original, err := s.Cache.GetOriginalURL(ctx, short)
	if err == nil {
		return original, nil
	}

	url, err := s.Repo.GetOriginalURL(ctx, short)
	if err != nil {
		return "", errors.New("URL not found")
	}
	return url.Original, nil
}
