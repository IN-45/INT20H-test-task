package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/google/uuid"
)

type Instructions struct {
	RecipeId    uuid.UUID
	Description string
	Priority    int
}

type RecipesProducts struct {
	RecipeId   uuid.UUID
	ProductId  uuid.UUID
	Amount     int
	AmountType string
}

type RecipeParams struct {
	Name               string
	Description        string
	AuthorId           uuid.UUID
	CookingTimeMinutes int
	ImageURL           string
	Instructions       []Instructions
	RecipesProducts    []RecipesProducts
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

	var instructions []storage_model.Instruction
	for i := range params.Instructions {
		params.Instructions[i].RecipeId = recipe.Id
		instructions = append(instructions, *storage_model.NewInstruction(
			params.Instructions[i].RecipeId,
			params.Instructions[i].Description,
			params.Instructions[i].Priority,
		))
	}
	if err := s.instructionRepository.AddInstructions(ctx, instructions); err != nil {
		return uuid.Nil, err
	}

	var recipesProducts []storage_model.RecipesProducts
	for i := range params.RecipesProducts {
		params.RecipesProducts[i].RecipeId = recipe.Id
		recipesProducts = append(recipesProducts, *storage_model.NewRecipesProducts(
			params.RecipesProducts[i].RecipeId,
			params.RecipesProducts[i].ProductId,
			params.RecipesProducts[i].Amount,
			params.RecipesProducts[i].AmountType,
		))
	}
	if err := s.recipeRepository.AddProducts(ctx, recipesProducts); err != nil {
		return uuid.Nil, err
	}

	return recipe.Id, nil
}
