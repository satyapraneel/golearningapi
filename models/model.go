package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	dbName := "netflix"
	collectionName := "watchlist"
	connectionString := "mongodb+srv://satyapraneel:hady4viVctX3dJuk@cluster0.4v4ognw.mongodb.net/?retryWrites=true&w=majority"
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
}
