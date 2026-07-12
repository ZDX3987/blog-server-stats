package bootstrap

import "github.com/redis/go-redis/v9"

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "47.113.97.58:6379",
		Password: "",
		DB:       1,
	})
	return client
}
