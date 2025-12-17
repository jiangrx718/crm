package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespRoleUpdateInfo struct {
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
}

func (s *Service) RoleUpdate(ctx context.Context, roleId, roleName, status string, permission []string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMRole.RoleId.Eq(roleId),
	}
	roleEntity, err := g.CRMRole.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if roleEntity == nil {
		return result, fmt.Errorf("role not found")
	}

	roleEntity.RoleName = roleName
	roleEntity.Status = status

	if _, err = g.CRMRole.Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Updates(&roleEntity); err != nil {
		logObj.Errorw("RoleUpdate role error", "role", roleEntity, "error", err)
		return result, err
	}

	result.Data = RespRoleUpdateInfo{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
