package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespPermissionService struct {
	PermissionId   string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
	PermissionUrl  string `json:"permission_url"`
	ParentId       string `json:"parent_id"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
}

func (s *Service) PermissionList(ctx context.Context) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	permissionDataList, count, err := ScanByPage()
	if err != nil {
		logObj.Errorw("PermissionList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listPermission []RespPermissionService
	for idx, _ := range permissionDataList {
		listPermission = append(listPermission, RespPermissionService{
			PermissionId:   permissionDataList[idx].PermissionId,
			PermissionName: permissionDataList[idx].PermissionName,
			PermissionUrl:  permissionDataList[idx].PermissionURL,
			ParentId:       permissionDataList[idx].ParentId,
			Status:         permissionDataList[idx].Status,
			CreatedAt:      permissionDataList[idx].CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	if len(listPermission) == 0 {
		listPermission = []RespPermissionService{}
	}
	result.Data = map[string]any{"list": listPermission, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage() ([]*model.CRMPermission, int64, error) {
	var (
		crmPermission = g.CRMPermission
		response      = make([]*model.CRMPermission, 0)
	)

	q := crmPermission.Debug()
	where := []gen.Condition{}

	count, err := q.Where(where...).Order(crmPermission.Id.Desc()).ScanByPage(&response, int(1), int(100))
	return response, count, err
}
