package handler

import (
	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	"github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
	"github.com/google/uuid"
)

func mapDtoProducts(products []*storage_model.Product) []*dtoProduct {
	var dtoProducts []*dtoProduct

	for _, product := range products {
		dtoProducts = append(dtoProducts, &dtoProduct{
			Id:         product.Id,
			Name:       product.Name,
			Category:   product.Category.Name,
			ImageURL:   product.ImageURL,
			AmountType: product.AmountType,
		})
	}

	return dtoProducts
}

func createProductParams(dto *dtoCreateProduct) service.ProductParams {
	return service.ProductParams{
		Name:       dto.Name,
		CategoryId: uuid.MustParse(dto.CategoryId),
		ImageURL:   dto.ImageURL,
		AmountType: dto.AmountType,
	}
}

func createInventoryParams(dto *dtoCreateInventory) service.InventoryParams {
	return service.InventoryParams{
		UserId:     uuid.MustParse(dto.UserId),
		ProductId:  uuid.MustParse(dto.ProductId),
		Amount:     dto.Amount,
		AmountType: dto.AmountType,
	}
}

func createRecipeParams(dto *dtoCreateRecipe) service.RecipeParams {
	instructions := []storage_model.Instruction{}
	for _, v := range dto.Instructions {
		instructions = append(instructions, *storage_model.NewInstruction(
			uuid.Nil,
			v.Description,
			v.Priority,
		))
	}
	recipesProducts := []storage_model.RecipesProducts{}
	for _, v := range dto.Products {
		recipesProducts = append(recipesProducts, *storage_model.NewRecipesProducts(
			uuid.Nil,
			uuid.MustParse(v.ProductId),
			v.Amount,
			v.AmountType,
		))
	}
	return service.RecipeParams{
		Name:               dto.Name,
		Description:        dto.Description,
		AuthorId:           uuid.MustParse(dto.AuthorId),
		CookingTimeMinutes: dto.CookingTimeMinutes,
		ImageURL:           dto.ImageURL,
		Instructions:       instructions,
		RecipesProducts:    recipesProducts,
	}
}

func mapDtoCategories(categories []*storage_model.Category) []*dtoCategory {
	var dtoCategories []*dtoCategory

	for _, category := range categories {
		dtoCategories = append(dtoCategories, &dtoCategory{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	return dtoCategories
}

func mapDtoInventories(inventories []*storage_model.Inventory) []*dtoInventory {
	var dtoInventories []*dtoInventory

	for _, inventory := range inventories {
		dtoInventories = append(dtoInventories, &dtoInventory{
			ProductId:  inventory.ProductId,
			Amount:     inventory.Amount,
			AmountType: inventory.AmountType,
		})
	}

	return dtoInventories
}
