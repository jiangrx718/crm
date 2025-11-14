package demo

import (
	"context"
	"strings"
	"web/internal/g"
	"web/internal/model"

	"github.com/pkg/errors"

	"gorm.io/gen"
)

func (d *Dao) Update(ctx context.Context, tx *g.Query, id string, content string) (info *gen.ResultInfo, err error) {
	if tx == nil {
		tx = g.Q
	}

	if strings.TrimSpace(id) == "" {
		err = errors.New("writing_online dao update id is empty")
		return
	}

	result, err := tx.Demo.Where(tx.Demo.DemoId.Eq(id)).Updates(&model.Demo{
		Content: content,
	})
	if err != nil {
		return nil, d.ConvertError(err)
	}

	return &result, nil
}
