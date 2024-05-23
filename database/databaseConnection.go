package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxt, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("shop").Collection(collectionName)
	return collection
}
