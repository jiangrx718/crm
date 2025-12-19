package permission

import (
	"context"
	"crm/gopkg/cache/redis"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// PermissionFind 查询权限
func (s *Service) PermissionFind(ctx context.Context, adminId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("login_auth:%s", adminId)
	rdb, err := redis.ClientAndErr("web")
	if err == nil && rdb != nil {
		if val, err := rdb.Get(ctx, cacheKey).Result(); err == nil && val == "1" {
			// 缓存命中，直接返回成功
			result.SetMessage("操作成功")
			return result, nil
		}
	}

	logObj.Info("用户接收参数：", zap.String("adminId", adminId))

	// 查询数据库验证用户状态
	admin, err := g.CRMAdmin.WithContext(ctx).Where(
		g.CRMAdmin.AdminId.Eq(adminId),
		g.CRMAdmin.Status.Eq(model.StatusOn),
	).Take()

	if err != nil {
		return result, fmt.Errorf("user not found or disabled")
	}
	if admin == nil {
		return result, fmt.Errorf("user not found or disabled")
	}

	// 查询用户所处的角色

	// 写入缓存，有效期 30 分钟
	if rdb != nil {
		rdb.Set(ctx, cacheKey, "1", 30*time.Minute)
	}

	result.SetMessage("操作成功")
	return result, nil
}
