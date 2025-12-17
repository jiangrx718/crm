package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) PermissionDelete(ctx context.Context, permissionId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查通知是否存在
	where := []gen.Condition{
		g.CRMPermission.PermissionId.Eq(permissionId),
	}
	permissionEntity, err := g.CRMPermission.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if permissionEntity == nil {
		return result, fmt.Errorf("permission not found")
	}

	if _, err := g.CRMPermission.Where(where...).Unscoped().Delete(); err != nil {
		logObj.Errorf("CRMPermission Delete permission Delete has error(%v)", err)
		return result, err
	}

	result.Data = map[string]string{
		"permission_id": permissionId,
	}

	result.SetMessage("操作成功")

	return result, nil
}
