package pageview

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/api"
)

type Handler struct {
	pgv *Service
}

func NewHandler(pgv *Service) *Handler {
	return &Handler{pgv: pgv}
}

func (h *Handler) SubmitPageViewHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var mpv PageViewMultiRequest
	var prList []PageViewRequest
	err := c.ShouldBindJSON(&prList)
	if err != nil {
		_ = c.Error(err)
		return
	}
	mpv.Request = prList
	fillHttpRequestParam(c, &mpv)
	_, err = h.pgv.SubmitPageView(ctx, &mpv)
	if err != nil {
		_ = c.Error(err)
	} else {
		c.JSON(http.StatusOK, api.SuccessResult(nil))
	}
}

func fillHttpRequestParam(c *gin.Context, request *PageViewMultiRequest) {
	request.IP = c.ClientIP()
	request.UserAgent = c.Request.UserAgent()
	request.Referer = c.Request.Referer()
	if v, exists := c.Get("visitor_id"); exists {
		request.VisitorID = v.(string)
	}
}
