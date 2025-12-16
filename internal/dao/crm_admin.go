package dao

import (
	"context"
	"crm/internal/model"
)

type CRMAdmin interface {
	Create(ctx context.Context, userName, userPhone, password string, departmentId int) (*model.CRMAdmin, error)
}
