package permission

import (
	"context"
	"crm/gopkg/cache/redis"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"
)

// PermissionMenu 获取菜单
func (s *Service) PermissionMenu(ctx context.Context, adminId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	fmt.Printf("permission menu adminId:%s\n", adminId)
	// 1. 获取用户权限信息 (复用 PermissionFind 逻辑或直接查缓存)
	// var permInfo RespPermissionFindInfo
	cacheKey := fmt.Sprintf("login_auth:%s", adminId)
	rdb, err := redis.ClientAndErr("crm")
	if err == nil && rdb != nil {
		if val, err := rdb.Get(ctx, cacheKey).Result(); err == nil && val != "" {
			// 尝试解析缓存中的结构体 (注意：PermissionFind 之前存的是 string "1"，现在为了菜单需要存结构体)
			// 如果之前只存了 "1"，这里会解析失败或者拿到默认值，需要重新查库
			// 为了稳妥，建议直接查库，或者优化 PermissionFind 存完整信息
			// 鉴于 PermissionFind 刚改为存 "1"，这里我们先直接查库，确保数据准确
		}
	}

	// 重新查询用户权限信息 (Copy from PermissionFind logic mostly)
	admin, err := g.CRMAdmin.WithContext(ctx).Where(
		g.CRMAdmin.AdminId.Eq(adminId),
		g.CRMAdmin.Status.Eq(model.StatusOn),
	).Take()
	if err != nil || admin == nil {
		return result, fmt.Errorf("用户不存在或已被禁用")
	}

	roleId := admin.RoleId
	roleInfo, err := g.CRMRole.WithContext(ctx).Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Take()
	if err != nil || roleInfo == nil {
		return result, fmt.Errorf("用户角色不存在")
	}

	// 2. 根据 IsSuper 决定查询逻辑
	var permissionDataList []*model.CRMPermission

	if roleInfo.IsSuper == model.IsSuperOn {
		// 超管：查询所有开启状态的权限
		permissionDataList, _, err = ScanByPage(model.StatusOn)
		if err != nil {
			logObj.Errorw("PermissionMenu Super ScanByPage error", "error", err)
			return result, err
		}
	} else {
		// 普通用户：查询角色关联的权限 ID
		rolePermissions, err := g.CRMRolePermission.WithContext(ctx).Where(
			g.CRMRolePermission.RoleId.Eq(roleId),
		).Find()
		if err != nil {
			return result, err
		}

		var permissionIds []string
		for _, rp := range rolePermissions {
			permissionIds = append(permissionIds, rp.PermissionId)
		}

		if len(permissionIds) > 0 {
			// 查询具体的权限详情 (且状态开启)
			permissionDataList, err = g.CRMPermission.WithContext(ctx).Where(
				g.CRMPermission.PermissionId.In(permissionIds...),
				g.CRMPermission.Status.Eq(model.StatusOn),
			).Order(g.CRMPermission.Position.Desc(), g.CRMPermission.Id.Desc()).Find()
			if err != nil {
				return result, err
			}
		}
	}

	// 3. 构建树形结构 (复用 PermissionList 的构建逻辑)
	// 将 model 转换为 resp 结构
	idMap := make(map[string]*RespPermissionService, len(permissionDataList))
	for _, p := range permissionDataList {
		idMap[p.PermissionId] = &RespPermissionService{
			PermissionId:   p.PermissionId,
			PermissionName: p.PermissionName,
			PermissionUrl:  p.PermissionURL,
			ParentId:       p.ParentId,
			Status:         p.Status,
			IsInit:         p.IsInit,
			Position:       p.Position,
			CreatedAt:      p.CreatedAt.Format("2006-01-02 15:04:05"),
			ChildList:      []*RespPermissionService{},
		}
	}

	var roots []*RespPermissionService
	for _, p := range permissionDataList {
		node := idMap[p.PermissionId]
		// 如果是根节点 (ParentId 为空)
		if p.ParentId == "" {
			roots = append(roots, node)
			continue
		}
		// 尝试挂载到父节点
		if parent, ok := idMap[p.ParentId]; ok && parent != nil {
			parent.ChildList = append(parent.ChildList, node)
		} else {
			// 特殊情况：如果是普通用户，可能只分配了子菜单而没有分配父菜单
			// 这种情况下，是否需要显示？通常如果不显示父菜单，子菜单也无法在树中找到位置
			// 策略：如果找不到父节点，且该用户不是超管（即部分权限），
			// 可能需要前端处理，或者后端将其作为顶层？
			// 根据需求 "对应层级展示"，通常意味着父节点必须存在。
			// 如果数据库中存在父节点但未分配给用户，则用户看不到该树枝。
			// 如果是数据不一致（父节点ID错误），则也无法挂载。
			// 这里保持与 PermissionList 一致逻辑：不挂载则不显示（或者作为游离节点，但 PermissionList 注释掉了 roots append）
		}
	}

	if len(roots) == 0 {
		roots = []*RespPermissionService{}
	}

	// 序列化结果以便存入 result.Data (result.Data is interface{})
	// 直接赋值即可
	result.Data = map[string]any{"list": roots}
	result.SetMessage("操作成功")
	return result, nil
}

// 辅助结构体定义 (如果 impl_find.go 里定义了 RespPermissionFindInfo，这里可以复用，或者重新定义)
// 注意：同一个包下可以直接使用 impl_find.go 中的 RespPermissionFindInfo
