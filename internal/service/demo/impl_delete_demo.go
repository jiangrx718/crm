package demo

import (
	"context"
	"errors"
	"web/gopkg/log"
	"web/gopkg/services"

	"go.uber.org/zap"
)

func (s *Service) DeleteDemo(ctx context.Context, demoId string) (services.Result, error) {

	logPrefix := "/internal/service/demo: Service.DeleteDemo()"

	demoData, err := s.demoDao.Get(ctx, demoId)
	if err != nil {
		log.Sugar().Error(logPrefix, zap.Any("demo dao Get error", err))
		return services.Failed(ctx, err)
	}
	if demoData == nil {
		return services.Failed(ctx, errors.New("数据不存在"))
	}
	if err = s.demoDao.Delete(ctx, demoData.DemoId); err != nil {
		log.Sugar().Error(logPrefix, zap.Any("delete demo error", err), zap.Any("demo_id", demoId))
		return services.Failed(ctx, err)
	}

	return services.Success(ctx, nil)
}
