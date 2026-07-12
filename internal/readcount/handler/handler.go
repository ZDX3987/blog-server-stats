package handler

import (
	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/readcount/model"
	"zhangdx.cn/blog-server-stats/internal/readcount/service"
)

type ReadCountHandler struct {
	readCountService *service.ReadCountService
}

func NewReadCountHandler(readCountService *service.ReadCountService) *ReadCountHandler {
	return &ReadCountHandler{readCountService}
}

func (h *ReadCountHandler) SubmitReadRequest(c *gin.Context) {
	var readCountForm model.ReadCountForm
	c.ShouldBindJSON(&readCountForm)
	h.readCountService.SaveReadCountRequest(&readCountForm)
}
