package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) RoleDelete(ctx context.Context, roleId string) (common.ServiceResult, error) {
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

	// 检查当前角色下是否存在用户
	adminWhere := []gen.Condition{
		g.CRMAdmin.DepartmentId.Eq(roleId),
	}
	adminList, err := g.CRMAdmin.Where(adminWhere...).Find()
	if err != nil {
		return result, err
	}
	if len(adminList) > 0 {
		return result, fmt.Errorf("当前角色下存在用户，请移除用户后删除")
	}

	if _, err := g.CRMRole.Where(where...).Unscoped().Delete(); err != nil {
		logObj.Errorf("CRMRole Delete role Delete has error(%v)", err)
		return result, err
	}

	result.Data = map[string]string{
		"role_id": roleId,
	}

	result.SetMessage("操作成功")

	return result, nil
}
