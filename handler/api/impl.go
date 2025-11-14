package api

import (
	"web/gopkg/gins"
	"web/handler/api/aws"
	"web/handler/api/demo"
	"web/handler/api/markdown"
	"web/handler/api/worker"
	"web/handler/api/workflow"
	"web/handler/middleware"

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
		markdown.NewHandler(g),
		aws.NewHandler(g),
		worker.NewHandler(g),
		workflow.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
