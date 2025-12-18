package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginUserId, exist := ctx.Get("user_id")
		if !exist {

		}
		fmt.Printf("loginUserId:%#v\n", loginUserId)
	}
}
