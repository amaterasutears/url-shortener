package cache

import (
	"context"
	"time"

	"github.com/amaterasutears/url-shortener/internal/service"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

var _ service.LinksRepository = (*RedisRepository)(nil)

func New(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) Put(code, original string) error {
	err := r.client.Set(context.Background(), code, original, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) FindOne(code string) (string, error) {
	original, err := r.client.Get(context.Background(), code).Result()
	if err != nil {
		return "", err
	}

	return original, nil
}
