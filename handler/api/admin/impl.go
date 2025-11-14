package admin

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/admin"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g            *gin.RouterGroup
	adminService service.AdminIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:            g,
		adminService: admin.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	//g := h.g.Group("/admin")
	//g.POST("/create", h.CreateData)
	//g.POST("/update", h.UpdateData)
	//g.POST("/delete", h.DeleteData)
	//g.POST("/status", h.StatusData)
	//g.GET("/list", h.ListData)
}
