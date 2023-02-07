package repository

import (
	"context"
	"database/sql"

	"github.com/IN-45/INT20H-test-task/pkg/db"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/google/uuid"

	"github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun/driver/pgdriver"
)

type RecipeRepository struct {
	db *db.TransactionRepository
}

func NewRecipeRepository(db *db.TransactionRepository) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) Create(ctx context.Context, recipe *storage_model.Recipe) error {
	_, err := r.db.NewInsert(ctx).Model(recipe).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("recipe already exists")
	}

	return err
}

func (r *RecipeRepository) AddProducts(ctx context.Context, products []storage_model.RecipeProducts) error {
	_, err := r.db.NewInsert(ctx).Model(&products).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("recipe already exists")
	}

	return err
}

func (r *RecipeRepository) GetAll(ctx context.Context) ([]*storage_model.Recipe, error) {
	var recipes []*storage_model.Recipe

	err := r.db.NewSelect(ctx).
		Model(&recipes).
		Scan(ctx)

	return recipes, err
}

func (r *RecipeRepository) GetProducts(ctx context.Context, recipeId uuid.UUID) ([]*storage_model.RecipeProducts, error) {
	var products []*storage_model.RecipeProducts

	err := r.db.NewSelect(ctx).
		Model(&products).
		Where("recipe_id = ?", recipeId).
		Relation("Product").
		Scan(ctx)

	return products, err
}

func (r *RecipeRepository) GetById(ctx context.Context, recipeId uuid.UUID) (*storage_model.Recipe, error) {
	recipe := new(storage_model.Recipe)

	err := r.db.NewSelect(ctx).
		Model(recipe).
		Where("id = ?", recipeId).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.NewNotFoundError("recipe not found")
		}
		return nil, err
	}

	return recipe, nil
}
