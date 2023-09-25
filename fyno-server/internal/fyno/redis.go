package fyno

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func NewRedis(logger zerolog.Logger) (*redis.Client, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	logger.Info().Msg("connecting to Redis")
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	// Check Redis connection
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	logger.Info().Msg("connected to Redis")
	return redisClient, nil
}
