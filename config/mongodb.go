package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBExample() {

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Get a handle to the recipes database
	recipesDb := client.Database("recipes")

	// Find a document
	var result bson.M
	err = recipesDb.Collection("recipes").FindOne(ctx, bson.D{{"name", "Test"}}).Decode(&result)

	// Print the result
	fmt.Println(result)

}

// func InsertDocument(){
// 	insertResult, err := recipesDb.Collection("recipes").InsertOne(ctx, bson.D{
// 		{"name", "Pizza"},
// 		{"ingredients", bson.A{
// 			"flour",
// 			"water",
// 			"yeast",
// 		}},
// 	})

// }
