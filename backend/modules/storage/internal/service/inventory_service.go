package service

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_repository "github.com/IN-45/INT20H-test-task/modules/storage/internal/repository"
	"github.com/google/uuid"
)

type InventoryParams struct {
	UserId     uuid.UUID
	ProductId  uuid.UUID
	Amount     int
	AmountType string
}

type InventoryService struct {
	inventoryRepository *storage_repository.InventoryRepository
}

func NewInventoryService(
	inventoryRepository *storage_repository.InventoryRepository,
) *InventoryService {
	return &InventoryService{
		inventoryRepository: inventoryRepository,
	}
}

func (s *InventoryService) Create(ctx context.Context, params InventoryParams) error {
	inventory := storage_model.NewInventory(
		params.UserId,
		params.ProductId,
		params.Amount,
		params.AmountType,
	)

	if err := s.inventoryRepository.Create(ctx, inventory); err != nil {
		return err
	}

	return nil
}
