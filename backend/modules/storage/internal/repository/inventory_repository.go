package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/google/uuid"

	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type InventoryRepository struct {
	db *bun.DB
}

func NewInventoryRepository(db *bun.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) GetAllInventoryProducts(ctx context.Context, userId uuid.UUID) ([]*storage_model.Inventory, error) {
	var inventory []*storage_model.Inventory

	err := r.db.NewSelect().
		Model(&inventory).
		Where("inventory.user_id = ?", userId).
		Scan(ctx)

	return inventory, err
}

func (r *InventoryRepository) AddItem(ctx context.Context, inventory *storage_model.Inventory) error {
	_, err := r.db.NewInsert().Model(inventory).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		// TODO change error
		return customerrors.NewAlreadyExistsError("product already exists")
	}

	return err
}
