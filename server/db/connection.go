package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongo(uri string) error {
	// Set a timeout for connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping MongoDB to confirm the connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	// Set the global MongoClient
	MongoClient = client
	log.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(database, collection string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatalf("MongoClient is not initialized. Call ConnectToMongo first.")
	}
	return MongoClient.Database(database).Collection(collection)
}