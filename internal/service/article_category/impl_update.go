package article_category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) CategoryUpdate(ctx context.Context, categoryId, parentId, categoryName, categoryImage, status string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMArticleCategory.CategoryId.Eq(categoryId),
	}
	categoryEntity, err := g.CRMArticleCategory.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if categoryEntity == nil {
		return result, fmt.Errorf("category not found")
	}

	categoryEntity.CategoryName = categoryName
	categoryEntity.CategoryImage = categoryImage
	categoryEntity.Status = status
	categoryEntity.Position = position

	if _, err = g.CRMArticleCategory.Where(
		g.CRMArticleCategory.CategoryId.Eq(categoryId),
	).Updates(&categoryEntity); err != nil {
		logObj.Errorw("CategoryUpdate error", "categoryEntity", categoryEntity, "error", err)
		return result, err
	}

	result.Data = RespArticleCategoryInfo{
		CategoryId:    categoryId,
		CategoryName:  categoryName,
		CategoryImage: categoryImage,
		ParentId:      parentId,
		Position:      position,
		Status:        status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
