package product_category

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/product_category"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g                      *gin.RouterGroup
	productCategoryService service.ProductCategoryIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:                      g,
		productCategoryService: product_category.NewService(),
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
