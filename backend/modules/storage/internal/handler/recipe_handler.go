package handler

import (
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	app.Get("/recipe", h.GetAll)
	app.Get("/recipe/:id", h.GetById)
}

// Create
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

	userId := ctx.Cookies("user_id")
	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user id not found")
	}

	dto.AuthorId = userId

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.recipeService.AddRecipe(ctx.Context(), createRecipeParams(dto))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(id)
}

// GetAll
//
//	@Summary	Get all recipes
//	@Tags		Recipe
//	@Success	200	{array}		dtoRecipe
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/recipe [get]
func (h *RecipeHandler) GetAll(ctx *fiber.Ctx) error {
	recipes, err := h.recipeService.GetAllRecipes(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapServiceRecipesToDto(recipes))
}

// GetById
//
//	@Summary	Get recipe by id
//	@Tags		Recipe
//	@Param		id	path		string	true	"Recipe ID"
//	@Success	200	{array}		dtoRecipe
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/recipe/{id} [get]
func (h *RecipeHandler) GetById(ctx *fiber.Ctx) error {
	dto := new(dtoGetRecipeById)

	if err := ctx.ParamsParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	recipe, err := h.recipeService.GetById(ctx.Context(), uuid.MustParse(dto.Id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapServiceRecipeToDto(recipe))
}

type dtoGetRecipeById struct {
	Id string `params:"id" validate:"required,uuid4"`
}

type dtoCreateRecipeProducts struct {
	ProductId  string `json:"product_id" validate:"required,uuid4"`
	Amount     int    `json:"amount" validate:"required"`
	AmountType string `json:"amount_type" validate:"required"`
}

type dtoCreateInstruction struct {
	Description string `json:"description"`
	Priority    int    `json:"priority" validate:"required"`
}

type dtoCreateRecipe struct {
	Name               string                    `json:"name" validate:"required"`
	Description        string                    `json:"description"`
	AuthorId           string                    `json:"-" validate:"uuid4"`
	CookingTimeMinutes int                       `json:"cooking_time_minutes"`
	ImageURL           string                    `json:"image_url"`
	Instructions       []dtoCreateInstruction    `json:"instructions,omitempty"`
	Products           []dtoCreateRecipeProducts `json:"products,omitempty"`
}

type dtoInstruction struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type dtoRecipeProduct struct {
	ProductId  uuid.UUID `json:"product_id"`
	Name       string    `json:"name"`
	Amount     int       `json:"amount"`
	AmountType string    `json:"amount_type"`
	ImageURL   string    `json:"image_url"`
	CategoryId uuid.UUID `json:"category_id"`
}

type dtoRecipe struct {
	Id                 uuid.UUID           `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	AuthorId           uuid.UUID           `json:"author_id"`
	CookingTimeMinutes int                 `json:"cooking_time_minutes"`
	ImageURL           string              `json:"image_url"`
	Instructions       []*dtoInstruction   `json:"instructions"`
	Products           []*dtoRecipeProduct `json:"products"`
}
