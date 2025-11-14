package aws

import (
	"context"
	"time"
	"web/gopkg/minio"
	"web/gopkg/services"
)

func (s *Service) AwsMinioPreview(ctx context.Context, objectName string) (services.Result, error) {
	previewUrl, err := minio.GeneratePresignedURL(objectName, time.Hour)
	if err != nil {
		return nil, err
	}

	return services.Success(ctx, map[string]interface{}{
		"preview_url": previewUrl,
	})
}
