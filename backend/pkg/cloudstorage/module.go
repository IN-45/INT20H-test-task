package cloudstorage

import (
	"context"

	"cloud.google.com/go/storage"
	"go.uber.org/fx"
	"google.golang.org/api/option"
)

var Module = fx.Provide(func(cfg Config) (*FileUploader, error) {
	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile(cfg.CloudCredentialsPath))
	if err != nil {
		return nil, err
	}

	return NewFileUploader(cfg, storageClient), nil
})
