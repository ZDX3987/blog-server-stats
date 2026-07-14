package job

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"zhangdx.cn/blog-server-stats/internal/infra"
	"zhangdx.cn/blog-server-stats/internal/readcount"
)

type ReadCountSyncJob struct {
	redisOperator *infra.RedisOperator
	repository    *readcount.Repository
}

func NewReadCountSyncJob(r *infra.RedisOperator, repo *readcount.Repository) *ReadCountSyncJob {
	return &ReadCountSyncJob{redisOperator: r, repository: repo}
}

func (job *ReadCountSyncJob) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				_ = job.runSyncTask(ctx)
			}
		}
	}()
}

func (job *ReadCountSyncJob) runSyncTask(ctx context.Context) error {
	itemIds, err := job.redisOperator.ListSet(ctx, "article:read:dirty")
	if err != nil {
		return err
	}
	for _, id := range itemIds {
		countKey := fmt.Sprintf("article:read:count:%s", id)
		val, err := job.redisOperator.Get(ctx, countKey)
		if err != nil {
			continue
		}
		readCount, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}
		log.Printf("update read count itemID: %s, count: %d\n", id, readCount)
		err = job.repository.UpdateReadCount(ctx, id, readCount)
		if err != nil {
			continue
		}
		err = job.redisOperator.Del(ctx, countKey)
		if err != nil {
			continue
		}
	}
	return nil
}
