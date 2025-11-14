package demo

import (
	"context"
	"strings"
	"web/internal/g"
)

func (d *Dao) Delete(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return nil
	}

	_, err := g.Demo.WithContext(ctx).Where(g.Demo.DemoId.Eq(id)).Delete()
	if err != nil {
		return d.ConvertError(err)
	}

	return nil
}
