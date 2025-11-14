package demo

import (
	"context"
	"strings"
	"web/internal/g"
	"web/internal/model"
)

func (d *Dao) Get(ctx context.Context, id string) (*model.Demo, error) {
	if strings.TrimSpace(id) == "" {
		return nil, nil
	}

	demoInfo, err := g.Demo.Where(g.Demo.DemoId.Eq(id)).First()
	if err != nil {
		return nil, d.ConvertError(err)
	}

	return demoInfo, nil
}
