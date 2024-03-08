package models

import (
	"context"
	"fmt"
	"go-crud/config"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateDoctor(t *testing.T) {
	// Setup: Connect to your test database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Example MongoDB URI with authentication:
	// "mongodb://username:password@host:port/database"
	// Adjust the URI below according to your MongoDB setup
	dsn := fmt.Sprintf("mongodb://%s:%s@140.113.151.61:27017/%s", config.MongodbUser, config.MongodbPass, config.MongodbTestDatabase)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Ensure you have selected the correct database and collection for your test
	db := client.Database("yourTestDatabaseName") // Replace with your test database name
	collection := db.Collection("doctors")        // Ensure this is the collection you want to test against

	// Define a sample doctor to insert
	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}

	// Assuming the email is unique and used to identify the inserted test data
	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	} else {
		// Attempt to find the inserted doctor to confirm creation
		var foundDoctor Doctor
		err = collection.FindOne(ctx, bson.M{"email": sampleDoctor.Email}).Decode(&foundDoctor)
		if err != nil {
			t.Errorf("Failed to find the created doctor: %v", err)
		} else if foundDoctor.Email != sampleDoctor.Email {
			t.Errorf("Created doctor's email does not match. Expected %s, got %s", sampleDoctor.Email, foundDoctor.Email)
		}

		// Cleanup
		result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
		if err != nil {
			t.Fatalf("Failed to clean up test data: %v", err)
		}
		if result.DeletedCount == 0 {
			t.Logf("No documents deleted for email %s", sampleDoctor.Email)
		}
	}

}
