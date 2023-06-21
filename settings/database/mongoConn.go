package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() (*mongo.Collection, error) {
	connectionString := "mongodb://localhost:27017/learn"
	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("Failed to create MongoDB client:", err)
	}

	// Connect MongoDB server
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	}

	//verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	database := client.Database("learn")
	collection := database.Collection("learn")
	return collection, nil
}
