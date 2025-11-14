package demo

import (
	"context"
	"time"
	"web/gopkg/gorms"
	"web/gopkg/log"
	"web/gopkg/paging"
	"web/gopkg/services"
	"web/internal/model"

	"go.uber.org/zap"
)

func (s *Service) PagingDemo(ctx context.Context, page gorms.Page) (services.Result, error) {
	logPrefix := "/internal/service/demo: Service.PagingDemo()"

	demoPaging, err := s.demoDao.Pagination(ctx, page)
	if err != nil {
		log.Sugar().Error(logPrefix, zap.Any("demo dao pagination error", err), zap.Any("page", page))
		return services.Failed(ctx, err)
	}
	return services.Success(ctx, paging.NewPaging(demoPaging.Total, NewDemoS(demoPaging.List)))
}

type Demo struct {
	DemoId      string             `json:"demo_id"`
	Name        string             `json:"name"`
	FileType    int                `json:"file_type"`
	ProjectType int                `json:"project_type"`
	Content     string             `json:"content"`
	Metadata    model.DemoMetadata `json:"metadata"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func NewDemo(demo *model.Demo) *Demo {
	if demo == nil {
		return nil
	}
	
	return &Demo{
		DemoId:      demo.DemoId,
		Name:        demo.Name,
		FileType:    demo.FileType,
		ProjectType: demo.ProjectType,
		Content:     demo.Content,
		Metadata:    demo.Metadata.Data(),
		CreatedAt:   demo.CreatedAt,
		UpdatedAt:   demo.UpdatedAt,
	}
}

func NewDemoS(demoEntities []*model.Demo) []*Demo {
	if len(demoEntities) == 0 {
		return nil
	}

	demoS := make([]*Demo, 0)
	for _, demoEntity := range demoEntities {
		demoItem := NewDemo(demoEntity)

		demoS = append(demoS, demoItem)
	}

	return demoS
}
