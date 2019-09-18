package interfaces

import "main.go/models"

type ICategoryRepository interface {
	GetCategoryRepository(code string) (models.CategoryModel, error)
}
