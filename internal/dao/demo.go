package dao

import (
	"context"
	"web/gopkg/gorms"
	"web/internal/g"
	"web/internal/model"

	"gorm.io/gen"
)

type Demo interface {
	Create(ctx context.Context, name string, fileType int, projectType int, content string, metadata model.DemoMetadata) (*model.Demo, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.Demo, error)
	Update(ctx context.Context, tx *g.Query, id string, content string) (*gen.ResultInfo, error)
	Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.Demo], error)
}
