package picture_book_item

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

func (s *Service) ItemStatus(ctx context.Context, id int, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	info, err := g.SPictureBookItem.Where(g.SPictureBookItem.Id.Eq(id)).Update(g.SPictureBookItem.Status, status)
	if err != nil {
		logObj.Errorw("ItemStatus Update error", "id", id, "error", err)
		return result, err
	}
	if info.RowsAffected == 0 {
		return result, fmt.Errorf("record not found")
	}

	result.SetMessage("操作成功")
	return result, nil
}
