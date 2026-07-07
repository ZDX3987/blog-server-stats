package router

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/readcount/handler"
)

func Init(r *gin.Engine) {
	stat := r.Group("/api/stat")
	stat.POST("/read", handler.Handler)
}
