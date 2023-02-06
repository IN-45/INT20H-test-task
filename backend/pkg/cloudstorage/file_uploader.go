package cloudstorage

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type FileUploader struct {
	bucketName    string
	storageClient *storage.Client
}

func NewFileUploader(
	cfg Config,
	storageClient *storage.Client,
) *FileUploader {
	return &FileUploader{
		bucketName:    cfg.ImagesBucketName,
		storageClient: storageClient,
	}
}

func (u *FileUploader) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	sw := u.storageClient.Bucket(u.bucketName).Object(file.Filename).NewWriter(ctx)

	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(sw, fileContent); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	return sw.Attrs().MediaLink, nil
}
