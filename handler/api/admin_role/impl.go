package admin_role

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/admin_role"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g                *gin.RouterGroup
	adminRoleService service.AdminRoleIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:                g,
		adminRoleService: admin_role.NewService(),
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
