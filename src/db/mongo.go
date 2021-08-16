package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/configs"
)

var MongoClient *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI(configs.MongoURI) // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	MongoClient = client.Database(configs.DBName)
}
