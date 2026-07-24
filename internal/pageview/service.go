package pageview

import "context"

type Service struct {
	repository *Repository
}

func (pvs *Service) SubmitPageView(ctx context.Context, requestParam *PageViewRequest) (bool, error) {
	pv := PageView{
		VisitorID:      requestParam.VisitorID,
		Path:           requestParam.Path,
		Referer:        requestParam.Referer,
		UserAgent:      requestParam.UserAgent,
		IP:             requestParam.IP,
		DurationSec:    requestParam.DurationSec,
		MaxScrollDepth: requestParam.MaxScrollDepth,
	}
	result, err := pvs.repository.SavePageView(ctx, &pv)
	return result, err
}
