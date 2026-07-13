package readcount

import (
	"fmt"
	"time"

	"zhangdx.cn/blog-server-stats/internal/apperror"
	"zhangdx.cn/blog-server-stats/internal/infra"
)

type Service struct {
	redisOperator *infra.RedisOperator
	repository    *Repository
}

func NewReadCountService(redisOperator *infra.RedisOperator, repo *Repository) *Service {
	return &Service{redisOperator, repo}
}

func (rcs *Service) SaveReadCountRequest(request *ReadCountRequest) (bool, *apperror.Error) {
	if request.ItemID == "" {
		return false, apperror.BusinessError("目标ID参数不能为空")
	}
	if !rcs.repository.IsValid(request.ItemID) {
		return false, apperror.BusinessError("目标ID不是有效的")
	}
	if !isValidRead(request.ReadDuration, request.ScrollDepth) {
		return false, nil
	}
	dedupKey := fmt.Sprintf("article:read:dedup:%s:%s", request.ItemID, request.Identity)
	ok := rcs.redisOperator.SetNx(dedupKey, "1", 30*time.Minute)
	if !ok {
		return false, nil
	}
	return true, nil
}

func isValidRead(duration int, depth int) bool {
	return duration >= 5 || depth >= 50
}
