package model

import (
	"github.com/google/uuid"
)

type Instruction struct {
	RecipeId    uuid.UUID
	Description string
	Priority    int
}

func NewInstruction(
	recipeId uuid.UUID,
	description string,
	priority int,
) *Instruction {
	return &Instruction{
		RecipeId:    recipeId,
		Description: description,
		Priority:    priority,
	}
}
