package interfaces

import "main.go/models"

type ICategoriesService interface {
	GetCategoryService(code string) (models.CategoryModel, error)
}
