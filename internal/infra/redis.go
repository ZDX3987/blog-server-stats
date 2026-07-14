package infra

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

func (operator *RedisOperator) Get(ctx context.Context, key string) (string, error) {
	val, err := operator.client.Get(ctx, key).Result()
	log.Printf("get redis value: %s\n", val)
	return val, err
}

func (operator *RedisOperator) Set(ctx context.Context, key, value string, expire time.Duration) {
	val, err := operator.client.Set(ctx, key, value, expire).Result()
	if err != nil {
		log.Fatalf("Error set key %s: %v", key, err)
	}
	log.Printf("set redis value: %s\n", val)
}

func (operator *RedisOperator) SetNx(ctx context.Context, key, value string, expire time.Duration) (bool, error) {
	val, err := operator.client.SetNX(ctx, key, value, expire).Result()
	log.Printf("set redis key: %s, value: %v\n", key, val)
	return val, err
}

func (operator *RedisOperator) AddSet(ctx context.Context, key, value string, expire time.Duration) (bool, error) {
	val, err := operator.client.SAdd(ctx, key, value, expire).Result()
	return val == 1, err
}

func (operator *RedisOperator) ListSet(ctx context.Context, key string) ([]string, error) {
	return operator.client.SMembers(ctx, key).Result()
}

func (operator *RedisOperator) Incr(ctx context.Context, key string) error {
	val, err := operator.client.Incr(ctx, key).Result()
	log.Printf("incr redis key: %s, value: %v\n", key, val)
	return err
}

func (operator *RedisOperator) Del(ctx context.Context, key string) error {
	val, err := operator.client.Del(ctx, key).Result()
	log.Printf("del redis key: %s, value: %v\n", key, val)
	return err
}
