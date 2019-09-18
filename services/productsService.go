package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"main.go/interfaces"
	"main.go/models"
)

type ProductsService struct {
	interfaces.IProductsRepository
}

func (service *ProductsService) Aggregate(category models.CategoryModel, contextStore string) bson.D {
	return service.GetProductsRepository(category, contextStore)
}
