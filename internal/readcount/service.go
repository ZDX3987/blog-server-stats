package readcount

import (
	"fmt"

	"zhangdx.cn/blog-server-stats/internal/infra"
)

type Service struct {
	redisOperator *infra.RedisOperator
	repository    *Repository
}

func NewReadCountService(redisOperator *infra.RedisOperator, repo *Repository) *Service {
	return &Service{redisOperator, repo}
}

func (rcs *Service) SaveReadCountRequest(request *ReadCountRequest) error {
	if request.ItemID == "" {
		return fmt.Errorf("目标ID参数不能为空")
	}
	if !rcs.repository.IsValid(request.ItemID) {
		return fmt.Errorf("目标ID不是有效的")
	}
	if !isValidRead(request.ReadDuration, request.ScrollDepth) {
		return nil
	}
	return nil
}

func isValidRead(duration int, depth int) bool {
	return duration >= 5 || depth >= 50
}
