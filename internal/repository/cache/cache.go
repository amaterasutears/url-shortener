package cache

import (
	"context"
	"time"

	"github.com/amaterasutears/url-shortener/internal/entity"
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
	err := r.client.Set(context.TODO(), code, original, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindOne(code string) (*entity.Link, error) {
	var link entity.Link

	err := r.client.Get(context.TODO(), code).Scan(&link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
