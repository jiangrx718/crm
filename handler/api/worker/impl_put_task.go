package worker

import (
	"web/gopkg/gins"
	"web/handler/api/demo/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) WorkerPutTask(ctx *gin.Context) {
	var req request.CreateDemoParams
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}
	res, err := h.demoService.CreateDemo(ctx, req.Name, req.FileType, req.ProjectType, req.Metadata)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}
	
	gins.StatusOK(ctx, res)
}
