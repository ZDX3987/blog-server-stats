package readcount

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/api"
)

type Handler struct {
	service *Service
}

func NewReadCountHandler(readCountService *Service) *Handler {
	return &Handler{readCountService}
}

func (rch *Handler) SubmitReadRequest(c *gin.Context) {
	var readCountRequest ReadCountRequest
	c.ShouldBindJSON(&readCountRequest)
	fillHttpRequestParam(c, &readCountRequest)
	err := rch.service.SaveReadCountRequest(&readCountRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, api.Success())
	}
}

func fillHttpRequestParam(c *gin.Context, request *ReadCountRequest) {
	request.IP = c.ClientIP()
	request.UserID = c.Request.UserAgent()
	request.Referer = c.Request.Referer()
}
