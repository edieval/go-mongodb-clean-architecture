package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetMongoClient() *mongo.Client {
	client, _ := mongo.Connect(GetConnectionContext(), options.Client().ApplyURI("mongodb://localhost:27017"))
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!")
	return client
}
