package admin

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

func (s *Service) AdminCreate(ctx context.Context, userName, userPhone, password string, departmentId int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)
	crmAdmin := model.CRMAdmin{
		UserName:     userName,
		UserPhone:    userPhone,
		Password:     password,
		DepartmentId: departmentId,
	}

	if createErr := g.CRMAdmin.Create(&crmAdmin); createErr != nil {
		logObj.Errorw("CRMAdmin Create crmAdmin error", "crmAdmin", crmAdmin, "error", createErr)
		return result, createErr
	}

	result.Data = model.CRMAdmin{
		UserName:     userName,
		UserPhone:    userPhone,
		Password:     password,
		DepartmentId: departmentId,
		Status:       model.StatusOn,
	}
	result.SetMessage("操作成功")
	return result, nil
}
