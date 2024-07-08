package application

import (
	"context"
	"fmt"
	"os"

	"github.com/amaterasutears/url-shortener/internal/handler"
	links "github.com/amaterasutears/url-shortener/internal/repository/mongo"
	cache "github.com/amaterasutears/url-shortener/internal/repository/redis"
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

	linksCollection := mongoClient.Database("url-shortener").Collection("links")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	linksRepository := links.New(linksCollection)
	cacheLinksRepository := cache.New(redisClient)

	linksService := service.New(linksRepository, cacheLinksRepository)

	validate := validator.New(validator.WithRequiredStructEnabled())

	app := fiber.New()

	app.Get("/a", handler.Shorten(linksService, validate))
	app.Get("/s/:code", handler.Redirect(linksService, validate))

	err = app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}
