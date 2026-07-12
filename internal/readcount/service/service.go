package service

import (
	"fmt"

	"zhangdx.cn/blog-server-stats/internal/infra/redis"
	"zhangdx.cn/blog-server-stats/internal/readcount/model"
)

type ReadCountService struct {
	redisOperator *redis.RedisOperator
}

func NewReadCountService(redisOperator *redis.RedisOperator) *ReadCountService {
	return &ReadCountService{redisOperator}
}

func (s *ReadCountService) SaveReadCountRequest(form *model.ReadCountForm) {
	fmt.Printf("打印参数：%s\n", form.ItemId)
}
