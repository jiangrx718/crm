package aws

import (
	"web/gopkg/gins"
	"web/internal/service"
	"web/internal/service/aws"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g          *gin.RouterGroup
	awsService service.Aws
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:          g,
		awsService: aws.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/aws")
	g.POST("/s3/upload", h.AwsS3Upload)
	//g.GET("/s3/download", h.AwsS3Download)
	g.POST("/minio/upload", h.AwsMinioUpload)
	g.POST("/minio/preview", h.AwsMinioPreview)
	g.POST("/minio/download", h.AwsMinioDownload)
	g.POST("/minio/download/file", h.AwsMinioDownloadFile)
}
