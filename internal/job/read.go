package job

import (
	"context"
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
	ticker := time.NewTicker(5 * time.Minute)
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
	log.Printf("runSyncTask start")
	ids, err := job.redisOperator.Client.HKeys(ctx, readcount.ReadCountKey).Result()
	if err != nil {
		log.Fatalf("runSyncTask get read count map error: %v\n", err)
		return err
	}
	log.Printf("runSyncTask read count size: %v\n", len(ids))
	for _, id := range ids {
		val, err := job.redisOperator.Client.HGetDel(ctx, readcount.ReadCountKey, id).Result()
		if len(val) == 0 || err != nil {
			continue
		}
		readCount, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			continue
		}
		log.Printf("update read count itemID: %s, count: %d\n", id, readCount)
		err = job.repository.IncrReadCount(ctx, id, readCount)
		if err != nil {
			continue
		}
	}
	log.Printf("runSyncTask end")
	return nil
}
