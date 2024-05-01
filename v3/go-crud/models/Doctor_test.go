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

var (
	dsn = fmt.Sprintf("mongodb://%s:%s@140.113.151.61:27017/", config.MongodbUser, config.MongodbPass)
)

func TestCreateDoctor(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// dsn := fmt.Sprintf("mongodb://%s:%s@140.113.151.61:27017/", config.MongodbUser, config.MongodbPass)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("doctors")

	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}
	// testfunc: Create
	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	} else {
		// check create
		var foundDoctor Doctor
		err = collection.FindOne(ctx, bson.M{"email": sampleDoctor.Email}).Decode(&foundDoctor)
		if err != nil {
			t.Errorf("Failed to find the created doctor: %v", err)
		} else if foundDoctor.Email != sampleDoctor.Email {
			t.Errorf("Created doctor's email does not match. Expected %s, got %s", sampleDoctor.Email, foundDoctor.Email)
		}

		// clean up
		result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
		if err != nil {
			t.Fatalf("Failed to clean up test data: %v", err)
		}
		if result.DeletedCount == 0 {
			t.Logf("No documents deleted for email %s", sampleDoctor.Email)
		}
	}

}

func TestGetDoctor(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("doctors")

	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}

	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	} else {
		// testfunc: Get
		getdoctor, err := GetDoctor(client, sampleDoctor.Email)
		if err != nil {
			t.Errorf("Failed to find the doctor we want to get: %v", err)
		} else if getdoctor.Email != sampleDoctor.Email { //is this possibile, chat?
			t.Errorf("Get method no doctor found. Expected %s, got %s", sampleDoctor.Email, getdoctor.Email)
		}
		log.Printf("GetDoctor() successed %v", getdoctor)

		// clean up
		result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
		if err != nil {
			t.Fatalf("Failed to clean up test data: %v", err)
		}
		if result.DeletedCount == 0 {
			t.Logf("No documents deleted for email %s", sampleDoctor.Email)
		}
	}

}

func TestUpdateDoctor(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("doctors")

	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}

	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	} else {
		// testfunc: Update
		updatedDoctor := Doctor{
			Name:  "Rename Doctor",
			Email: "test@example.com",
		}
		err := UpdateDoctor(client, sampleDoctor.ID, &updatedDoctor)
		if err != nil {
			t.Errorf("Failed to find the created doctor: %v", err)
		} else if updatedDoctor.Name == sampleDoctor.Name {
			t.Errorf("Update doctor's email does not match. Expected %s, got %s", sampleDoctor.Name, updatedDoctor.Name)
		}
		log.Printf("UpdateDoctor() successed %v", updatedDoctor)

		// clean up
		result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
		if err != nil {
			t.Fatalf("Failed to clean up test data: %v", err)
		}
		if result.DeletedCount == 0 {
			t.Logf("No documents deleted for email %s", sampleDoctor.Email)
		}
	}

}

func TestDeleteDoctor(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("doctors")

	sampleDoctor := Doctor{
		Name:  "Test Doctor",
		Email: "test@example.com",
	}

	err = CreateDoctor(client, sampleDoctor)
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	} else {
		// testfunc: Delete

		err := DeleteDoctor(client, sampleDoctor.ID)
		if err != nil {
			t.Errorf("Failed to delete the created doctor: %v", err)
			// clean up
			result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
			if err != nil {
				t.Fatalf("Failed to clean up test data: %v", err)
			}
			if result.DeletedCount == 0 {
				t.Logf("No documents deleted for email %s", sampleDoctor.Email)
			}
		}
		log.Printf("Delete Doctor() successed")

	}

}
