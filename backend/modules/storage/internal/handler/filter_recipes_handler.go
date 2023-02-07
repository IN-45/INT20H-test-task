package handler

import (
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FilterRecipesHandler struct {
	filterRecipesService *storage_service.FilterRecipesService
}

func NewFilterRecipesHandler(
	filterRecipesService *storage_service.FilterRecipesService,
) *FilterRecipesHandler {
	return &FilterRecipesHandler{
		filterRecipesService: filterRecipesService,
	}
}

func RegisterFilterRecipeHandler(app *fiber.App, h *FilterRecipesHandler) {
	app.Get("/filter-recipes", h.GetFilterRecipes)
}

// GetFilterRecipes
//
//	@Summary	Get filter recipes
//	@Tags		Recipe
//	@Success	200	{array}		storage_service.DtoFilterRecipe
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/filter-recipes [get]
func (h *FilterRecipesHandler) GetFilterRecipes(ctx *fiber.Ctx) error {
	userId := ctx.Cookies("user_id")
	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user id not found")
	}

	filterRecipes, err := h.filterRecipesService.FilterRecipes(ctx.Context(), uuid.MustParse(userId))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapServiceFilterRecipesToDto(filterRecipes))
}

type dtoFilterProduct struct {
	ProductId    uuid.UUID `json:"product_id"`
	Name         string    `json:"name"`
	Amount       int       `json:"amount"`
	MissedAmount int       `json:"missed_amount"`
	AmountType   string    `json:"amount_type"`
	ImageURL     string    `json:"image_url"`
	CategoryId   uuid.UUID `json:"category_id"`
}

type dtoFilterRecipe struct {
	Id                 uuid.UUID           `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	AuthorId           uuid.UUID           `json:"author_id"`
	CookingTimeMinutes int                 `json:"cooking_time_minutes"`
	ImageURL           string              `json:"image_url"`
	Instructions       []*dtoInstruction   `json:"instructions"`
	Products           []*dtoFilterProduct `json:"products"`
}
