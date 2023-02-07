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
	for _, instruction := range params.Instructions {
		instructions = append(instructions, *storage_model.NewInstruction(
			recipe.Id,
			instruction.Description,
			instruction.Priority,
		))
	}
	if err := s.instructionRepository.AddInstructions(ctx, instructions); err != nil {
		return uuid.Nil, err
	}

	var recipesProducts []storage_model.RecipesProducts
	for _, product := range params.RecipesProducts {
		recipesProducts = append(recipesProducts, *storage_model.NewRecipesProducts(
			recipe.Id,
			product.ProductId,
			product.Amount,
			product.AmountType,
		))
	}
	if err := s.recipeRepository.AddProducts(ctx, recipesProducts); err != nil {
		return uuid.Nil, err
	}

	return recipe.Id, nil
}

func (s *RecipeService) GetAllRecipes(ctx context.Context) ([]DtoRecipe, error) {
	var dtoRecipes []DtoRecipe

	recipes, err := s.recipeRepository.GetAllRecipes(ctx)
	if err != nil {
		return nil, err
	}

	for _, recipe := range recipes {
		dtoRecipe, err := s.fillDtoRecipe(ctx, recipe)
		if err != nil {
			return nil, err
		}

		dtoRecipes = append(dtoRecipes, *dtoRecipe)
	}

	return dtoRecipes, nil
}

func (s *RecipeService) GetRecipeById(ctx context.Context, id uuid.UUID) (*DtoRecipe, error) {
	recipe, err := s.recipeRepository.GetRecipeById(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.fillDtoRecipe(ctx, recipe)
}

func (s *RecipeService) fillDtoRecipe(ctx context.Context, recipe *storage_model.Recipe) (*DtoRecipe, error) {
	dtoRecipe := new(DtoRecipe)
	// get instructions for current recipe
	instructions, err := s.instructionRepository.GetInstructionsById(ctx, recipe.Id)
	if err != nil {
		return nil, err
	}
	var dtoInstructions []dtoInstruction
	for _, instruction := range instructions {
		dtoInstructions = append(dtoInstructions, dtoInstruction{
			Description: instruction.Description,
			Priority:    instruction.Priority,
		})
	}
	// get products for current recipe
	recipeProducts, err := s.recipeRepository.GetProductsForRecipe(ctx, recipe.Id)
	if err != nil {
		return nil, err
	}
	var dtoProducts []dtoProduct
	for _, product := range recipeProducts {
		dtoProducts = append(dtoProducts, dtoProduct{
			ProductId:  product.ProductId,
			Name:       product.Product.Name,
			Amount:     product.Amount,
			AmountType: product.AmountType,
			ImageURL:   product.Product.ImageURL,
			CategoryId: product.Product.CategoryId,
		})
	}

	dtoRecipe.Id = recipe.Id
	dtoRecipe.Name = recipe.Name
	dtoRecipe.Description = recipe.Description
	dtoRecipe.AuthorId = recipe.AuthorId
	dtoRecipe.CookingTimeMinutes = recipe.CookingTimeMinutes
	dtoRecipe.ImageURL = recipe.ImageURL
	dtoRecipe.Instructions = dtoInstructions
	dtoRecipe.Products = dtoProducts

	return dtoRecipe, nil
}

type dtoInstruction struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type dtoProduct struct {
	ProductId  uuid.UUID `json:"product_id"`
	Name       string    `json:"name"`
	Amount     int       `json:"amount"`
	AmountType string    `json:"amount_type"`
	ImageURL   string    `json:"image_url"`
	CategoryId uuid.UUID `json:"category_id"`
}

type DtoRecipe struct {
	Id                 uuid.UUID        `json:"id"`
	Name               string           `json:"name"`
	Description        string           `json:"description"`
	AuthorId           uuid.UUID        `json:"author_id"`
	CookingTimeMinutes int              `json:"cooking_time_minutes"`
	ImageURL           string           `json:"image_url"`
	Instructions       []dtoInstruction `json:"instructions"`
	Products           []dtoProduct     `json:"products"`
}
