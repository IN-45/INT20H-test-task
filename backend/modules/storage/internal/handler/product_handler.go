package handler

import (
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	validator         *validator.Validate
	productRepository *storage_repository.ProductRepository
	productService    *storage_service.ProductService
}

func NewProductHandler(
	validator *validator.Validate,
	productRepository *storage_repository.ProductRepository,
	productService *storage_service.ProductService,
) *ProductHandler {
	return &ProductHandler{
		validator:         validator,
		productRepository: productRepository,
		productService:    productService,
	}
}

func RegisterProductHandler(app *fiber.App, h *ProductHandler) {
	app.Get("/products", h.GetAll)
	app.Get("/product/:id", h.GetById)
	app.Post("/product", h.Create)
	app.Get("/amount_types", h.GetAmountTypes)
}

// GetAll
//
//	@Summary	Get all products
//	@Tags		Product
//	@Success	200	{array}		dtoProduct
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/products [get]
func (h *ProductHandler) GetAll(ctx *fiber.Ctx) error {
	products, err := h.productRepository.GetAllWithCategories(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapDtoProducts(products))
}

// GetById
//
//	@Summary	Get product by id
//	@Tags		Product
//	@Param		id	path		string	true	"Product ID"
//	@Success	200	{array}		dtoProduct
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/product/{id} [get]
func (h *ProductHandler) GetById(ctx *fiber.Ctx) error {
	dto := new(dtoGetProductById)

	if err := ctx.ParamsParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product, err := h.productRepository.GetByIdWithCategory(ctx.Context(), uuid.MustParse(dto.Id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(&dtoProduct{
		Id:       product.Id,
		Name:     product.Name,
		Category: product.Category.Name,
		ImageURL: product.ImageURL,
	})
}

// Create
//
//	@Summary	Create new product
//	@Tags		Product
//	@Param		input	body		dtoCreateProduct	true	"Product"
//	@Success	200		{object}	uuid.UUID
//	@Failure	400		{object}	fiber.Error
//	@Failure	401		{object}	customerrors.UnauthorizedError
//	@Failure	500		{object}	fiber.Error
//	@Router		/product [post]
func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	dto := new(dtoCreateProduct)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.productService.Create(ctx.Context(), createProductParams(dto))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(id)
}

// GetAmountTypes
//
//	@Summary	Get amount types for products
//	@Tags		Product
//	@Success	200	{array}		string
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/amount_types [get]
func (h *ProductHandler) GetAmountTypes(ctx *fiber.Ctx) error {
	amountTypes, err := h.productRepository.GetAmountTypes()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(amountTypes)
}

type dtoProduct struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	ImageURL   string    `json:"image_url"`
	AmountType string    `json:"amount_type"`
}

type dtoGetProductById struct {
	Id string `params:"id" validate:"required,uuid4"`
}

type dtoCreateProduct struct {
	Name       string `json:"name" validate:"required"`
	CategoryId string `json:"category_id" validate:"required,uuid4"`
	ImageURL   string `json:"image_url"`
	AmountType string `json:"amount_type"`
}
