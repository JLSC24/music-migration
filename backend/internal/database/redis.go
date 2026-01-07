package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// ConnectRedis establishes connection to Redis
func ConnectRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("⚠️  Redis connection failed: %v (cache disabled)", err)
		return client
	}

	log.Println("✅ Redis connected successfully")
	return client
}

// CacheKey generates a consistent cache key
func CacheKey(prefix, id string) string {
	return fmt.Sprintf("%s:%s", prefix, id)
}
