package models

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateDoctor(t *testing.T) {
	// Setup: Connect to your test database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("140.113.151.61:8090"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Define a sample doctor to insert
	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}

	// Call the function under test
	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	}

	// Add cleanup code to remove any data added to the database during the test
	// Add assertions to check if the doctor was correctly inserted, if necessary
}
