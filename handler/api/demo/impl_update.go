package demo

import (
	"web/gopkg/gins"
	"web/handler/api/demo/request"

	"github.com/gin-gonic/gin"
)

// UpdateDemo 更新
// @Tags APi:示例
// @Summary 更新
// @Description 更新
// @Accept json
// @Produce	json
// @Param param body request.UpdateDemoParams true "请求参数"
// @Success 200 {object} services.BaseResult{data=view.WritingOnline} "操作成功"
// @Router /api/demo/update [post]
func (h *Handler) UpdateDemo(ctx *gin.Context) {
	var req request.UpdateDemoParams
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.demoService.UpdateDemo(ctx, req.DemoId, req.Content)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
