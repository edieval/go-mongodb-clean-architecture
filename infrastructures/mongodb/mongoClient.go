package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() *mongo.Client {
	client, _ := mongo.Connect(GetConnectionContext(), options.Client().ApplyURI("mongodb://localhost:27017"))
	return client
}
