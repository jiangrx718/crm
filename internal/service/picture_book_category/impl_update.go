package picture_book_category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

func (s *Service) CategoryUpdate(ctx context.Context, categoryId, categoryName, status string, position, categoryType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	entity, err := g.SPictureBookCategory.Where(g.SPictureBookCategory.CategoryId.Eq(categoryId)).Take()
	if err != nil {
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	entity.CategoryName = categoryName
	entity.Status = status
	entity.Position = position
	entity.CategoryType = categoryType

	if _, err := g.SPictureBookCategory.Where(g.SPictureBookCategory.CategoryId.Eq(categoryId)).Updates(entity); err != nil {
		logObj.Errorw("SPictureBookCategory Update error", "entity", entity, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
