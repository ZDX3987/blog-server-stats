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
	_ = c.ShouldBindJSON(&readCountRequest)
	fillHttpRequestParam(c, &readCountRequest)
	result, err := rch.service.SaveReadCountRequest(&readCountRequest)
	if err != nil {
		_ = c.Error(err)
	} else {
		c.JSON(http.StatusOK, api.SuccessResult(result))
	}
}

func fillHttpRequestParam(c *gin.Context, request *ReadCountRequest) {
	request.IP = c.ClientIP()
	request.UserAgent = c.Request.UserAgent()
	request.Referer = c.Request.Referer()
	request.Identity = BuildIdentity(c)
}
