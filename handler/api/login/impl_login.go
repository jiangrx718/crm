package login

import (
	"crm/gopkg/gins"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// DoLogin 用户登录
func (h *Handler) DoLogin(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.loginService.Login(ctx, req.UserName, req.Password)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
