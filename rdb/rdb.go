package rdb

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb = initClient()

func initClient() redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.113.97.58:6379",
		Password: "",
		DB:       1,
	})
	return *rdb
}

func Get(key string) string {
	ctx := context.Background()
	val, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error getting key %s: %v", key, err)
	}
	log.Printf("get redis value: %s\n", val)
	return val
}

func Set(key, value string) {
	ctx := context.Background()
	val, err := Rdb.Set(ctx, key, value, 20*time.Second).Result()
	if err != nil {
		log.Fatalf("Error set key %s: %v", key, err)
	}
	log.Printf("set redis value: %s\n", val)
}
