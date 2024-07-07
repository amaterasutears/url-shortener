package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amaterasutears/url-shortener/internal/handler"
	"github.com/amaterasutears/url-shortener/internal/repository/cache"
	"github.com/amaterasutears/url-shortener/internal/repository/link"
	"github.com/amaterasutears/url-shortener/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUsername := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongo:27017", mongoUsername, mongoPassword)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := mongoClient.Database("url-shortener").Collection("links")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	defer func() {
		err := redisClient.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	linkRepo := link.New(collection)
	cacheRepo := cache.New(redisClient)
	svc := service.New(linkRepo, cacheRepo)

	v := validator.New(validator.WithRequiredStructEnabled())

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/a", handler.Shorten(svc, v))
	app.Get("/s/:code", handler.Redirect(svc, v))

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
