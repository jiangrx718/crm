package demo

import (
	"context"
	"errors"
	"web/gopkg/log"
	"web/gopkg/services"

	"go.uber.org/zap"
)

func (s *Service) UpdateDemo(ctx context.Context, demoId string, content string) (services.Result, error) {
	logPrefix := "/internal/service/demo: Service.UpdateDemo()"

	// 获取对应的详情
	demoEntity, err := s.demoDao.Get(ctx, demoId)
	if err != nil {
		log.Sugar().Error(logPrefix, zap.Any("demo dao get data error", err), zap.Any("demo_id", demoId))
		return nil, err
	}
	if demoEntity == nil {
		return nil, errors.New("数据不存在")
	}

	if _, err := s.demoDao.Update(ctx, nil, demoId, content); err != nil {
		log.Sugar().Error(logPrefix, zap.Any("demo dao update error", err),
			zap.Any("demo_id", demoId),
			zap.Any("content", content))
		return services.Failed(ctx, err)
	}

	demoEntity.Content = content
	return services.Success(ctx, demoEntity)
}
