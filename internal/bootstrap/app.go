package bootstrap

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/infra"
	"zhangdx.cn/blog-server-stats/internal/readcount"
	"zhangdx.cn/blog-server-stats/internal/router"
)

func NewApp() {
	redisClient := NewRedisClient()
	redisOperator := infra.NewRedisOperator(redisClient)
	readCountService := readcount.NewReadCountService(redisOperator)
	readCountHandler := readcount.NewReadCountHandler(readCountService)
	r := gin.Default()
	router.Init(r, readCountHandler)
	r.Run(":8081")
}
