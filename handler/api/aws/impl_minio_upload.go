package aws

import (
	"web/gopkg/gins"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AwsMinioUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	if file == nil {
		gins.ServerError(ctx, err)
		return
	}

	res, err := h.awsService.AwsMinioUpload(ctx, file)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
