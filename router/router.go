package router

import (
	"blog-server-stats/readCount"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	stat := r.Group("/api/stat")
	stat.POST("/read", readCount.Handler)
}
