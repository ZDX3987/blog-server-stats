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
	ctx := c.Request.Context()
	var readCountRequest ReadCountRequest
	if err := c.ShouldBindJSON(&readCountRequest); err != nil {
		_ = c.Error(err)
		return
	}
	fillHttpRequestParam(c, &readCountRequest)
	result, err := rch.service.SaveReadCountRequest(ctx, &readCountRequest)
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
	if v, exists := c.Get(visitorIdKey); exists {
		request.VisitorID = v.(string)
	}
}
