package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespPermissionUpdateInfo struct {
	PermissionId   string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
	PermissionURL  string `json:"permission_url"`
	ParentId       string `json:"parent_id"`
	Status         string `json:"status"`
}

func (s *Service) PermissionUpdate(ctx context.Context, permissionId, permissionName, permissionUrl, parentId, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
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

	permissionEntity.PermissionName = permissionName
	permissionEntity.PermissionURL = permissionUrl
	permissionEntity.ParentId = parentId
	permissionEntity.Status = status

	if _, err = g.CRMPermission.Where(
		g.CRMPermission.PermissionId.Eq(permissionId),
	).Updates(&permissionEntity); err != nil {
		logObj.Errorw("PermissionUpdate permission error", "permission", permissionEntity, "error", err)
		return result, err
	}

	result.Data = RespPermissionUpdateInfo{
		PermissionId:   permissionId,
		PermissionName: permissionName,
		PermissionURL:  permissionUrl,
		ParentId:       parentId,
		Status:         status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
