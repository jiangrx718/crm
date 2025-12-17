package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespRoleService struct {
	RoleId    string `json:"role_id"`
	RoleName  string `json:"role_name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (s *Service) RoleList(ctx context.Context, offset, limit int64) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	roleDataList, count, err := ScanByPage(offset, limit)
	if err != nil {
		logObj.Errorw("AdminList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listRole []RespRoleService
	for idx, _ := range roleDataList {
		listRole = append(listRole, RespRoleService{
			RoleId:    roleDataList[idx].RoleId,
			RoleName:  roleDataList[idx].RoleName,
			Status:    roleDataList[idx].Status,
			CreatedAt: roleDataList[idx].CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	if len(listRole) == 0 {
		listRole = []RespRoleService{}
	}
	result.Data = map[string]any{"list": listRole, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(offset, limit int64) ([]*model.CRMRole, int64, error) {
	var (
		crmRole  = g.CRMRole
		response = make([]*model.CRMRole, 0)
	)

	q := crmRole.Debug()
	where := []gen.Condition{}

	count, err := q.Where(where...).Order(crmRole.Id.Desc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
