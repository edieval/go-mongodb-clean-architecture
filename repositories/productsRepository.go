package repositories

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main.go/infrastructures/mongodb"
	"main.go/models"
)

type ProductsRepository struct {
	*mongo.Client
}

func (client ProductsRepository) GetProductsRepository(category models.CategoryModel, contextStore string) map[string]interface{} {
	productsCollection := client.Database("opus-category").Collection("products")

	var query []interface{}
	query = append(
		query,
		buildPopulationMatch(category.PopulationRules),
		getPublishedMatch(),
		addContextualized(contextStore),
		buildFacet(category.Filters),
	)

	cursor, _ := productsCollection.Aggregate(mongodb.GetQueryContext(), query)

	var result map[string]interface{} //bson.D

	cursor.Next(mongodb.GetQueryContext())
	err := cursor.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	var temporaryBytes, _ = bson.MarshalExtJSON(result, true, true)

	var jsonDocument map[string]interface{}

	err = json.Unmarshal(temporaryBytes, &jsonDocument)

	return jsonDocument
}

func addContextualized(context string) interface{} {
	if context == "" {
		context = "380"
	}
	addField := map[string]interface{}{
		"$addFields": map[string]interface{}{
			"contextualized": "$contextualized." + context,
			"allProducts":    "allProducts",
		},
	}

	return addField
}

func buildFacet(filters []models.Filter) interface{} {
	facets := map[string]interface{}{
		"_total":    getTotalFacet(),
		"_products": buildProductsFacet(),
		"prices":    buildPriceFacet(),
	}

	for _, filter := range filters {
		if filter.Type == "standard" || filter.Type == "toggle" {
			var filterFacet = map[string]interface{}{
				filter.Code: buildFilterFacet(filter),
			}
			facets = concatJson(facets, filterFacet)
		}
	}

	return map[string]interface{}{
		"$facet": facets,
	}
}

func buildPriceFacet() []interface{} {
	var priceFacet []interface{}

	match := map[string]interface{}{
		"$match": map[string]interface{}{
			"contextualized.price.price": map[string]interface{}{
				"$exists": true,
			},
		},
	}

	group := map[string]interface{}{
		"$group": map[string]interface{}{
			"_id": "$allProducts",
			"min": map[string]interface{}{
				"$min": "$contextualized.price.price",
			},
			"max": map[string]interface{}{
				"$max": "$contextualized.price.price",
			},
		},
	}
	priceFacet = append(priceFacet, match, group)

	return priceFacet
}

func buildFilterFacet(filter models.Filter) []interface{} {
	var filterFacet []interface{}

	match := map[string]interface{}{
		"$match": map[string]interface{}{
			filter.Path: map[string]interface{}{
				"$exists": true,
			},
		},
	}

	group := map[string]interface{}{
		"$group": map[string]interface{}{
			"_id": "$" + filter.Path,
			"count": map[string]interface{}{
				"$sum": 1,
			},
		},
	}

	sortDirection := 1
	if filter.StandardFilterFields.Sort == "count_desc" {
		sortDirection = -1
	}

	sort := map[string]interface{}{
		"$sort": map[string]interface{}{
			"count": sortDirection,
		},
	}

	filterFacet = append(filterFacet, match, group, sort)
	return filterFacet
}

func buildProductsFacet() []interface{} {
	var productsFacet []interface{}

	project := map[string]interface{}{
		"$project": map[string]interface{}{
			"_id":  1,
			"type": "product",
			"sortedField": map[string]interface{}{
				"$ifNull": []interface{}{"$scores.score4", 0},
			},
		},
	}

	sort := map[string]interface{}{
		"$sort": map[string]interface{}{
			"_id":         1,
			"sortedField": -1,
		},
	}

	skip := map[string]interface{}{
		"$skip": 0,
	}

	limit := map[string]interface{}{
		"$limit": 30,
	}

	productsFacet = append(productsFacet, project, sort, skip, limit)

	return productsFacet
}

func getTotalFacet() []interface{} {
	var totalFacet []interface{}

	totalFacet = append(totalFacet, map[string]interface{}{
		"$count": "total",
	})
	return totalFacet
}

func getPublishedMatch() interface{} {
	match := map[string]interface{}{
		"$match": map[string]interface{}{
			"publication.published": true,
		},
	}

	return match
}

func getProject() interface{} {
	return map[string]interface{}{
		"$project": map[string]interface{}{
			"_id":   1,
			"price": 1,
		},
	}
}

func buildPopulationMatch(popRules []models.PopulationRules) interface{} {
	var acc []interface{}
	for _, populationRule := range popRules {
		var modelRule = map[string]interface{}{
			"model": map[string]interface{}{
				"$in": populationRule.ModelCodes,
			},
		}
		acc = append(acc, concatJson(modelRule, buildAttributesRules(populationRule.AttributeRules)))
	}

	match := map[string]interface{}{
		"$match": map[string]interface{}{
			"$or": acc,
		},
	}

	return match
}

func buildAttributesRules(attributeRules []models.AttributeRule) map[string]interface{} {
	var result map[string]interface{}
	for _, attributeRule := range attributeRules {
		var attrib = map[string]interface{}{
			"attributes." + attributeRule.AttributeCode: map[string]interface{}{
				convertOperator(attributeRule.Operator): attributeRule.AttributeValues,
			},
		}

		result = concatJson(result, attrib)
	}
	return result
}

func convertOperator(operator string) string {
	if operator == "=" {
		return "$in"
	} else {
		return "$nin"
	}
}

func concatJson(json1 map[string]interface{}, json2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range json1 {
		result[k] = v
	}
	for k, v := range json2 {
		result[k] = v
	}
	return result
}
