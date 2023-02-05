package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type CategoryRepository struct {
	db *bun.DB
}

func NewCategoryRepository(db *bun.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]*storage_model.Category, error) {
	var categories []*storage_model.Category

	err := r.db.NewSelect().Model(&categories).Scan(ctx)

	return categories, err
}

func (r *CategoryRepository) Create(ctx context.Context, category *storage_model.Category) error {
	_, err := r.db.NewInsert().Model(category).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("category already exists")
	}

	return err
}
