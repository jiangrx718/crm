package workflow

import (
	"web/gopkg/gins"
	"web/handler/middleware"

	"web/internal/service"
	"web/internal/service/workflow"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g               *gin.RouterGroup
	workflowService service.Workflow
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:               g,
		workflowService: workflow.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/workflow")
	g.POST("/chat", middleware.EventStreamHeadersMiddleware(), h.ChatStream)
}
