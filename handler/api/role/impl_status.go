package role

import (
	"crm/gopkg/gins"
	"crm/handler/api/role/request"

	"github.com/gin-gonic/gin"
)

// RoleStatus 管理权限-角色管理-状态
func (h *Handler) RoleStatus(ctx *gin.Context) {
	var req request.RoleStatusReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.roleService.RoleStatus(ctx, req.RoleId, req.Status)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
