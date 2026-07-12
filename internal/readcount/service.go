package readcount

import (
	"fmt"

	"zhangdx.cn/blog-server-stats/internal/infra"
)

type Service struct {
	redisOperator *infra.RedisOperator
}

func NewReadCountService(redisOperator *infra.RedisOperator) *Service {
	return &Service{redisOperator}
}

func (rcs *Service) SaveReadCountRequest(form *ReadCountRequest) error {
	fmt.Printf("打印参数：%+v\n", form)
	if form.ItemID == "" {
		return fmt.Errorf("目标ID参数不能为空")
	}
	return nil
}
