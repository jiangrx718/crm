package aws

import (
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"web/gopkg/services"

	"github.com/gin-gonic/gin"
)

func (s *Service) AwsS3Upload(ctx *gin.Context, fileInfo *multipart.FileHeader) (services.Result, error) {
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

	fileContent, err := fileInfo.Open()
	if err != nil {
		return services.Failed(ctx, errors.New("file open error: "+err.Error()))
	}

	defer func() {
		_ = fileContent.Close()
	}()

	bytes, err := io.ReadAll(fileContent)
	if err != nil {
		return services.Failed(ctx, errors.New("file read error: "+err.Error()))
	}

	if bytes == nil {
		return services.Failed(ctx, errors.New("file header is missing"))
	}

	//source := files.FileName(fileInfo.Filename + fileInfo.LowerBaseExt())
	//path, err := storage.UploadByFileBytes(constant.ProjectName, constant.ContrastFileModuleName, source, fileContent)
	//if err != nil {
	//	log.Error(ctx, logPrefix, zap.Any("storage v2 upload by file reader error", err))
	//	return services.Failed(ctx, err)
	//}
	//
	//return services.Success(ctx, view.FileInfo{
	//	Name: source.Base(),
	//	Path: path,
	//})

	//storage.UploadByReader()
	//res, err := h.fileService.UploadFileProxy(ctx, files.FileName(fileInfo.Filename), bytes)
	//if err != nil {
	//	return services.Failed(ctx, errors.New("file upload error: "+err.Error()))
	//}
	//
	return services.Success(ctx, map[string]interface{}{
		"file_tag": "",
	})
}
