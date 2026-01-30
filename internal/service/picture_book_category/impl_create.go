package picture_book_category

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespCategoryCreateInfo struct {
	CategoryId string `json:"category_id"`
}

func (s *Service) CategoryCreate(ctx context.Context, categoryName, status string, position, categoryType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	categoryId := utils.GenUUID()
	entity := model.SPictureBookCategory{
		CategoryId:   categoryId,
		CategoryName: categoryName,
		Status:       status,
		Position:     position,
		CategoryType: categoryType,
	}

	if err := g.SPictureBookCategory.Create(&entity); err != nil {
		logObj.Errorw("SPictureBookCategory Create error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespCategoryCreateInfo{
		CategoryId: categoryId,
	}
	result.SetMessage("操作成功")
	return result, nil
}
