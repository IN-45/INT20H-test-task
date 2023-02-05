package service

import (
	"context"
	"time"

	user_model "github.com/IN-45/INT20H-test-task/modules/user/internal/model"
	user_repository "github.com/IN-45/INT20H-test-task/modules/user/internal/repository"
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/IN-45/INT20H-test-task/pkg/hash"
	"github.com/IN-45/INT20H-test-task/pkg/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthService struct {
	userRepository *user_repository.UserRepository
	tokenGenerator *jwt.TokenGenerator
}

func NewAuthService(
	userRepository *user_repository.UserRepository,
	tokenGenerator *jwt.TokenGenerator,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		tokenGenerator: tokenGenerator,
	}
}

func (s *AuthService) SignUp(ctx context.Context, email string, password string) error {
	passwordHash, err := hash.HashString(password)
	if err != nil {
		return errors.Wrap(err, "password hashing error")
	}

	user := user_model.NewUser(uuid.New(), email, passwordHash)

	if err := s.userRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", customerrors.NewNotFoundError("user not found")
	}

	if !hash.IsEqualWithHash(password, user.PasswordHash) {
		return "", errors.New("incorrect password")
	}

	return s.tokenGenerator.GenerateNewAccessToken(24 * time.Hour)
}
