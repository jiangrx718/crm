package demo

import (
	"context"
	"web/gopkg/services"
	"web/internal/model"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func (s *Service) CreateDemo(ctx context.Context, name string, fileType int, projectType int, metadata model.DemoMetadata) (services.Result, error) {
	logPrefix := "/internal/service/demo: Service.CreateDemo()"
	demoEntity, err := s.demoDao.Create(ctx, name, fileType, projectType, "", metadata)
	if err != nil {
		hlog.Errorf("%s  发生 demo dao Create() error: %v\n", logPrefix, err)
		return services.Failed(ctx, err)
	}

	return services.Success(ctx, &demoEntity)
}
