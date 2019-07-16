package services

import (
	"main.go/interfaces"
	"main.go/models"
)

type CategoriesService struct {
	interfaces.ICategoriesRepository
}

func (service *CategoriesService) GetCategoryService(code string) (models.CategoryModel, error) {
	category, err := service.GetCategoryRepository(code)
	return category, err
}
