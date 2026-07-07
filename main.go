package main

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/router"
)

func main() {
	r := gin.Default()
	router.Init(r)
	r.Run(":8081")
}
