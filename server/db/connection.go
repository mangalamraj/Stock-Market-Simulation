package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectToMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	MongoDB = client.Database("stock-market-simulation")
	log.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(database, collection string) *mongo.Collection {
	if MongoDB == nil {
		log.Fatalf("MongoDB is not initialized. Call ConnectToMongo first.")
	}
	return MongoDB.Collection(collection)
}
