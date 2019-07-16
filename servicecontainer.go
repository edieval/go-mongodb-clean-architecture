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

	categoriesRepository := &repositories.CategoryRepository{mongodbClient}

	categoriesService := &services.CategoriesService{categoriesRepository}

	playerController := controllers.CategoriesController{categoriesService}

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
