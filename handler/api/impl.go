package api

import (
	"crm/gopkg/gins"
	"crm/handler/api/demo"
	"crm/handler/api/worker"
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

	g := h.engine.Group("/api", middleware.RequestCapture())
	handlers := []gins.Handler{
		demo.NewHandler(g),
		worker.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
