package readcount

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/api"
)

type Handler struct {
	readCountService *Service
}

func NewReadCountHandler(readCountService *Service) *Handler {
	return &Handler{readCountService}
}

func (rch *Handler) SubmitReadRequest(c *gin.Context) {
	var readCountRequest ReadCountRequest
	c.ShouldBindJSON(&readCountRequest)
	readCountRequest.IP = c.ClientIP()
	err := rch.readCountService.SaveReadCountRequest(&readCountRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, api.Success())
	}
}
