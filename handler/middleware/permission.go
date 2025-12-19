package middleware

import (
	"crm/gopkg/gins"
	"crm/internal/service/permission"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginUserId, exist := ctx.Get("user_id")
		if !exist {
			gins.Unauthorized(ctx)
			ctx.Abort()
			return
		}

		permissionService := permission.NewService()
		result, err := permissionService.PermissionFind(ctx, loginUserId.(string))
		if err != nil {
			gins.ServerError(ctx, err)
			ctx.Abort()
			return
		}

		fmt.Printf("result: %+v\n", result)

		ctx.Next()
	}
}
