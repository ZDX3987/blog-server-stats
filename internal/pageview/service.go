package pageview

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (pvs *Service) SubmitPageView(ctx context.Context, requestParam *PageViewMultiRequest) (bool, error) {
	requestList := requestParam.Request
	if len(requestList) == 0 {
		return false, nil
	}
	mpv := make([]*PageView, len(requestList))
	for i, request := range requestList {
		pv := PageView{
			VisitorID:      requestParam.VisitorID,
			Path:           request.Path,
			Referer:        requestParam.Referer,
			UserAgent:      requestParam.UserAgent,
			IP:             requestParam.IP,
			DurationSec:    request.DurationSec,
			MaxScrollDepth: request.MaxScrollDepth,
			OccurredAt:     request.OccurredAt,
		}
		mpv[i] = &pv
	}
	result, err := pvs.repository.SavePageView(ctx, mpv)
	return result, err
}
