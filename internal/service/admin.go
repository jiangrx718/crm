package service

import (
	"context"
	"crm/internal/common"
)

type AdminIFace interface {
	AdminCreate(ctx context.Context, userName, userPhone, password string, departmentId int) (common.ServiceResult, error)
}
