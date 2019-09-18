package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/infrastructures/mongodb"
	"main.go/models"
)

type CategoryRepository struct {
	*mongo.Client
}

func (client CategoryRepository) GetCategoryRepository(code string) (models.CategoryModel, error) {
	categoriesCollection := client.Database("opus-category").Collection("categories")

	filter := bson.M{"code": code}

	var category models.CategoryModel
	err := categoriesCollection.FindOne(mongodb.GetQueryContext(), filter).Decode(&category)

	return category, err
}
