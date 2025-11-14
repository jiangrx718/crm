package request

type UpdateDemoParams struct {
	DemoId  string `json:"demo_id" binding:"required"` // 案例ID
	Content string `json:"content" binding:"required"` // 文档内容
}
