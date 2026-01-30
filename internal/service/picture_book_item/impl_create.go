package picture_book_item

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespItemCreateInfo struct {
	Id int `json:"id"`
}

func (s *Service) ItemCreate(ctx context.Context, bookId, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	entity := model.SPictureBookItem{
		BookId:   bookId,
		Title:    title,
		Pic:      pic,
		BPic:     bPic,
		Audio:    audio,
		Content:  content,
		Status:   status,
		Position: position,
	}

	if err := g.SPictureBookItem.Create(&entity); err != nil {
		logObj.Errorw("SPictureBookItem Create error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespItemCreateInfo{
		Id: entity.Id,
	}
	result.SetMessage("操作成功")
	return result, nil
}
