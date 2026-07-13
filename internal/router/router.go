package router

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/middleware"
	"zhangdx.cn/blog-server-stats/internal/readcount"
)

func Init(r *gin.Engine, readCountHandler *readcount.Handler) {
	stat := r.Group("/api/stat")
	stat.Use(middleware.VisitorMiddleware())
	stat.POST("/read", readCountHandler.SubmitReadRequest)
}
