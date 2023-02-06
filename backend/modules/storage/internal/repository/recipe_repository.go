package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"

	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type RecipeRepository struct {
	db *bun.DB
}

func NewRecipeRepository(db *bun.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) Create(ctx context.Context, recipe *storage_model.Recipe) error {
	_, err := r.db.NewInsert().Model(recipe).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("recipe already exists")
	}

	return err
}

func (r *RecipeRepository) AddProducts(ctx context.Context, products []storage_model.RecipesProducts) error {
	_, err := r.db.NewInsert().Model(&products).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("recipe already exists")
	}

	return err
}
