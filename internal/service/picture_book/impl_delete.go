package picture_book

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
)

func (s *Service) BookDelete(ctx context.Context, bookId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	if _, err := g.SPictureBook.Where(g.SPictureBook.BookId.Eq(bookId)).Delete(); err != nil {
		logObj.Errorw("SPictureBook Delete error", "bookId", bookId, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
