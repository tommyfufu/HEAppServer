package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    mongodbAddr     string
    MongodbUser     string
    MongodbPass     string
    MongodbDatabase string
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    mongodbAddr = os.Getenv("MONGO_ADDR")
    MongodbUser = os.Getenv("MONGO_USER")
    MongodbPass = os.Getenv("MONGO_PASS")
    MongodbDatabase = os.Getenv("MONGO_DATABASE")

    // Debug: Print the environment variables
    log.Printf("MONGO_ADDR: %s", mongodbAddr)
    log.Printf("MONGO_USER: %s", MongodbUser)
    log.Printf("MONGO_PASS: %s", MongodbPass)
    log.Printf("MONGO_DATABASE: %s", MongodbDatabase)
}

func ConnectMongoDB() *mongo.Client {
    mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s/", MongodbUser, MongodbPass, mongodbAddr)
    log.Printf("MongoDB URI: %s", mongodbURI) // Debug: Print the MongoDB URI

    clientOptions := options.Client().ApplyURI(mongodbURI)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB %v, %v", MongodbDatabase, err)
    }

    pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer pingCancel()

    if err = client.Ping(pingCtx, nil); err != nil {
        log.Fatalf("Error pinging MongoDB %v, %v", MongodbDatabase, err)
    }

    log.Printf("Connected to MongoDB %v", MongodbDatabase)
    return client
}
