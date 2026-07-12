package router

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/readcount"
)

func Init(r *gin.Engine, readCountHandler *readcount.Handler) {
	stat := r.Group("/api/stat")
	stat.POST("/read", readCountHandler.SubmitReadRequest)
}
