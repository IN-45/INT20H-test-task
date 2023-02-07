package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/google/uuid"

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

func (r *RecipeRepository) GetAllRecipes(ctx context.Context) ([]*storage_model.Recipe, error) {
	var recipes []*storage_model.Recipe

	err := r.db.NewSelect().
		Model(&recipes).
		Scan(ctx)

	print(recipes)
	return recipes, err
}

func (r *RecipeRepository) GetProductsForRecipe(ctx context.Context, recipeId uuid.UUID) ([]*storage_model.RecipesProducts, error) {
	var products []*storage_model.RecipesProducts

	err := r.db.NewSelect().
		Model(&products).
		Where("recipe_id = ?", recipeId).
		Relation("Product").
		Scan(ctx)

	return products, err
}

func (r *RecipeRepository) GetRecipeById(ctx context.Context, recipeId uuid.UUID) (*storage_model.Recipe, error) {
	recipe := new(storage_model.Recipe)

	err := r.db.NewSelect().
		Model(recipe).
		Where("id = ?", recipeId).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}
