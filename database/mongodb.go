package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectMongo(uri string) {
	if uri == "" {
		panic("Mongo URI is valid")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to MongoDB")

	DB = client
}

func GetCollection(dbName, collName string) *mongo.Collection {
	return DB.Database(dbName).Collection(collName)
}
