package model

import (
	"github.com/google/uuid"
)

type Product struct {
	Id       uuid.UUID `bun:",pk"`
	Name     string
	ImageURL string

	CategoryId uuid.UUID
	Category   *Category `bun:"rel:belongs-to"`
}

func NewProduct(
	id uuid.UUID,
	categoryId uuid.UUID,
	name string,
	imageURL string,
) *Product {
	return &Product{
		Id:         id,
		CategoryId: categoryId,
		Name:       name,
		ImageURL:   imageURL,
	}
}
