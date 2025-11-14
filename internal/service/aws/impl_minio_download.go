package aws

import (
	"context"
	"web/gopkg/minio"
	"web/gopkg/services"
)

func (s *Service) AwsMinioDownload(ctx context.Context, objectName string) (services.Result, error) {
	downloadedFile := "downloaded-" + objectName
	if err := minio.DownloadFile(objectName, downloadedFile); err != nil {
		return nil, err
	}

	return services.Success(ctx, map[string]interface{}{
		"download": downloadedFile,
	})
}

func (s *Service) AwsMinioDownloadFile(ctx context.Context, objectKey string) (services.Result, error) {
	bytes, err := minio.DownloadFileByte(objectKey)
	if err != nil {
		return nil, err
	}

	return services.Success(ctx, map[string]interface{}{
		"bytes": bytes,
	})
}
