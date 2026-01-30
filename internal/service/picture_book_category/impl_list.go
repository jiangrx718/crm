package picture_book_category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"

	"gorm.io/gen"
)

type RespCategoryService struct {
	Id           int    `json:"id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryType int    `json:"category_type"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
	CreatedAt    string `json:"created_at"`
}

func (s *Service) CategoryList(ctx context.Context, offset, limit int64, categoryName string, categoryType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.SPictureBookCategory.Debug()
	var conditions []gen.Condition
	if categoryName != "" {
		conditions = append(conditions, g.SPictureBookCategory.CategoryName.Like("%"+categoryName+"%"))
	}
	if categoryType != 0 {
		conditions = append(conditions, g.SPictureBookCategory.CategoryType.Eq(categoryType))
	}

	list, count, err := q.Where(conditions...).Order(g.SPictureBookCategory.Position.Desc(), g.SPictureBookCategory.Id.Desc()).FindByPage(int(offset), int(limit))
	if err != nil {
		logObj.Errorw("CategoryList Find error", "error", err)
		return result, err
	}

	var respList []RespCategoryService
	for _, v := range list {
		respList = append(respList, RespCategoryService{
			Id:           v.Id,
			CategoryId:   v.CategoryId,
			CategoryName: v.CategoryName,
			CategoryType: v.CategoryType,
			Status:       v.Status,
			Position:     v.Position,
			CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	result.Data = map[string]any{"list": respList, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}
