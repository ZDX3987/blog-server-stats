package router

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/middleware"
	"zhangdx.cn/blog-server-stats/internal/readcount"
)

func Init(r *gin.Engine, readCountHandler *readcount.Handler) {
	stat := r.Group("/api/stats")
	stat.Use(middleware.VisitorMiddleware(), middleware.ErrorMiddleware())
	stat.POST("/read", readCountHandler.SubmitReadRequest)
	stat.POST("/page-view")
}
