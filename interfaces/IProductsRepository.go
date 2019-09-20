package interfaces

import (
	"main.go/models"
)

type IProductsRepository interface {
	GetProductsRepository(models.CategoryModel, string) map[string]interface{}
}
