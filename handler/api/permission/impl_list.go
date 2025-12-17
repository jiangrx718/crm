package permission

import (
	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

// PermissionList 管理权限-权限设置-列表
func (h *Handler) PermissionList(ctx *gin.Context) {

	result, err := h.permissionService.PermissionList(ctx)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
