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
			Id:       product.Id,
			Name:     product.Name,
			Category: product.Category.Name,
			ImageURL: product.ImageURL,
		})
	}

	return dtoProducts
}

func createDtoToProductParams(dto *dtoCreateProduct) service.ProductParams {
	return service.ProductParams{
		Name:       dto.Name,
		CategoryId: uuid.MustParse(dto.CategoryId),
		ImageURL:   dto.ImageURL,
	}
}

func createDtoToInventoryParams(dto *dtoCreateInventory) service.InventoryParams {
	return service.InventoryParams{
		UserId:     dto.UserId,
		ProductId:  uuid.MustParse(dto.ProductId),
		Amount:     dto.Amount,
		AmountType: dto.AmountType,
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
			UserId:     inventory.UserId,
			ProductId:  inventory.ProductId,
			Amount:     inventory.Amount,
			AmountType: inventory.AmountType,
		})
	}

	return dtoInventories
}
