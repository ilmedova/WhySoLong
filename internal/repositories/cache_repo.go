package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type CacheRepository struct {
	Client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{Client: client}
}

func (r *CacheRepository) SaveShortURL(ctx context.Context, short, original string) error {
	return r.Client.Set(ctx, short, original, 0).Err()
}

func (r *CacheRepository) GetOriginalURL(ctx context.Context, short string) (string, error) {
	return r.Client.Get(ctx, short).Result()
}
