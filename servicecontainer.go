package main

import (
	"context"
	"log"
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

	err := mongodbClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!")

	categoriesCollection := mongodbClient.Database("opus-category").Collection("categories")

	categoriesRepository := &repositories.CategoryRepository{categoriesCollection}
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
