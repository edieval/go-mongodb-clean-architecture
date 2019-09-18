package controllers

import (
	"encoding/json"
	routing "github.com/jackwhelpton/fasthttp-routing"
	"log"
	"main.go/interfaces"
)

type CategoriesController struct {
	interfaces.ICategoriesService
	interfaces.IProductsService
}

func (controller *CategoriesController) GetCategoryWithProducts(c *routing.Context) error {

	categoryCode := c.Param("categoryCode")
	contextStore := c.Query("context")

	category, error := controller.GetCategoryService(categoryCode)
	if error != nil {
		log.Fatal(error)
	}

	aggregateResult := controller.Aggregate(category, contextStore)

	err := json.NewEncoder(c.Response.BodyWriter()).Encode(aggregateResult)
	return err
}
