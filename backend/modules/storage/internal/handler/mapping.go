package handler

import (
	storage_model "github.com/IN-45/INT20H-test-task/modules/storage/internal/model"
	storage_service "github.com/IN-45/INT20H-test-task/modules/storage/internal/service"
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

func createProductParams(dto *dtoCreateProduct) storage_service.ProductParams {
	return storage_service.ProductParams{
		Name:       dto.Name,
		CategoryId: uuid.MustParse(dto.CategoryId),
		ImageURL:   dto.ImageURL,
		AmountType: dto.AmountType,
	}
}

func createInventoryParams(dto *dtoCreateInventory) storage_service.InventoryParams {
	return storage_service.InventoryParams{
		UserId:     uuid.MustParse(dto.UserId),
		ProductId:  uuid.MustParse(dto.ProductId),
		Amount:     dto.Amount,
		AmountType: dto.AmountType,
	}
}

func createRecipeParams(dto *dtoCreateRecipe) storage_service.RecipeParams {
	var instructions []storage_service.Instructions
	for _, v := range dto.Instructions {
		instructions = append(instructions, storage_service.Instructions{
			RecipeId:    uuid.Nil,
			Description: v.Description,
			Priority:    v.Priority,
		})
	}

	var RecipeProducts []storage_service.RecipeProducts
	for _, v := range dto.Products {
		RecipeProducts = append(RecipeProducts, storage_service.RecipeProducts{
			RecipeId:   uuid.Nil,
			ProductId:  uuid.MustParse(v.ProductId),
			Amount:     v.Amount,
			AmountType: v.AmountType,
		})
	}

	return storage_service.RecipeParams{
		Name:               dto.Name,
		Description:        dto.Description,
		AuthorId:           uuid.MustParse(dto.AuthorId),
		CookingTimeMinutes: dto.CookingTimeMinutes,
		ImageURL:           dto.ImageURL,
		Instructions:       instructions,
		RecipeProducts:     RecipeProducts,
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

func mapServiceRecipesToDto(recipes []*storage_service.Recipe) []*dtoRecipe {
	var dtoRecipes []*dtoRecipe

	for _, recipe := range recipes {
		dtoRecipes = append(dtoRecipes, mapServiceRecipeToDto(recipe))
	}

	return dtoRecipes
}

func mapServiceRecipeToDto(recipe *storage_service.Recipe) *dtoRecipe {
	return &dtoRecipe{
		Id:                 recipe.Recipe.Id,
		Name:               recipe.Recipe.Name,
		Description:        recipe.Recipe.Description,
		AuthorId:           recipe.Recipe.AuthorId,
		CookingTimeMinutes: recipe.Recipe.CookingTimeMinutes,
		ImageURL:           recipe.Recipe.ImageURL,
		Instructions:       mapDtoInstructions(recipe.Instructions),
		Products:           mapDtoRecipeProducts(recipe.Products),
	}
}

func mapDtoInstructions(instructions []*storage_model.Instruction) []*dtoInstruction {
	var dtoInstructions []*dtoInstruction

	for _, instruction := range instructions {
		dtoInstructions = append(dtoInstructions, &dtoInstruction{
			Description: instruction.Description,
			Priority:    instruction.Priority,
		})
	}

	return dtoInstructions
}

func mapDtoRecipeProducts(products []*storage_model.RecipeProducts) []*dtoRecipeProduct {
	var dtoRecipeProducts []*dtoRecipeProduct

	for _, product := range products {
		dtoRecipeProducts = append(dtoRecipeProducts, &dtoRecipeProduct{
			ProductId:  product.ProductId,
			Name:       product.Product.Name,
			Amount:     product.Amount,
			AmountType: product.AmountType,
			ImageURL:   product.Product.ImageURL,
			CategoryId: product.Product.CategoryId,
		})
	}

	return dtoRecipeProducts
}

func mapServiceFilterRecipesToDto(recipes []*storage_service.FilterRecipe) []*dtoFilterRecipe {
	var dtoFilterRecipes []*dtoFilterRecipe

	for _, recipe := range recipes {
		dtoFilterRecipes = append(dtoFilterRecipes, mapServiceFilterRecipeToDto(recipe))
	}

	return dtoFilterRecipes
}

func mapServiceFilterRecipeToDto(recipe *storage_service.FilterRecipe) *dtoFilterRecipe {
	return &dtoFilterRecipe{
		Id:                 recipe.Recipe.Id,
		Name:               recipe.Recipe.Name,
		Description:        recipe.Recipe.Description,
		AuthorId:           recipe.Recipe.AuthorId,
		CookingTimeMinutes: recipe.Recipe.CookingTimeMinutes,
		ImageURL:           recipe.Recipe.ImageURL,
		Instructions:       mapDtoInstructions(recipe.Instructions),
		Products:           mapDtoFilterProducts(recipe.Products),
	}
}

func mapDtoFilterProducts(products []*storage_service.FilterProduct) []*dtoFilterProduct {
	var dtoFilterProducts []*dtoFilterProduct

	for _, product := range products {
		dtoFilterProducts = append(dtoFilterProducts, &dtoFilterProduct{
			ProductId:    product.ProductId,
			Name:         product.Name,
			Amount:       product.Amount,
			MissedAmount: product.MissedAmount,
			AmountType:   product.AmountType,
			ImageURL:     product.ImageURL,
			CategoryId:   product.CategoryId,
		})
	}

	return dtoFilterProducts
}
