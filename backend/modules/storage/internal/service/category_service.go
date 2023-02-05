package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	categoryRepository *storage_repository.CategoryRepository
}

func NewCategoryService(
	categoryRepository *storage_repository.CategoryRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, name string) (uuid.UUID, error) {
	category := storage_model.NewCategory(
		uuid.New(),
		name,
	)

	if err := s.categoryRepository.Create(ctx, category); err != nil {
		return uuid.Nil, err
	}

	return category.Id, nil
}
