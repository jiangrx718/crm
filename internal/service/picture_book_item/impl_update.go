package picture_book_item

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

func (s *Service) ItemUpdate(ctx context.Context, id int, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	entity, err := g.SPictureBookItem.Where(g.SPictureBookItem.Id.Eq(id)).Take()
	if err != nil {
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	entity.Title = title
	entity.Pic = pic
	entity.BPic = bPic
	entity.Audio = audio
	entity.Content = content
	entity.Status = status
	entity.Position = position

	if _, err := g.SPictureBookItem.Where(g.SPictureBookItem.Id.Eq(id)).Updates(entity); err != nil {
		logObj.Errorw("SPictureBookItem Update error", "entity", entity, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
