package main

import (
	"blog-server-stats/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)
	r.Run(":8081")
}
