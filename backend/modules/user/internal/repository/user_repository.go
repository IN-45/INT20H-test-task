package repository

import (
	"context"
	"database/sql"

	usermodel "github.com/IN-45/INT20H-test-task/modules/user/internal/model"
	customerrors "github.com/IN-45/INT20H-test-task/pkg/custom_errors"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	user := new(usermodel.User)

	if err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *usermodel.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)

	if e, ok := err.(pgdriver.Error); ok && e.IntegrityViolation() {
		return customerrors.NewAlreadyExistsError("user already exists")
	}

	return errors.WithStack(err)
}
