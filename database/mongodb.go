package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectMongo(uri string) {
	if uri == "" {
		log.Fatal("Mongo URI is valid")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection failed", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB ping error", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	DB = client
}

func GetCollection(dbName, collectionName string) *mongo.Collection {
	return DB.Database(dbName).Collection(collectionName)
}
