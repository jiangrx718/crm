package demo

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/demo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	demoService service.Demo
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		demoService: demo.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/demo")
	g.POST("/create", h.CreateDemo)
	g.GET("/get", h.GetDemo)
	g.GET("/list", h.PagingDemo)
	g.POST("/update", h.UpdateDemo)
	g.POST("/delete", h.DeleteDemo)
	g.POST("/s3/show", h.ShowS3Demo)
	g.POST("/s3/upload", h.UploadFile)
}
