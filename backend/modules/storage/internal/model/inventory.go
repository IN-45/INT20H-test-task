package model

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	UserId     uuid.UUID
	AddedAt    time.Time
	Amount     int
	AmountType string

	ProductId uuid.UUID
	Product   *Product `bun:"rel:belongs-to"`
}

func NewInventory(
	userId uuid.UUID,
	productId uuid.UUID,
	amount int,
	amountType string,
) *Inventory {
	return &Inventory{
		UserId:     userId,
		AddedAt:    time.Now(),
		Amount:     amount,
		AmountType: amountType,
		ProductId:  productId,
	}
}
