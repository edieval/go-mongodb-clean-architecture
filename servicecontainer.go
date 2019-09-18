package main

import (
	"main.go/controllers"
	"main.go/infrastructures/mongodb"
	"main.go/repositories"
	"main.go/services"
	"sync"
)

type IServiceContainer interface {
	InjectCategoriesController() controllers.CategoriesController
}

type kernel struct{}

func (k *kernel) InjectCategoriesController() controllers.CategoriesController {

	mongodbClient := mongodb.GetMongoClient()

	categoryRepository := &repositories.CategoryRepository{mongodbClient}
	productsRepository := &repositories.ProductsRepository{mongodbClient}

	categoriesService := &services.CategoriesService{categoryRepository}
	productsService := &services.ProductsService{productsRepository}

	playerController := controllers.CategoriesController{categoriesService, productsService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
