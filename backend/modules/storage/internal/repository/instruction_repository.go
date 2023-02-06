package repository

import (
	"context"

	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type InstructionRepository struct {
	db *bun.DB
}

func NewInstructionRepository(db *bun.DB) *InstructionRepository {
	return &InstructionRepository{db: db}
}

func (r *InstructionRepository) AddInstructions(ctx context.Context, instructions []storage_model.Instruction) error {
	_, err := r.db.NewInsert().Model(&instructions).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("recipe already exists")
	}

	return err
}
