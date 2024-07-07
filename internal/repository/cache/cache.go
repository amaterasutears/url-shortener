package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func New(client *redis.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Put(code, original string) error {
	err := r.client.Set(context.Background(), code, original, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindOne(code string) (string, error) {
	original, err := r.client.Get(context.Background(), code).Result()
	if err != nil {
		return "", err
	}

	return original, nil
}
