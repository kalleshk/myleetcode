package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection URI
	const mongoURI = "mongodb://localhost:27017/"

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	fmt.Println("Connected to MongoDB!")

	// Select database and collection
	db := client.Database("BookParking")
	collection := db.Collection("Parking")

	// Perform CRUD operations
	// 1. Create
	newDoc := bson.D{
		{Key: "name", Value: "John Doe"},
		{Key: "age", Value: 30},
		{Key: "city", Value: "New York"},
	}
	insertResult, err := collection.InsertOne(ctx, newDoc)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	// 2. Read
	filter := bson.D{{Key: "name", Value: "John Doe"}}
	var result bson.D
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}
	fmt.Printf("Found document: %v\n", result)

	// 3. Update
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "city", Value: "San Francisco"}}}}
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}
	fmt.Printf("Matched %v documents and modified %v documents\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// 4. Delete
	//deleteResult, err := collection.DeleteOne(ctx, filter)
	//if err != nil {
	//	log.Fatalf("Failed to delete document: %v", err)
	//}
	//fmt.Printf("Deleted %v documents\n", deleteResult.DeletedCount)
}
