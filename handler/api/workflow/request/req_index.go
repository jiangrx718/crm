package request

type WorkflowIndexReq struct {
	WorkflowID     string `json:"workflow_id" binding:"required"`
	Query          string `json:"query" binding:"required"`
	TaskType       string `json:"task_type" binding:"required"`
	TargetLanguage string `json:"target_language,omitempty"`
}
