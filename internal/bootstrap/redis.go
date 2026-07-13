package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"zhangdx.cn/blog-server-stats/internal/config"
)

func NewRedisClient(c config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	return client
}
