package application

import (
	"context"
	"fmt"
	"os"

	"github.com/amaterasutears/url-shortener/internal/handler"
	mongo_repo "github.com/amaterasutears/url-shortener/internal/repository/mongo"
	redis_repo "github.com/amaterasutears/url-shortener/internal/repository/redis"
	"github.com/amaterasutears/url-shortener/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() error {
	mongoUsername := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongo:27017", mongoUsername, mongoPassword)

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	defer func() error {
		err = mongoClient.Disconnect(context.Background())
		if err != nil {
			return err
		}

		return nil
	}()

	mongoDatabase := os.Getenv("MONGO_DB")
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	linksCollection := mongoClient.Database(mongoDatabase).Collection(mongoCollection)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	defer func() error {
		err = redisClient.Close()
		if err != nil {
			return err
		}

		return nil
	}()

	mongoRepository := mongo_repo.New(linksCollection)
	redisRepository := redis_repo.New(redisClient)

	shortenerService := service.New(mongoRepository, redisRepository)

	validate := validator.New(validator.WithRequiredStructEnabled())

	app := fiber.New()

	app.Get("/a", handler.Shorten(shortenerService, validate))
	app.Get("/s/:code", handler.Redirect(shortenerService, validate))

	err = app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}
