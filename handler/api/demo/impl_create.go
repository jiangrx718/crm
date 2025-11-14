package demo

import (
	"web/gopkg/gins"
	"web/handler/api/demo/request"

	"github.com/gin-gonic/gin"
)

// CreateDemo 创建
// @Tags APi:示例
// @Summary 创建
// @Description 创建
// @Accept json
// @Produce	json
// @Param param body request.CreateDemoParams true "请求参数"
// @Success 200 {object} services.BaseResult{data=view.WritingOnline} "操作成功"
// @Router /api/demo/create [post]
func (h *Handler) CreateDemo(ctx *gin.Context) {
	var req request.CreateDemoParams
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.demoService.CreateDemo(ctx, req.Name, req.FileType, req.ProjectType, req.Metadata)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, result)
}
