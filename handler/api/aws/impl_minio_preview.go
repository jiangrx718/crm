package aws

import (
	"web/gopkg/gins"
	"web/handler/api/aws/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AwsMinioPreview(ctx *gin.Context) {
	var req request.AwsMinioPreview
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.awsService.AwsMinioPreview(ctx, req.ObjectName)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
