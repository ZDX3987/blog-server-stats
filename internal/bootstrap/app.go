package bootstrap

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/infra/redis"
	"zhangdx.cn/blog-server-stats/internal/readcount/handler"
	"zhangdx.cn/blog-server-stats/internal/readcount/service"
	"zhangdx.cn/blog-server-stats/internal/router"
)

func NewApp() {
	redisClient := NewRedisClient()
	redisOperator := redis.NewRedisOperator(redisClient)
	readCountService := service.NewReadCountService(redisOperator)
	readCountHandler := handler.NewReadCountHandler(readCountService)
	r := gin.Default()
	router.Init(r, readCountHandler)
	r.Run(":8081")
}
