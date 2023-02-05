package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenGenerator struct {
	secretKey string
}

func NewTokenGenerator(cfg Config) *TokenGenerator {
	return &TokenGenerator{
		secretKey: cfg.SecretKey,
	}
}

func (g *TokenGenerator) GenerateNewAccessToken(ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(g.secretKey))
}

func (g *TokenGenerator) GenerateExpiredToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now()})

	return token.SignedString([]byte(g.secretKey))
}
