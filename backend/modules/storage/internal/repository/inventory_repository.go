package repository

import (
	"context"
	"database/sql"

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

func (r *InventoryRepository) GetProductById(ctx context.Context, userId uuid.UUID, productId uuid.UUID) (*storage_model.Inventory, error) {
	inventory := new(storage_model.Inventory)

	if err := r.db.NewSelect().
		Model(inventory).
		Where("inventory.user_id = ?", userId).
		Where("inventory.product_id = ?", productId).
		Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return inventory, nil
}

func (r *InventoryRepository) AddItem(ctx context.Context, inventory *storage_model.Inventory) error {
	_, err := r.db.NewInsert().
		Model(inventory).
		On("CONFLICT (user_id, product_id) DO UPDATE").
		Set("amount=?", inventory.Amount).
		Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("product in user inventory already exists")
	}

	return err
}
