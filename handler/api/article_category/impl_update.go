package article_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/article_category/request"

	"github.com/gin-gonic/gin"
)

// CategoryUpdate 内容管理-文章分类-更新
func (h *Handler) CategoryUpdate(ctx *gin.Context) {
	var req request.ArticleCategoryUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleCategoryService.CategoryUpdate(ctx, req.CategoryId, req.ParentId, req.CategoryName, req.CategoryImage, req.Status, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
