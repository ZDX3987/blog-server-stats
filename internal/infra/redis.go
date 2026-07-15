package infra

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOperator struct {
	Client *redis.Client
}

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{Client: client}
}

func (operator *RedisOperator) Get(ctx context.Context, key string) (string, error) {
	val, err := operator.Client.Get(ctx, key).Result()
	log.Printf("get redis value: %s\n", val)
	return val, err
}

func (operator *RedisOperator) Set(ctx context.Context, key, value string, expire time.Duration) {
	val, err := operator.Client.Set(ctx, key, value, expire).Result()
	if err != nil {
		log.Fatalf("Error set key %s: %v", key, err)
	}
	log.Printf("set redis value: %s\n", val)
}

func (operator *RedisOperator) SetNx(ctx context.Context, key, value string, expire time.Duration) (bool, error) {
	val, err := operator.Client.SetNX(ctx, key, value, expire).Result()
	log.Printf("set redis key: %s, value: %v\n", key, val)
	return val, err
}

func (operator *RedisOperator) AddSet(ctx context.Context, key, value string, expire time.Duration) (bool, error) {
	val, err := operator.Client.SAdd(ctx, key, value, expire).Result()
	return val == 1, err
}

func (operator *RedisOperator) ListSet(ctx context.Context, key string) ([]string, error) {
	return operator.Client.SMembers(ctx, key).Result()
}

func (operator *RedisOperator) Incr(ctx context.Context, key string) error {
	val, err := operator.Client.Incr(ctx, key).Result()
	log.Printf("incr redis key: %s, value: %v\n", key, val)
	return err
}

func (operator *RedisOperator) Del(ctx context.Context, keys ...string) error {
	val, err := operator.Client.Del(ctx, keys...).Result()
	log.Printf("del redis key: %s, value: %v\n", keys, val)
	return err
}
