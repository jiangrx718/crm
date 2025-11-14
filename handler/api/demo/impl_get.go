package demo

import (
	"web/gopkg/gins"
	"web/handler/api/demo/request"

	"github.com/gin-gonic/gin"
)

// GetDemo 详情
// @Tags APi:示例
// @Summary 详情
// @Description 详情
// @Produce	json
// @Param param query request.GetDemoParams false "请求参数"
// @Success 200 {object} services.BaseResult{data=view.WritingOnline} "操作成功"
// @Router /api/demo/get [get]
func (h *Handler) GetDemo(ctx *gin.Context) {
	var req request.GetDemoParams
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.demoService.GetDemo(ctx, req.DemoId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
