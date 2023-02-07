package repository

import (
	"context"

	"github.com/IN-45/INT20H-test-task/pkg/db"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun/driver/pgdriver"
)

type CategoryRepository struct {
	db *db.TransactionRepository
}

func NewCategoryRepository(db *db.TransactionRepository) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]*storage_model.Category, error) {
	var categories []*storage_model.Category

	err := r.db.NewSelect(ctx).Model(&categories).Scan(ctx)

	return categories, err
}

func (r *CategoryRepository) Create(ctx context.Context, category *storage_model.Category) error {
	_, err := r.db.NewInsert(ctx).Model(category).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("category already exists")
	}

	return err
}
