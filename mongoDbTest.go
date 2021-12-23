package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://127.0.0.1:27017"

func runMongoDbTest() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected and pinged.")

		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Select Test
		println("Select test")


		// Insert Test
		println("Insert test")
		// doc := bson.D{{"usr_email", "indra.nureska@gmail.com"}, {"password", "Password1"}, {"last_login", ""}, {"firstname", ""}, {"lastname", ""}}

		// result, err := userCollection.InsertOne(context.TODO(), doc)
		// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

		// Update Test
		println("Update test")

		// Delete Test
		println("Delete test")
	}

}
