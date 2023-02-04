package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(
	id uuid.UUID,
	email string,
	passwordHash string,
) *User {
	return &User{
		Id:           id,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}
