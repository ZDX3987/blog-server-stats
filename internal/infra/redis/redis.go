package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOperator struct {
	client *redis.Client
}

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{client: client}
}

func (operator *RedisOperator) Get(key string) string {
	ctx := context.Background()
	val, err := operator.client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error getting key %s: %v", key, err)
	}
	log.Printf("get redis value: %s\n", val)
	return val
}

func (operator *RedisOperator) Set(key, value string) {
	ctx := context.Background()
	val, err := operator.client.Set(ctx, key, value, 20*time.Second).Result()
	if err != nil {
		log.Fatalf("Error set key %s: %v", key, err)
	}
	log.Printf("set redis value: %s\n", val)
}
