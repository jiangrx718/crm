package login

import (
	"context"
	"crm/internal/common"
)

type RespAdminCreateInfo struct {
	AdminId      string `json:"admin_id"`
	UserName     string `json:"user_name"`
	UserPhone    string `json:"user_phone"`
	DepartmentId int    `json:"department_id"`
	Status       string `json:"status"`
}

func (s *Service) Login(ctx context.Context, userName, password string) (common.ServiceResult, error) {
	var (
		result = common.NewCRMServiceResult()
	)

	result.SetMessage("操作成功")
	return result, nil
}
