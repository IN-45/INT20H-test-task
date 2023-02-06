package service

import (
	"context"
	"errors"

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

func (s *InventoryService) AddItem(ctx context.Context, params InventoryParams) error {
	inventory := storage_model.NewInventory(
		params.UserId,
		params.ProductId,
		params.Amount,
		params.AmountType,
	)

	storedInventory, err := s.inventoryRepository.GetProductById(ctx, params.UserId, params.ProductId)
	if err != nil {
		return err
	}

	if storedInventory != nil {
		if storedInventory.AmountType == inventory.AmountType {
			inventory.Amount += storedInventory.Amount
		} else {
			return errors.New("invalid amount type")
		}
	}

	if err := s.inventoryRepository.AddItem(ctx, inventory); err != nil {
		return err
	}

	return nil
}
