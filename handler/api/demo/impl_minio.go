package demo

import (
	"fmt"
	"os"
	"path/filepath"
	"web/gopkg/gins"
	"web/gopkg/minio"
	"web/gopkg/services"
	"web/gopkg/utils/date"
	"web/gopkg/utils/files"

	"github.com/gin-gonic/gin"
)

// UploadFile 文件上传代理
// @Tags 客户侧：文件管理
// @Summary 文件上传代理
// @Description 文件上传代理
// @Accept multipart/form-data
// @Produce	json
// @Param param body request.UploadFileProxy true "请求参数"
// @Success 200 {object} services.BaseResult{data=view.FileInfo} "操作成功"
// @Router /api/demo/s3/upload [post]
func (h *Handler) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("------------1111111-------")
		gins.BadRequest(ctx, err)
		return
	}

	if file == nil {
		fmt.Println("------------22222222-------")
		gins.ServerError(ctx, err)
		return
	}

	// 文件扩展名验证（白名单机制）
	//allowedExtensions := map[string]bool{
	//	".doc": true, ".docx": true, // Word文档
	//	".xls": true, ".xlsx": true, // Excel文档
	//	".pdf": true, // PDF文档
	//}
	//
	//ext := strings.ToLower(filepath.Ext(file.Filename))
	//if !allowedExtensions[ext] {
	//	gins.ServerError(ctx, err)
	//	return
	//}

	// 创建本地临时目录（如果不存在）
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		gins.ServerError(ctx, err)
		return
	}

	// 生成唯一的本地文件路径
	objectName := date.GetCurrentDateYearMonthDayHIS() + filepath.Ext(file.Filename)
	localFilePath := filepath.Join(uploadDir, objectName)

	// 将上传的文件保存到本地
	if err := ctx.SaveUploadedFile(file, localFilePath); err != nil {
		gins.ServerError(ctx, err)
		return
	}

	defer files.DeleteFileIfExists(localFilePath)
	_, err = minio.UploadFile(objectName, localFilePath)
	if err != nil {
		fmt.Println("------------33333333-------")
		gins.ServerError(ctx, err)
		return
	}

	var res services.Result
	gins.StatusOK(ctx, res)
}
