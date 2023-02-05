package handler

import (
	"github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	validator          *validator.Validate
	categoryRepository *repository.CategoryRepository
	categoryService    *service.CategoryService
}

func NewCategoryHandler(
	validator *validator.Validate,
	categoryRepository *repository.CategoryRepository,
	categoryService *service.CategoryService,
) *CategoryHandler {
	return &CategoryHandler{
		validator:          validator,
		categoryRepository: categoryRepository,
		categoryService:    categoryService,
	}
}

func RegisterCategoryHandler(app *fiber.App, h *CategoryHandler) {
	app.Get("/categories", h.GetAll)
	app.Post("/category", h.Create)
}

// GetAll
//
//	@Summary	Get all categories
//	@Tags		Category
//	@Success	200	{array}		dtoCategory
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/categories [get]
func (h *CategoryHandler) GetAll(ctx *fiber.Ctx) error {
	categories, err := h.categoryRepository.GetAll(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapDtoCategories(categories))
}

// Create
//
//	@Summary	Create new category
//	@Tags		Category
//	@Param		input	body		dtoCreateCategory	true	"Category"
//	@Success	200		{object}	uuid.UUID
//	@Failure	400		{object}	fiber.Error
//	@Failure	401		{object}	customerrors.UnauthorizedError
//	@Failure	500		{object}	fiber.Error
//	@Router		/category [post]
func (h *CategoryHandler) Create(ctx *fiber.Ctx) error {
	dto := new(dtoCreateCategory)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.categoryService.Create(ctx.Context(), dto.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(id)
}

type dtoCategory struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type dtoCreateCategory struct {
	Name string `json:"name" validate:"required"`
}
