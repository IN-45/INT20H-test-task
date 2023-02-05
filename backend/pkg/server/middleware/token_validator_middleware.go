package middleware

import (
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/IN-45/INT20H-test-task/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type tokenValidatorMiddleware struct {
	tokenValidator *jwt.TokenValidator
}

func NewTokenValidatorMiddleware(tokenValidator *jwt.TokenValidator) *tokenValidatorMiddleware {
	return &tokenValidatorMiddleware{tokenValidator: tokenValidator}
}

func (m *tokenValidatorMiddleware) validateToken(ctx *fiber.Ctx) error {
	if err := m.tokenValidator.ValidateToken(ctx.Cookies("token")); err != nil {
		if _, ok := err.(*customerrors.UnauthorizedError); ok {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Next()
}
