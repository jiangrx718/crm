package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespPermissionCreateInfo struct {
	PermissionId   string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
	PermissionURL  string `json:"permission_url"`
	ParentId       string `json:"parent_id"`
	Status         string `json:"status"`
}

func (s *Service) PermissionCreate(ctx context.Context, permissionName, permissionUrl, parentId, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	permissionId := utils.GenUUID()
	crmPermission := model.CRMPermission{
		PermissionId:   permissionId,
		PermissionName: permissionName,
		PermissionURL:  permissionUrl,
		ParentId:       parentId,
		Status:         status,
	}
	if createErr := g.CRMPermission.Create(&crmPermission); createErr != nil {
		logObj.Errorw("CRMPermission Create crmPermission error", "crmPermission", crmPermission, "error", createErr)
		return result, createErr
	}

	result.Data = RespPermissionCreateInfo{
		PermissionId:   permissionId,
		PermissionName: permissionName,
		PermissionURL:  permissionUrl,
		ParentId:       parentId,
		Status:         status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
