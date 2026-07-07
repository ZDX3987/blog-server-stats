package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const VisitorIdCookieName = "visitor_id"

func VisitorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		visitorId, err := context.Cookie(VisitorIdCookieName)
		if err != nil || visitorId == "" {
			// TODO 未完成visitorID定义
			visitorId = "v_"
			maxAge := int((365 * 24 * time.Hour).Seconds())
			context.SetSameSite(http.SameSiteLaxMode)
			context.SetCookie(VisitorIdCookieName, visitorId, maxAge, "/", "", false, true)
		}
		context.Set(VisitorIdCookieName, visitorId)
		context.Next()
	}
}
