package model

import (
	"github.com/google/uuid"
)

type Category struct {
	Id   uuid.UUID `bun:",pk"`
	Name string
}

func NewCategory(
	id uuid.UUID,
	name string,
) *Category {
	return &Category{
		Id:   id,
		Name: name,
	}
}
