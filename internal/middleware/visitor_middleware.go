package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const VisitorIdCookieName = "visitor_id"
const visitorIdPrefix = "v_"
const maxVisitorIdAge = 365 * 24 * time.Hour
const cookieDomain = "zhangdx.cn"

func VisitorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		visitorId, err := context.Cookie(VisitorIdCookieName)
		if err != nil || visitorId == "" {
			visitorId = visitorIdPrefix + uuid.NewString()
			maxAge := int(maxVisitorIdAge.Seconds())
			context.SetSameSite(http.SameSiteLaxMode)
			context.SetCookie(VisitorIdCookieName, visitorId, maxAge, "/", "", false, true)
		}
		context.Set(VisitorIdCookieName, visitorId)
		context.Next()
	}
}
