package pageview

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/api"
)

type Handler struct {
	pgv *Service
}

func (h *Handler) SubmitPageViewHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var pv PageViewRequest
	err := c.ShouldBindJSON(pv)
	if err != nil {
		_ = c.Error(err)
		return
	}
	fillHttpRequestParam(c, &pv)
	_, err = h.pgv.SubmitPageView(ctx, &pv)
	if err != nil {
		_ = c.Error(err)
	} else {
		c.JSON(http.StatusOK, api.SuccessResult(nil))
	}
}

func fillHttpRequestParam(c *gin.Context, request *PageViewRequest) {
	request.IP = c.ClientIP()
	request.UserAgent = c.Request.UserAgent()
	request.Referer = c.Request.Referer()
	if v, exists := c.Get("visitor_id"); exists {
		request.VisitorID = v.(string)
	}
}
