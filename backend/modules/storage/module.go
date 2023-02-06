package storage

import (
	storage_handler "github.com/IN-45/INT20H-test-task/modules/storage/internal/handler"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage_handler.NewProductHandler,
		storage_handler.NewCategoryHandler,
		storage_handler.NewUploaderHandler,
		storage_service.NewProductService,
		storage_service.NewCategoryService,
		storage_repository.NewProductRepository,
		storage_repository.NewCategoryRepository,
	),
	fx.Invoke(
		storage_handler.RegisterProductHandler,
		storage_handler.RegisterCategoryHandler,
		storage_handler.RegisterUploaderHandler,
	),
)
