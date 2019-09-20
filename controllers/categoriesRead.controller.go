package controllers

import (
	"encoding/json"
	routing "github.com/jackwhelpton/fasthttp-routing"
	"log"
	"main.go/interfaces"
	"main.go/models"
	"strconv"
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

	categoryInfo := CategoryInfo{
		Type:        category.CategoryType,
		Code:        category.Code,
		Label:       category.Label,
		Description: category.Description,
		Publication: Publication{
			Status: category.Publication.Status,
		},
	}

	//TOTAL ##############################################
	var test, _ = json.Marshal(aggregateResult["_total"].([]interface{})[0])

	var obj2 models.Total

	json.Unmarshal(test, &obj2)

	total, _ := strconv.ParseInt(obj2.Total.Number, 10, 0)

	//PRODUCTS  ##############################################
	var test2, _ = json.Marshal(aggregateResult["_products"].([]interface{}))

	var contenus []models.Product

	json.Unmarshal(test2, &contenus)
	var qfdqdf []Content

	for _, c := range contenus {
		qfdqdf = append(qfdqdf, Content{Code: c.Id, Type: c.Type})
	}

	//FILTERS  ##############################################
	var filters []Filter
	for _, categoryFilter := range category.Filters {
		if categoryFilter.Type == "standard" || categoryFilter.Type == "toggle" {
			var filterResults []models.FacetResult

			var filterValues []FilterValue

			var test3, _ = json.Marshal(aggregateResult[categoryFilter.Code].([]interface{}))
			json.Unmarshal(test3, &filterResults)

			for _, c := range filterResults {
				totalFilterValue, _ := strconv.ParseInt(c.Count.Number, 10, 0)
				filterValues = append(filterValues, FilterValue{Type: "standard", Value: c.Id, Total: totalFilterValue})
			}
			filters = append(filters, Filter{Type: categoryFilter.Type, Code: categoryFilter.Code, Label: categoryFilter.Label, Values: filterValues})
		} else if categoryFilter.Type == "range" || categoryFilter.IsExtended {
			var filterRangeResult models.FacetRangeResult
			var test, _ = json.Marshal(aggregateResult[categoryFilter.Code].([]interface{})[0])
			json.Unmarshal(test, &filterRangeResult)
			min, _ := strconv.ParseFloat(filterRangeResult.Min.Double, 64)
			max, _ := strconv.ParseFloat(filterRangeResult.Max.Double, 64)
			filters = append(filters, Filter{Type: categoryFilter.Type, Code: categoryFilter.Code, Label: categoryFilter.Label, Min: min, Max: max})
		}
	}

	contents := Contents{
		Offset:  0,
		Limit:   30,
		Total:   total,
		Context: contextStore,
		Values:  qfdqdf,
		Filters: filters,
	}

	response := CategoryResponse{categoryInfo, contents}

	err := json.NewEncoder(c.Response.BodyWriter()).Encode(response)
	return err
}
