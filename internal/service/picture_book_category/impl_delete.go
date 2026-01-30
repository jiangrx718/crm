package picture_book_category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
)

func (s *Service) CategoryDelete(ctx context.Context, categoryId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	if _, err := g.SPictureBookCategory.Where(g.SPictureBookCategory.CategoryId.Eq(categoryId)).Delete(); err != nil {
		logObj.Errorw("SPictureBookCategory Delete error", "categoryId", categoryId, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
