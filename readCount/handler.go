package readCount

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/rdb"
)

type ReadCountForm struct {
	ItemId string `json:"itemId"`
	//read_time time.Time
}

func Handler(c *gin.Context) {
	var param ReadCountForm
	c.ShouldBind(&param)
	log.Printf("params: %s\n", param.ItemId)
	rdb.Set("itemId", param.ItemId)
	val := rdb.Get("itemId")
	c.JSON(http.StatusOK, gin.H{"msg": "OK", "redisVal": val})
}
