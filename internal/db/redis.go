package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(ctx context.Context) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis")
	}
	return redisClient
}
