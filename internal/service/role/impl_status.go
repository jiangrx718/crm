package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespRoleStatusInfo struct {
	RoleId string `json:"role_id"`
	Status string `json:"status"`
}

func (s *Service) RoleStatus(ctx context.Context, roleId, status string) (common.ServiceResult, error) {
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

	roleEntity.Status = status

	if _, err = g.CRMRole.Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Updates(&roleEntity); err != nil {
		logObj.Errorw("RoleStatus role error", "role", roleEntity, "error", err)
		return result, err
	}

	result.Data = RespRoleStatusInfo{
		RoleId: roleId,
		Status: status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
