package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"main.go/models"
)

type IProductsRepository interface {
	GetProductsRepository(models.CategoryModel, string) bson.D
}
