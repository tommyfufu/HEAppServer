package config

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (

	// HEWebdb (MongoDB), which is used to store HEWeb users' information
	mongodbAddr     = "X"
	MongodbUser     = "X"
	MongodbPass     = "X"
	MongodbDatabase = "X"
)

func ConnectMongoDB() *mongo.Client {
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s/", MongodbUser, MongodbPass, mongodbAddr)
	clientOptions := options.Client().ApplyURI(mongodbURI)

	// Create a context with a 10-second timeout that can be used for the initial connection step
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB using the defined context
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB %v, %v", MongodbDatabase, err)
	}

	// Use a new context to check the MongoDB connection
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB %v, %v", MongodbDatabase, err)
	}

	log.Printf("Connected to MongoDB %v", MongodbDatabase)
	return client
}
