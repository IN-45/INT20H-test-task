package service

import (
	"context"
	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/google/uuid"
	"sort"
)

type FilterRecipesService struct {
	recipeService    *RecipeService
	inventoryService *InventoryService
}

func NewFilterRecipesService(
	recipeService *RecipeService,
	inventoryService *InventoryService,
) *FilterRecipesService {
	return &FilterRecipesService{
		recipeService:    recipeService,
		inventoryService: inventoryService,
	}
}

func (f *FilterRecipesService) FilterRecipes(ctx context.Context, userId uuid.UUID) ([]*FilterRecipe, error) {
	recipes, err := f.recipeService.GetAllRecipes(ctx)
	if err != nil {
		return nil, err
	}

	inventoryProducts, err := f.inventoryService.inventoryRepository.GetExistingProducts(ctx, userId)
	if err != nil {
		return nil, err
	}

	var filterRecipes []*FilterRecipe

	for _, recipe := range recipes {
		var filterProducts []*FilterProduct
		for _, recipeProduct := range recipe.Products {
			missedAmount := 0
			inventoryProduct := f.findInInventory(recipeProduct, inventoryProducts)
			if inventoryProduct == nil {
				missedAmount = recipeProduct.Amount
			} else if inventoryProduct.Amount < recipeProduct.Amount {
				missedAmount = recipeProduct.Amount - inventoryProduct.Amount
			}
			filterProducts = append(filterProducts, &FilterProduct{
				ProductId:    recipeProduct.ProductId,
				Name:         recipeProduct.Product.Name,
				Amount:       recipeProduct.Amount,
				MissedAmount: missedAmount,
				AmountType:   recipeProduct.AmountType,
				ImageURL:     recipeProduct.Product.ImageURL,
				CategoryId:   recipeProduct.Product.CategoryId,
			})
		}

		filterRecipes = append(filterRecipes, &FilterRecipe{
			Recipe:       recipe.Recipe,
			Instructions: recipe.Instructions,
			Products:     filterProducts,
		})
	}

	return f.sort(filterRecipes), nil
}

func (f *FilterRecipesService) findInInventory(
	recipeProduct *storage_model.RecipeProducts,
	inventoryProducts []*storage_model.Inventory) *storage_model.Inventory {
	for _, product := range inventoryProducts {
		if product.ProductId == recipeProduct.ProductId && product.AmountType == recipeProduct.AmountType {
			return product
		}
	}
	return nil
}

func (f *FilterRecipesService) sort(recipes []*FilterRecipe) []*FilterRecipe {
	sort.Slice(recipes, func(i, j int) bool {
		sumI := 0
		sumJ := 0
		for _, product := range recipes[i].Products {
			if product.MissedAmount > 0 {
				sumI++
			}
		}
		for _, product := range recipes[j].Products {
			if product.MissedAmount > 0 {
				sumJ++
			}
		}

		return sumI < sumJ
	})

	return recipes
}

type FilterProduct struct {
	ProductId    uuid.UUID `json:"product_id"`
	Name         string    `json:"name"`
	Amount       int       `json:"amount"`
	MissedAmount int       `json:"missed_amount"`
	AmountType   string    `json:"amount_type"`
	ImageURL     string    `json:"image_url"`
	CategoryId   uuid.UUID `json:"category_id"`
}

type FilterRecipe struct {
	Recipe       *storage_model.Recipe
	Instructions []*storage_model.Instruction
	Products     []*FilterProduct
}
