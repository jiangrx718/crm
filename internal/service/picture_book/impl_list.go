package picture_book

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"

	"gorm.io/gen"
)

type RespBookService struct {
	Id           int    `json:"id"`
	BookId       string `json:"book_id"`
	Title        string `json:"title"`
	Icon         string `json:"icon"`
	CategoryId   string `json:"category_id"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type"`
	CreatedAt    string `json:"created_at"`
}

func (s *Service) BookList(ctx context.Context, offset, limit int64, title string, categoryType int, categoryId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.SPictureBook.Debug()
	var conditions []gen.Condition
	if title != "" {
		conditions = append(conditions, g.SPictureBook.Title.Like("%"+title+"%"))
	}
	if categoryType != 0 {
		conditions = append(conditions, g.SPictureBook.CategoryType.Eq(categoryType))
	}
	if categoryId != "" {
		conditions = append(conditions, g.SPictureBook.CategoryId.Eq(categoryId))
	}

	list, count, err := q.Where(conditions...).Order(g.SPictureBook.Position.Desc(), g.SPictureBook.Id.Desc()).FindByPage(int(offset), int(limit))
	if err != nil {
		logObj.Errorw("BookList Find error", "error", err)
		return result, err
	}

	var respList []RespBookService
	for _, v := range list {
		respList = append(respList, RespBookService{
			Id:           v.Id,
			BookId:       v.BookId,
			Title:        v.Title,
			Icon:         v.Icon,
			CategoryId:   v.CategoryId,
			Status:       v.Status,
			Position:     v.Position,
			CategoryType: v.CategoryType,
			CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	result.Data = map[string]any{"list": respList, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}
