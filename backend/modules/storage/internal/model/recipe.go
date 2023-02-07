package model

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	Id                 uuid.UUID `bun:",pk"`
	Name               string
	Description        string
	AuthorId           uuid.UUID
	CookingTimeMinutes int
	ImageURL           string
	CreatedAt          time.Time
}

func NewRecipe(
	id uuid.UUID,
	name string,
	description string,
	authorId uuid.UUID,
	cookingTimeMinutes int,
	imageURL string,
) *Recipe {
	return &Recipe{
		Id:                 id,
		Name:               name,
		Description:        description,
		AuthorId:           authorId,
		CookingTimeMinutes: cookingTimeMinutes,
		ImageURL:           imageURL,
		CreatedAt:          time.Now(),
	}
}

type RecipeProducts struct {
	RecipeId uuid.UUID
	Recipe   *Recipe `bun:"rel:belongs-to"`

	ProductId uuid.UUID
	Product   *Product `bun:"rel:belongs-to"`

	Amount     int
	AmountType string
}

func NewRecipeProducts(
	recipeId uuid.UUID,
	productId uuid.UUID,
	amount int,
	amountType string,
) *RecipeProducts {
	return &RecipeProducts{
		RecipeId:   recipeId,
		ProductId:  productId,
		Amount:     amount,
		AmountType: amountType,
	}
}
