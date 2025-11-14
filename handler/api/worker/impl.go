package worker

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/demo"
	"crm/internal/service/worker"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g             *gin.RouterGroup
	workerService service.Worker
	demoService   service.Demo
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:             g,
		workerService: worker.NewService(),
		demoService:   demo.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/worker")
	g.POST("/put/task", h.WorkerPutTask)
}
