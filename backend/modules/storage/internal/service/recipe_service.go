package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/google/uuid"
)

type RecipeParams struct {
	Name               string
	Description         string
	AuthorId           uuid.UUID
	CookingTimeMinutes int
	ImageURL           string
	Instructions       []storage_model.Instruction
	RecipesProducts    []storage_model.RecipesProducts
}

type RecipeService struct {
	recipeRepository      *storage_repository.RecipeRepository
	instructionRepository *storage_repository.InstructionRepository
}

func NewRecipeService(
	recipeRepository *storage_repository.RecipeRepository,
	instructionRepository *storage_repository.InstructionRepository,
) *RecipeService {
	return &RecipeService{
		recipeRepository:      recipeRepository,
		instructionRepository: instructionRepository,
	}
}

func (s *RecipeService) AddRecipe(ctx context.Context, params RecipeParams) (uuid.UUID, error) {
	// TODO add transaction
	recipe := storage_model.NewRecipe(
		uuid.New(),
		params.Name,
		params.Description,
		params.AuthorId,
		params.CookingTimeMinutes,
		params.ImageURL,
	)

	if err := s.recipeRepository.Create(ctx, recipe); err != nil {
		return uuid.Nil, err
	}

	for i := range params.Instructions {
		params.Instructions[i].RecipeId = recipe.Id
	}
	if err := s.instructionRepository.AddInstructions(ctx, params.Instructions); err != nil {
		return uuid.Nil, err
	}

	for i := range params.RecipesProducts {
		params.RecipesProducts[i].RecipeId = recipe.Id
	}
	if err := s.recipeRepository.AddProducts(ctx, params.RecipesProducts); err != nil {
		return uuid.Nil, err
	}

	return recipe.Id, nil
}
