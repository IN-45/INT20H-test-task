package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/google/uuid"
)

type ProductParams struct {
	Name       string
	CategoryId uuid.UUID
	ImageURL   string
}

type ProductService struct {
	productRepository *storage_repository.ProductRepository
}

func NewProductService(
	productRepository *storage_repository.ProductRepository,
) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) Create(ctx context.Context, params ProductParams) (uuid.UUID, error) {
	product := storage_model.NewProduct(
		uuid.New(),
		params.CategoryId,
		params.Name,
		params.ImageURL,
	)

	if err := s.productRepository.Create(ctx, product); err != nil {
		return uuid.Nil, err
	}

	return product.Id, nil
}
