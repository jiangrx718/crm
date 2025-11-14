package request

type WorkerPutTaskRequest struct {
	Name string `json:"name" required:"true"`
}
