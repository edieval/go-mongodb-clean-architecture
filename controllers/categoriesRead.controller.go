package controllers

import (
	"encoding/json"
	routing "github.com/jackwhelpton/fasthttp-routing"
	"log"
	"main.go/interfaces"
)

type CategoriesController struct {
	interfaces.ICategoriesService
}

func (controller *CategoriesController) GetCategoryWithProducts(c *routing.Context) error {

	categoryCode := c.Param("categoryCode")

	category, error := controller.GetCategoryService(categoryCode)
	if error != nil {
		log.Fatal(error)
	}

	err := json.NewEncoder(c.Response.BodyWriter()).Encode(category)
	return err
}
