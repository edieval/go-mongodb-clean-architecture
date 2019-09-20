package interfaces

import (
	"main.go/models"
)

type IProductsService interface {
	Aggregate(category models.CategoryModel, contextStore string) map[string]interface{}
}
