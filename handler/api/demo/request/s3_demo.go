package request

type ShowS3Demo struct {
	FilePath string `json:"file_path" binding:"required"` // 文件路径
}
