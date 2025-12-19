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
	DepartmentId string `json:"department_id"`
	Status       string `json:"status"`
}

func (s *Service) AdminCreate(ctx context.Context, userName, userPhone, password, status, departmentId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	adminEntity, err := g.CRMAdmin.Where(
		g.CRMAdmin.Where(g.CRMAdmin.UserName.Eq(userName)).
			Or(g.CRMAdmin.UserPhone.Eq(userPhone)),
	).Take()

	if err != nil && err.Error() != "record not found" {
		logObj.Errorw("AdminCreate Check Exist Error", "error", err)
		return result, err
	}
	if adminEntity != nil {
		result.SetCode(10001) // 业务错误码
		result.SetMessage("用户名或手机号已存在")
		return result, nil // 返回 nil error，让 controller 处理 result
	}

	adminId := utils.GenUUID()
	crmAdmin := model.CRMAdmin{
		AdminId:      adminId,
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
		AdminId:      adminId,
		UserName:     userName,
		UserPhone:    userPhone,
		DepartmentId: departmentId,
		Status:       status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
