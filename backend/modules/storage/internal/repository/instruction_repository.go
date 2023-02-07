package repository

import (
	"context"

	"github.com/IN-45/INT20H-test-task/pkg/db"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/google/uuid"
)

type InstructionRepository struct {
	db *db.TransactionRepository
}

func NewInstructionRepository(db *db.TransactionRepository) *InstructionRepository {
	return &InstructionRepository{db: db}
}

func (r *InstructionRepository) Create(ctx context.Context, instructions []storage_model.Instruction) error {
	_, err := r.db.NewInsert(ctx).Model(&instructions).Exec(ctx)

	return err
}

func (r *InstructionRepository) GetByRecipeId(ctx context.Context, id uuid.UUID) ([]*storage_model.Instruction, error) {
	var instructions []*storage_model.Instruction

	err := r.db.NewSelect(ctx).
		Model(&instructions).
		Where("recipe_id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
