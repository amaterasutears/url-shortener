package mongo

import (
	"context"

	"github.com/amaterasutears/url-shortener/internal/entity"
	"github.com/amaterasutears/url-shortener/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

var _ service.LinksRepository = (*MongoRepository)(nil)

func New(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		collection: collection,
	}
}

func (r *MongoRepository) Put(code, original string) error {
	_, err := r.collection.InsertOne(
		context.Background(),
		bson.M{"code": code, "original": original},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoRepository) FindOne(code string) (string, error) {
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
