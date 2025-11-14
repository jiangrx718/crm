package request

type AwsMinioDownload struct {
	ObjectName string `json:"object_name" binding:"required"`
}

type AwsMinioDownloadFile struct {
	ObjectKey string `json:"object_key" binding:"required"`
}

type AwsMinioPreview struct {
	ObjectName string `json:"object_name" binding:"required"`
}
