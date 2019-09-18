package services

import (
	"main.go/interfaces"
	"main.go/models"
)

type CategoriesService struct {
	interfaces.ICategoryRepository
}

func (service *CategoriesService) GetCategoryService(code string) (models.CategoryModel, error) {
	return service.GetCategoryRepository(code)
}
