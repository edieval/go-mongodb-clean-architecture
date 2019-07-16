package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/infrastructures/mongodb"
	"main.go/models"
)

type CategoryRepository struct {
	*mongo.Client
}

func (client CategoryRepository) GetCategoryRepository(code string) (models.CategoryModel, error) {
	categoriesCollection := client.Database("opus-category").Collection("categories")

	filter := bson.M{"code": code}

	var category models.CategoryModel
	err := categoriesCollection.FindOne(mongodb.GetQueryContext(), filter).Decode(&category)

	return category, err
}

//
//func getCategoryProducts(collection *mongo.Collection, category models.CategoryModel) {
//	var pipeline interface{}
//	error := bson.UnmarshalExtJSON([]byte(`[{"$match":{"publication.published":true}},{"$match":{"$or":[{"model":{"$in":["200711"]}}]}},{"$addFields":{"price":"$prices.380.price","allProducts":"allProducts"}},{"$facet":{"_total":[{"$count":"total"}],"_products":[{"$project":{"_id":1,"type":"product","sortedField":{"$ifNull":["$scores.score4",0]}}},{"$sort":{"sortedField":-1,"_id":1}},{"$skip":0},{"$limit":30}],"prices":[{"$match":{"price":{"$exists":true}}},{"$group":{"_id":"$allProducts","min":{"$min":"$price"},"max":{"$max":"$price"}}}],"products-type":[{"$match":{"attributes.08547":{"$exists":true}}},{"$group":{"_id":"$attributes.08547","count":{"$sum":1}}},{"$sort":{"count":1}}]}}]`), true, &pipeline)
//	if error != nil {
//		panic(error)
//	}
//
//	_ = []bson.M{
//		bson.M{
//			"$match": bson.M{"publication.published": true},
//		},
//		bson.M{
//			"$match": bson.M{"$or": true},
//		},
//		bson.M{
//			"$project": bson.M{"_id": 1},
//		},
//	}
//
//	cursor, _ := collection.Aggregate(mongodb.GetQueryContext(), pipeline)
//
//	for cursor.Next(mongodb.GetQueryContext()) {
//		elem := &bson.D{}
//		if err := cursor.Decode(elem); err != nil {
//			log.Fatal(err)
//		}
//
//		log.Print(elem)
//	}
//}
