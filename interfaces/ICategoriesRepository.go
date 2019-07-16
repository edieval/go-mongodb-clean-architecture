package interfaces

import "main.go/models"

type ICategoriesRepository interface {
	GetCategoryRepository(code string) (models.CategoryModel, error)
}
