package handler

import (
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RecipeHandler struct {
	validator             *validator.Validate
	recipeRepository      *storage_repository.RecipeRepository
	instructionRepository *storage_repository.RecipeRepository
	recipeService         *storage_service.RecipeService
}

func NewRecipeHandler(
	validator *validator.Validate,
	recipeRepository *storage_repository.RecipeRepository,
	instructionRepository *storage_repository.RecipeRepository,
	recipeService *storage_service.RecipeService,
) *RecipeHandler {
	return &RecipeHandler{
		validator:             validator,
		recipeRepository:      recipeRepository,
		instructionRepository: instructionRepository,
		recipeService:         recipeService,
	}
}

func RegisterRecipeHandler(app *fiber.App, h *RecipeHandler) {
	app.Post("/recipe", h.Create)
}

// CreateRecipe
//
//	@Summary	Create recipe and add products and instructions
//	@Tags		Recipe
//	@Param		input	body		dtoCreateRecipe	true	"Recipe"
//	@Success	200		{object}	uuid.UUID
//	@Failure	400		{object}	fiber.Error
//	@Failure	401		{object}	customerrors.UnauthorizedError
//	@Failure	500		{object}	fiber.Error
//	@Router		/recipe [post]
func (h *RecipeHandler) Create(ctx *fiber.Ctx) error {
	dto := new(dtoCreateRecipe)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	user_id := ctx.Cookies("user_id")
	if user_id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user id not found")
	}

	dto.AuthorId = user_id

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.recipeService.AddRecipe(ctx.Context(), createRecipeParams(dto))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(id)
}

type dtoCreateRecipesProducts struct {
	ProductId  string `json:"product_id" validate:"required,uuid4"`
	Amount     int    `json:"amount" validate:"required"`
	AmountType string `json:"amount_type" validate:"required"`
}

type dtoCreateInstruction struct {
	Description string `json:"description"`
	Priority    int    `json:"priority" validate:"required"`
}

type dtoCreateRecipe struct {
	Name               string                     `json:"name" validate:"required"`
	Description        string                     `json:"description"`
	AuthorId           string                     `json:"-" validate:"uuid4"`
	CookingTimeMinutes int                        `json:"cooking_time_minutes"`
	ImageURL           string                     `json:"image_url"`
	Instructions       []dtoCreateInstruction     `json:"instructions,omitempty"`
	Products           []dtoCreateRecipesProducts `json:"products,omitempty"`
}
