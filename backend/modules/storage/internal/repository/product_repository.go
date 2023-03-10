package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/IN-45/INT20H-test-task/pkg/db"
	"github.com/google/uuid"

	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/uptrace/bun/driver/pgdriver"
)

type ProductRepository struct {
	db *db.TransactionRepository
}

func NewProductRepository(db *db.TransactionRepository) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetById(ctx context.Context, id uuid.UUID) (*storage_model.Product, error) {
	product := new(storage_model.Product)

	err := r.db.NewSelect(ctx).
		Model(product).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetAllWithCategories(ctx context.Context) ([]*storage_model.Product, error) {
	var products []*storage_model.Product

	err := r.db.NewSelect(ctx).
		Model(&products).
		Relation("Category").
		Scan(ctx)

	return products, err
}

func (r *ProductRepository) GetByIdWithCategory(ctx context.Context, id uuid.UUID) (*storage_model.Product, error) {
	product := new(storage_model.Product)

	err := r.db.NewSelect(ctx).
		Model(product).
		Where("product.id = ?", id).
		Relation("Category").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *storage_model.Product) error {
	_, err := r.db.NewInsert(ctx).Model(product).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("product already exists")
	}

	return err
}

func (r *ProductRepository) GetAmountTypes(ctx context.Context) ([]string, error) {
	var types []string

	rows, err := r.db.Query(ctx, "select enumlabel from pg_enum")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var scanType string
		err = rows.Scan(&scanType)
		if err != nil {
			return nil, err
		}
		types = append(types, scanType)
	}
	return types, err
}
