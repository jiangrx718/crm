package request

import "web/internal/model"

type CreateDemoParams struct {
	Name        string             `json:"name" binding:"required"`         // 标题
	FileType    int                `json:"file_type" binding:"required"`    // 类型
	ProjectType int                `json:"project_type" binding:"required"` // 类型
	Metadata    model.DemoMetadata `json:"metadata"`
}
