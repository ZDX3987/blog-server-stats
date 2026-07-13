package readcount

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const visitorIdKey = "visitor_id"
const userIdKey = "user_id"
const anonymous = "anonymous"
const unknown = "unknown"

func BuildIdentity(c *gin.Context) string {
	if userID, exists := c.Get(userIdKey); exists {
		return fmt.Sprintf("user%v", userID)
	}
	visitorID, exists := c.Get(visitorIdKey)
	if !exists {
		return fmt.Sprintf("%s:%s", anonymous, unknown)
	}
	return fmt.Sprintf("%s:%v", anonymous, visitorID)
}
