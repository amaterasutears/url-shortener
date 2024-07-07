package links

import (
	"context"

	"github.com/amaterasutears/url-shortener/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{
		collection: collection,
	}
}

func (r *Repository) Put(code, original string) error {
	_, err := r.collection.InsertOne(
		context.Background(),
		bson.M{"code": code, "original": original},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindOne(code string) (string, error) {
	filter := bson.M{"code": code}

	var link entity.Link

	err := r.collection.FindOne(context.Background(), filter).Decode(&link)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		return "", err
	}

	return link.Original, nil
}