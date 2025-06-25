package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectToRedis() {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatal("REDIS_URL environment variable is not set")
	}

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatal("Invalid REDIS_URL format:", err)
	}

	client := redis.NewClient(opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	Redis = client
	log.Println("Connected to Redis successfully")
}
