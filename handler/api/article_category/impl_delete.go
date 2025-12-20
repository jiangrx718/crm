package article_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/article_category/request"

	"github.com/gin-gonic/gin"
)

// CategoryDelete 内容管理-文章分类-删除
func (h *Handler) CategoryDelete(ctx *gin.Context) {
	var req request.ArticleCategoryDeleteReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleCategoryService.CategoryDelete(ctx, req.CategoryId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
