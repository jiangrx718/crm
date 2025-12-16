package api

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin"
	"crm/handler/api/admin_role"
	"crm/handler/api/article"
	"crm/handler/api/article_category"
	"crm/handler/api/login"
	"crm/handler/api/order"
	"crm/handler/api/permission"
	"crm/handler/api/product"
	"crm/handler/api/product_category"
	"crm/handler/api/role"
	"crm/handler/api/system"
	"crm/handler/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

func NewHandler(engine *gin.Engine) gins.Handler {
	return &Handler{
		engine: engine,
	}
}

func (h *Handler) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append([]string{}, config.AllowHeaders...)
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))

	// 登录路由
	l := h.engine.Group("/")
	loginHandlers := []gins.Handler{
		login.NewHandler(l),
	}
	for _, handler := range loginHandlers {
		handler.RegisterRoutes()
	}

	// 业务路由
	g := h.engine.Group("/api", middleware.RequestCapture(), middleware.JWTAuth())
	handlers := []gins.Handler{
		admin.NewHandler(g),
		admin_role.NewHandler(g),
		article.NewHandler(g),
		article_category.NewHandler(g),
		order.NewHandler(g),
		permission.NewHandler(g),
		product.NewHandler(g),
		product_category.NewHandler(g),
		role.NewHandler(g),
		system.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
