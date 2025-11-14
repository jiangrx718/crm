package request

type DeleteDemoParams struct {
	DemoId string `json:"demo_id" binding:"required"` // 案例ID
}
