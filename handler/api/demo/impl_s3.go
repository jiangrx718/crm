package demo

import (
	"crm/gopkg/gins"
	"crm/gopkg/storage"
	"crm/handler/api/demo/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowS3Demo(ctx *gin.Context) {
	var req request.ShowS3Demo
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	fmt.Printf("文件S3路径: %+v\n", req.FilePath)
	bytes, _ := storage.DownloadFile(req.FilePath)

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Writer.Write(bytes)
}
