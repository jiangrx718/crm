package demo

import (
	"context"
	"web/gopkg/gorms"
	"web/internal/model"

	"web/internal/g"
)

func (d *Dao) Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.Demo], error) {
	paging, err := gorms.PaginationQuery(
		g.Demo.Order(
			g.Demo.CreatedAt.Desc(),
		).FindByPage, gorms.Page{
			PageIndex: page.PageIndex,
			PageSize:  page.PageSize,
		})
	if err != nil {
		return nil, d.ConvertError(err)
	}

	return paging, nil
}
