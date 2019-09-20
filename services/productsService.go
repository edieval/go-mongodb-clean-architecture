package services

import (
	"main.go/interfaces"
	"main.go/models"
)

type ProductsService struct {
	interfaces.IProductsRepository
}

func (service *ProductsService) Aggregate(category models.CategoryModel, contextStore string) map[string]interface{} {
	return service.GetProductsRepository(category, contextStore)
}
