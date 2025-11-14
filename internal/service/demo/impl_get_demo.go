package demo

import (
	"context"
	"web/gopkg/log"
	"web/gopkg/services"

	"go.uber.org/zap"
)

func (s *Service) GetDemo(ctx context.Context, demoId string) (services.Result, error) {
	logPrefix := "/internal/service/demo: Service.GetDemo()"

	demoData, err := s.demoDao.Get(ctx, demoId)
	if err != nil {
		log.Sugar().Error(logPrefix, zap.Any("demo dao Get error", err))
		return services.Failed(ctx, err)
	}
	
	return services.Success(ctx, demoData)
}
