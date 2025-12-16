package crm_admin

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/g"
	"crm/internal/model"

	"go.uber.org/zap"
)

func (d *Dao) Create(ctx context.Context, userName, userPhone, password string, departmentId int) (*model.CRMAdmin, error) {
	logPrefix := "/internal/dao/crm_admin: Dao.Create()"

	crmAdminItem := model.CRMAdmin{
		UserName:     userName,
		UserPhone:    userPhone,
		Password:     password,
		DepartmentId: departmentId,
		Status:       model.StatusOn,
	}

	if err := g.CRMAdmin.Create(&crmAdminItem); err != nil {
		log.Sugar().Error(ctx, logPrefix, zap.Any("crm_admin record", crmAdminItem))
		return nil, err
	}
	return &crmAdminItem, nil
}
