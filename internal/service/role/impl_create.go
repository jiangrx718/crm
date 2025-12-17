package role

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespRoleCreateInfo struct {
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
}

func (s *Service) RoleCreate(ctx context.Context, roleName, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	roleId := utils.GenUUID()
	crmRole := model.CRMRole{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	if createErr := g.CRMRole.Create(&crmRole); createErr != nil {
		logObj.Errorw("CRMRole Create crmRole error", "crmRole", crmRole, "error", createErr)
		return result, createErr
	}

	result.Data = RespRoleCreateInfo{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	result.SetMessage("操作成功")
	return nil, nil
}
