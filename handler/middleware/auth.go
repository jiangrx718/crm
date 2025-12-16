package middleware

import (
	"crm/gopkg/auth"
	"crm/gopkg/gins"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractToken(ctx)
		if token == "" {
			gins.Unauthorized(ctx)
			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil || claims == nil || claims.UserID == "" {
			gins.Unauthorized(ctx)
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}

func extractToken(ctx *gin.Context) string {
	// 优先解析 Authorization: Bearer <token>
	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return strings.TrimSpace(parts[1])
		}
	}

	// 兼容 query token
	if q := ctx.Query("token"); q != "" {
		return q
	}

	// 兼容 Cookie token
	if c, err := ctx.Cookie("token"); err == nil && c != "" {
		return c
	}
	return ""
}
