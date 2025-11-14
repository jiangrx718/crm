package service

import (
	"context"
	"mime/multipart"
	"web/gopkg/services"

	"github.com/gin-gonic/gin"
)

type Aws interface {
	AwsMinioUpload(ctx *gin.Context, fileInfo *multipart.FileHeader) (services.Result, error)
	AwsMinioPreview(ctx context.Context, objectName string) (services.Result, error)
	AwsMinioDownload(ctx context.Context, objectName string) (services.Result, error)
	AwsMinioDownloadFile(ctx context.Context, objectKey string) (services.Result, error)
	AwsS3Upload(ctx *gin.Context, fileInfo *multipart.FileHeader) (services.Result, error)
}
