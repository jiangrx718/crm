package aws

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"web/gopkg/minio"
	"web/gopkg/services"
	"web/gopkg/utils/date"
	"web/gopkg/utils/files"

	"github.com/gin-gonic/gin"
)

func (s *Service) AwsMinioUpload(ctx *gin.Context, fileInfo *multipart.FileHeader) (services.Result, error) {
	// 文件扩展名验证（白名单机制）
	allowedExtensions := map[string]bool{
		".doc": true, ".docx": true, // Word文档
		".xls": true, ".xlsx": true, // Excel文档
		".pdf": true,                                                                        // PDF文档
		".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".bmp": true, ".svg": true, // 图片
		".zip": true, ".tar.gz": true, ".tar.xz": true, ".gz": true, ".bz2": true, // 压缩文件
	}

	ext := strings.ToLower(filepath.Ext(fileInfo.Filename))
	if !allowedExtensions[ext] {
		return services.Failed(ctx, errors.New("file extension not allowed"))
	}

	// 创建本地临时目录（如果不存在）
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return services.Failed(ctx, err)
	}

	// 生成唯一的本地文件路径
	objectName := date.GetCurrentDateYearMonthDayHIS() + filepath.Ext(fileInfo.Filename)
	localFilePath := filepath.Join(uploadDir, objectName)

	// 将上传的文件保存到本地
	if err := ctx.SaveUploadedFile(fileInfo, localFilePath); err != nil {
		return services.Failed(ctx, err)
	}

	defer files.DeleteFileIfExists(localFilePath)
	_, err := minio.UploadFile(objectName, localFilePath)
	if err != nil {
		return services.Failed(ctx, err)
	}

	return services.Success(ctx, map[string]interface{}{
		"file_tag": objectName,
	})
}
