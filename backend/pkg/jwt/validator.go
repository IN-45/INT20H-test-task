package jwt

import (
	customerrors "github.com/IN-45/INT20H-test-task/pkg/customerrors"
	"github.com/golang-jwt/jwt"
)

type TokenValidator struct {
	secretKey string
}

func NewTokenValidator(cfg Config) *TokenValidator {
	return &TokenValidator{
		secretKey: cfg.SecretKey,
	}
}

func (v *TokenValidator) ValidateToken(token string) error {
	parsed, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(v.secretKey), nil
	})
	if err != nil {
		return customerrors.NewUnauthorizedError()
	}

	if !parsed.Valid {
		return customerrors.NewUnauthorizedError()
	}

	return nil
}
