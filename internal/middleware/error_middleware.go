package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/api"
	"zhangdx.cn/blog-server-stats/internal/apperror"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err

		var appErr *apperror.Error
		if errors.As(err, &appErr) {
			c.JSON(http.StatusOK, &api.ResponseResult{Msg: appErr.Msg, Code: appErr.Code})
			return
		}
		c.JSON(http.StatusInternalServerError, api.Fail("服务异常"))
	}
}
