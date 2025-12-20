package article_category

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/article_category"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g                      *gin.RouterGroup
	articleCategoryService service.ArticleCategoryIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:                      g,
		articleCategoryService: article_category.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/article/category")
	g.POST("/create", h.CategoryCreate)
	g.POST("/update", h.CategoryUpdate)
	g.POST("/delete", h.CategoryDelete)
	g.GET("/list", h.CategoryList)
}
