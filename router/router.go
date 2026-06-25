package router

import (
	"zhangdx.cn/blog-server-stats/readCount"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	stat := r.Group("/api/stat")
	stat.POST("/read", readCount.Handler)
}
