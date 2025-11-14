package aws

import (
	"web/gopkg/gins"
	"web/handler/api/aws/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AwsMinioDownload(ctx *gin.Context) {
	var req request.AwsMinioDownload
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.awsService.AwsMinioDownload(ctx, req.ObjectName)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}

func (h *Handler) AwsMinioDownloadFile(ctx *gin.Context) {
	var req request.AwsMinioDownloadFile
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.awsService.AwsMinioDownloadFile(ctx, req.ObjectKey)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
