package markdown

import (
	"web/gopkg/gins"
	"web/internal/service"
	"web/internal/service/markdown"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g               *gin.RouterGroup
	markdownService service.Markdown
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:               g,
		markdownService: markdown.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/markdown")
	g.GET("/file", h.MarkdownExtractFile)
	g.GET("/title", h.MarkdownExtractTitle)
	g.GET("/section", h.MarkdownExtractSection)
	g.GET("/title-section", h.MarkdownExtractTitleSection)
}
