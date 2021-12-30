package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Users struct {
	UserEmail string    `bson:"usr_email"`
	Password  string    `bson:"password"`
	LastLogin time.Time `bson:"last_login"`
	FirstName string    `bson:"firstname"`
	LastName  string    `bson:"lastname"`
}

// Connection URI
const uri = "mongodb://127.0.0.1:27017"

func runMongoDbTest() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected and pinged.")

		// Select database and collection
		userCollection := client.Database("blogdb").Collection("users")

		// Select Test
		fmt.Println("Select test")
		// convert id string to ObjectId
		objectId, err := primitive.ObjectIDFromHex("61c42fe7dda6e33854e3c2e1")
		if err != nil {
			fmt.Println("Invalid id")
		} else {
			var result bson.M
			filter := bson.M{"_id": objectId}
			err := userCollection.FindOne(ctx, filter).Decode(&result)
			if err != nil {
				panic(err)
			} else {
				fmt.Println(result)
			}
		}

		// Insert Test
		fmt.Println("Insert test")
		doc := &Users{
			UserEmail: "indra.nureska@gmail.com",
			Password:  "Password1",
			FirstName: "Indra",
			LastName:  "Nureska",
			LastLogin: time.Now(),
		}
		insertResult, err := userCollection.InsertOne(ctx, doc)
		fmt.Printf("Inserted document with _id: %v\n", insertResult.InsertedID)

		if err != nil {
			panic(err)
		} else {
			//Update Test
			fmt.Println("Update test")
			updateResult, err := userCollection.UpdateOne(
				ctx, bson.M{"_id": insertResult.InsertedID},
				bson.D{{"$set", bson.D{{"firstname", "Indra"}}}},
			)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Updated %v Documents!\n", updateResult.ModifiedCount)

			// Delete Test
			fmt.Println("Delete test")
			delResult, err := userCollection.DeleteOne(ctx, bson.M{"_id": insertResult.InsertedID})
			if err != nil {
				panic(err)
			}
			fmt.Printf("DeleteOne removed %v document(s)\n", delResult.DeletedCount)
		}

	}

}
