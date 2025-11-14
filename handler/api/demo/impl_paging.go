package demo

import (
	"web/gopkg/gins"
	"web/handler/api/demo/request"

	"github.com/gin-gonic/gin"
)

// PagingDemo 列表-分页
// @Tags APi:示例
// @Summary 列表-分页
// @Description 列表-分页
// @Produce	json
// @Param param query request.PagingDemoParams true "请求参数"
// @Success 200 {object} services.BaseResult{data=view.Paging[view.WritingKnowledge]{}} "操作成功"
// @Router /api/demo/list [get]
func (h *Handler) PagingDemo(ctx *gin.Context) {
	var req request.PagingDemoParams

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.demoService.PagingDemo(ctx, req.Page)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
