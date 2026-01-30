package service

import (
	"context"
	"crm/internal/common"
)

type PictureBookIFace interface {
	BookCreate(ctx context.Context, categoryId, title, icon, status string, position, categoryType int) (common.ServiceResult, error)
	BookUpdate(ctx context.Context, bookId, categoryId, title, icon, status string, position, categoryType int) (common.ServiceResult, error)
	BookList(ctx context.Context, offset, limit int64, title string, categoryType int, categoryId string) (common.ServiceResult, error)
	BookDelete(ctx context.Context, bookId string) (common.ServiceResult, error)
	BookStatus(ctx context.Context, bookId, status string) (common.ServiceResult, error)
	BookDetail(ctx context.Context, bookId string) (common.ServiceResult, error)
}
