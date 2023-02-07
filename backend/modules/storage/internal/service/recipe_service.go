package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/IN-45/INT20H-test-task/pkg/db"
	"github.com/google/uuid"
)

type Instructions struct {
	RecipeId    uuid.UUID
	Description string
	Priority    int
}

type RecipeProducts struct {
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
	RecipeProducts     []RecipeProducts
}

type Recipe struct {
	Recipe       *storage_model.Recipe
	Instructions []*storage_model.Instruction
	Products     []*storage_model.RecipeProducts
}

type RecipeService struct {
	recipeRepository      *storage_repository.RecipeRepository
	instructionRepository *storage_repository.InstructionRepository
	transactionManager    *db.TransactionManager
}

func NewRecipeService(
	recipeRepository *storage_repository.RecipeRepository,
	instructionRepository *storage_repository.InstructionRepository,
	transactionManager *db.TransactionManager,
) *RecipeService {
	return &RecipeService{
		recipeRepository:      recipeRepository,
		instructionRepository: instructionRepository,
		transactionManager:    transactionManager,
	}
}

func (s *RecipeService) AddRecipe(ctx context.Context, params RecipeParams) (uuid.UUID, error) {
	recipe := storage_model.NewRecipe(
		uuid.New(),
		params.Name,
		params.Description,
		params.AuthorId,
		params.CookingTimeMinutes,
		params.ImageURL,
	)

	if err := s.transactionManager.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := s.recipeRepository.Create(ctx, recipe); err != nil {
			return err
		}

		if err := s.createRecipeInstructions(ctx, recipe.Id, params); err != nil {
			return err
		}

		if err := s.createRecipeProducts(ctx, recipe.Id, params); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return uuid.Nil, err
	}

	return recipe.Id, nil
}

func (s *RecipeService) GetAllRecipes(ctx context.Context) ([]*Recipe, error) {
	var recipes []*Recipe

	allRecipes, err := s.recipeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, recipe := range allRecipes {
		instructions, err := s.instructionRepository.GetByRecipeId(ctx, recipe.Id)
		if err != nil {
			return nil, err
		}

		products, err := s.recipeRepository.GetProducts(ctx, recipe.Id)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, &Recipe{
			Recipe:       recipe,
			Instructions: instructions,
			Products:     products,
		})
	}

	return recipes, nil
}

func (s *RecipeService) GetById(ctx context.Context, id uuid.UUID) (*Recipe, error) {
	recipe, err := s.recipeRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	instructions, err := s.instructionRepository.GetByRecipeId(ctx, recipe.Id)
	if err != nil {
		return nil, err
	}

	products, err := s.recipeRepository.GetProducts(ctx, recipe.Id)
	if err != nil {
		return nil, err
	}

	return &Recipe{
		Recipe:       recipe,
		Instructions: instructions,
		Products:     products,
	}, nil
}

func (s *RecipeService) createRecipeInstructions(ctx context.Context, recipeId uuid.UUID, params RecipeParams) error {
	var instructions []storage_model.Instruction

	for _, instruction := range params.Instructions {
		instructions = append(instructions, *storage_model.NewInstruction(
			recipeId,
			instruction.Description,
			instruction.Priority,
		))
	}

	if err := s.instructionRepository.Create(ctx, instructions); err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) createRecipeProducts(ctx context.Context, recipeId uuid.UUID, params RecipeParams) error {
	var RecipeProducts []storage_model.RecipeProducts

	for _, product := range params.RecipeProducts {
		RecipeProducts = append(RecipeProducts, *storage_model.NewRecipeProducts(
			recipeId,
			product.ProductId,
			product.Amount,
			product.AmountType,
		))
	}

	if err := s.recipeRepository.AddProducts(ctx, RecipeProducts); err != nil {
		return err
	}

	return nil
}
