package handler

import (
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	_ "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InventoryHandler struct {
	validator           *validator.Validate
	inventoryRepository *storage_repository.InventoryRepository
	inventoryService    *storage_service.InventoryService
}

func NewInventoryHandler(
	validator *validator.Validate,
	inventoryRepository *storage_repository.InventoryRepository,
	inventoryService *storage_service.InventoryService,
) *InventoryHandler {
	return &InventoryHandler{
		validator:           validator,
		inventoryRepository: inventoryRepository,
		inventoryService:    inventoryService,
	}
}

func RegisterInventoryHandler(app *fiber.App, h *InventoryHandler) {
	app.Post("/inventory", h.Create)
	app.Get("/inventory", h.GetAllInventoryProducts)
}

// GetAllInventoryProducts
//
//	@Summary	Get all inventory products
//	@Tags		Inventory
//	@Success	200	{array}		dtoInventory
//	@Failure	401	{object}	customerrors.UnauthorizedError
//	@Failure	500	{object}	fiber.Error
//	@Router		/inventory [get]
func (h *InventoryHandler) GetAllInventoryProducts(ctx *fiber.Ctx) error {
	user_id := ctx.Cookies("user_id")
	if user_id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user id not found")
	}

	inventories, err := h.inventoryRepository.GetAllInventoryProducts(ctx.Context(), uuid.MustParse(user_id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(mapDtoInventories(inventories))
}

// AddProductToInventory
//
//	@Summary	Add product to inventory
//	@Tags		Inventory
//	@Param		input	body		dtoCreateInventory	true	"Inventory"
//	@Success	200		{object}	uuid.UUID
//	@Failure	400		{object}	fiber.Error
//	@Failure	401		{object}	customerrors.UnauthorizedError
//	@Failure	500		{object}	fiber.Error
//	@Router		/inventory [post]
func (h *InventoryHandler) Create(ctx *fiber.Ctx) error {
	dto := new(dtoCreateInventory)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	user_id := ctx.Cookies("user_id")
	if user_id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user id not found")
	}

	dto.UserId = user_id

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := h.inventoryService.AddItem(ctx.Context(), createInventoryParams(dto))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON("created")
}

type dtoCreateInventory struct {
	UserId     string `json:"-" validate:"uuid4"`
	ProductId  string `json:"product_id" validate:"required,uuid4"`
	Amount     int    `json:"amount"`
	AmountType string `json:"amount_type"`
}

type dtoInventory struct {
	ProductId  uuid.UUID `json:"product_id"`
	Amount     int       `json:"amount"`
	AmountType string    `json:"amount_type"`
}
