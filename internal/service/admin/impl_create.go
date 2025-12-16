package admin

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/gopkg/utils/str"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"
)

type RespAdminCreateInfo struct {
	AdminId      string `json:"admin_id"`
	UserName     string `json:"user_name"`
	UserPhone    string `json:"user_phone"`
	DepartmentId int    `json:"department_id"`
	Status       int    `json:"status"`
}

func (s *Service) AdminCreate(ctx context.Context, userName, userPhone, password string, departmentId, status int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	crmAdmin := model.CRMAdmin{
		AdminId:      utils.GenUUID(),
		UserName:     userName,
		UserPhone:    userPhone,
		Password:     str.MD5String(fmt.Sprintf("%s%s", password, model.SaltValue)),
		DepartmentId: departmentId,
		Status:       status,
	}

	if createErr := g.CRMAdmin.Create(&crmAdmin); createErr != nil {
		logObj.Errorw("CRMAdmin Create crmAdmin error", "crmAdmin", crmAdmin, "error", createErr)
		return result, createErr
	}

	result.Data = RespAdminCreateInfo{
		AdminId:      utils.GenUUID(),
		UserName:     userName,
		UserPhone:    userPhone,
		DepartmentId: departmentId,
		Status:       status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
