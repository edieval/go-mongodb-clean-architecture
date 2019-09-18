package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"main.go/models"
)

type IProductsService interface {
	Aggregate(category models.CategoryModel, contextStore string) bson.D
}
