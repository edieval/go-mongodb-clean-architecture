package services

import (
	"log"
	"main.go/interfaces"
	"main.go/models"
)

type CategoriesService struct {
	interfaces.ICategoriesRepository
}

func (service *CategoriesService) GetCategoryService(code string) (models.CategoryModel, error) {
	category, err := service.GetCategoryRepository(code)
	log.Print("category", category)

	return category, err
}
