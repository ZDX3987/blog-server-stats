package readcount

import (
	"context"
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

func (rcs *Service) SaveReadCountRequest(ctx context.Context, request *ReadCountRequest) (bool, error) {
	if request.ItemID == "" {
		return false, apperror.BusinessError("目标ID参数不能为空")
	}
	if !rcs.repository.IsValid(ctx, request.ItemID) {
		return false, apperror.BusinessError("目标ID不是有效的")
	}
	if !isValidRead(request.ReadDuration, request.ScrollDepth) {
		return false, nil
	}
	dedupKey := fmt.Sprintf("article:read:dedup:%s:%s", request.ItemID, request.Identity)
	ok, err := rcs.redisOperator.SetNx(ctx, dedupKey, "1", 30*time.Minute)
	if !ok || err != nil {
		return false, err
	}
	ok, err = rcs.redisOperator.AddSet(ctx, "article:read:dirty", request.ItemID, 30*time.Minute)
	if !ok || err != nil {
		return false, err
	}
	log := &ReadCountLog{
		ItemId:       request.ItemID,
		Identity:     request.Identity,
		VisitorID:    request.VisitorID,
		VisitorIp:    request.IP,
		UserAgent:    request.UserAgent,
		Referer:      request.Referer,
		ReadDuration: request.ReadDuration,
		ReadDepth:    request.ScrollDepth,
	}
	result, err := rcs.repository.InsertReadLog(ctx, log)
	return result, err
}

func isValidRead(duration int, depth int) bool {
	return duration >= 5 || depth >= 50
}
