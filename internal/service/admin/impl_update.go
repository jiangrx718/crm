package admin

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils/str"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

type RespAdminUpdateInfo struct {
	AdminId      string `json:"admin_id"`
	UserName     string `json:"user_name"`
	UserPhone    string `json:"user_phone"`
	DepartmentId int    `json:"department_id"`
	Status       string `json:"status"`
}

func (s *Service) AdminUpdate(ctx context.Context, adminId, password, status string, departmentId int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMAdmin.AdminId.Eq(adminId),
	}
	adminEntity, err := g.CRMAdmin.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if adminEntity == nil {
		return result, fmt.Errorf("admin not found")
	}

	adminEntity.DepartmentId = departmentId
	adminEntity.Status = status
	if password != "" {
		adminEntity.Password = str.MD5String(fmt.Sprintf("%s%s", password, model.SaltValue))
	}

	if _, err = g.CRMAdmin.Where(
		g.CRMAdmin.AdminId.Eq(adminId),
	).Updates(&adminEntity); err != nil {
		logObj.Errorw("AdminUpdate admin error", "admin", adminEntity, "error", err)
		return result, err
	}

	result.Data = RespAdminUpdateInfo{
		AdminId:      adminId,
		UserName:     adminEntity.UserName,
		UserPhone:    adminEntity.UserPhone,
		DepartmentId: departmentId,
		Status:       status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
