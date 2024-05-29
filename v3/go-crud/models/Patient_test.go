package models

import (
	"context"

	"go-crud/config"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreatePatient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("patients")
	// testfunc: Create
	for i := 0; i < len(TestUsePatientsData); i++ {
		samplePatient := TestUsePatientsData[i]
		err = CreatePatient(client, TestUsePatientsData[i])
		if err != nil {
			t.Errorf("Failed to create doctor: %v", err)
		} else {
			// check create
			var foundPatient Patient
			err = collection.FindOne(ctx, bson.M{"email": samplePatient.Email}).Decode(&foundPatient)
			if err != nil {
				t.Errorf("Failed to find the created doctor: %v", err)
			} else if foundPatient.Email != samplePatient.Email {
				t.Errorf("Created doctor's email does not match. Expected %s, got %s", samplePatient.Email, foundPatient.Email)
			}

			// clean up
			result, err := collection.DeleteOne(ctx, bson.M{"email": samplePatient.Email})
			if err != nil {
				t.Fatalf("Failed to clean up test data: %v", err)
			}
			if result.DeletedCount == 0 {
				t.Logf("No documents deleted for email %s", samplePatient.Email)
			}
		}
	}

}

func TestGetPatient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.MongodbDatabase)
	collection := db.Collection("patients")

	// testfunc: Get
	var samplePatient Patient
	if err := collection.FindOne(ctx, bson.M{"name": "John Doe"}).Decode(&samplePatient); err != nil {
		t.Fatalf("Failed to find sample patient: %v", err)
	}
	samplePatientID := samplePatient.ID.Hex()
	// GetPatient function is called with the correct ID
	gotPatient, err := GetPatient(client, samplePatientID)
	if err != nil {
		t.Errorf("Failed to get patient: %v", err)
	} else if gotPatient.ID != samplePatient.ID {
		t.Errorf("GetPatient() returned wrong patient. Expected ID %v, got ID %v", samplePatient.ID, gotPatient.ID)
	}
	// gotPatientJSON, err := json.MarshalIndent(gotPatient, "", "  ")
	// if err != nil {
	// 	t.Logf("Failed to marshal gotPatient data: %v", err)
	// } else {
	// 	t.Logf("Got Patient Data:\n%s\n", string(gotPatientJSON))
	// }

}

func TestGetAllPatients(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	gotPatients, err := GetAllPatients(client)
	if err != nil {
		t.Errorf("Failed to get all patients: %v", err)
	}
	if len(gotPatients) == 0 {
		t.Errorf("GetAllPatients() returned no patients")
	} else {
		t.Logf("Retrieved %d patients", len(gotPatients))
	}

	// for _, patient := range gotPatients {
	// 	t.Logf("Patient ID: %s, Name: %s", patient.ID.Hex(), patient.Name)
	// }

}

// func TestUpdatePatient(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer client.Disconnect(ctx)

// 	db := client.Database(config.MongodbDatabase)
// 	collection := db.Collection("doctors")

// 	sampleDoctor := Doctor{
// 		Name:  "Test Doctor",
// 		Email: "test@example.com",
// 	}

// 	err = CreateDoctor(client, sampleDoctor)
// 	if err != nil {
// 		t.Errorf("Failed to create doctor: %v", err)
// 	} else {
// 		// testfunc: Update
// 		updatedDoctor := Doctor{
// 			Name:  "Rename Doctor",
// 			Email: "test@example.com",
// 		}
// 		err := UpdateDoctor(client, sampleDoctor.ID, &updatedDoctor)
// 		if err != nil {
// 			t.Errorf("Failed to find the created doctor: %v", err)
// 		} else if updatedDoctor.Name == sampleDoctor.Name {
// 			t.Errorf("Update doctor's email does not match. Expected %s, got %s", sampleDoctor.Name, updatedDoctor.Name)
// 		}
// 		log.Printf("UpdateDoctor() successed %v", updatedDoctor)

// 		// clean up
// 		result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
// 		if err != nil {
// 			t.Fatalf("Failed to clean up test data: %v", err)
// 		}
// 		if result.DeletedCount == 0 {
// 			t.Logf("No documents deleted for email %s", sampleDoctor.Email)
// 		}
// 	}

// }

// func TestDeletePatient(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer client.Disconnect(ctx)

// 	db := client.Database(config.MongodbDatabase)
// 	collection := db.Collection("doctors")

// 	sampleDoctor := Doctor{
// 		Name:  "Test Doctor",
// 		Email: "test@example.com",
// 	}

// 	err = CreateDoctor(client, sampleDoctor)
// 	if err != nil {
// 		t.Errorf("Failed to create doctor: %v", err)
// 	} else {
// 		// testfunc: Delete

// 		err := DeleteDoctor(client, sampleDoctor.ID)
// 		if err != nil {
// 			t.Errorf("Failed to delete the created doctor: %v", err)
// 			// clean up
// 			result, err := collection.DeleteOne(ctx, bson.M{"email": sampleDoctor.Email})
// 			if err != nil {
// 				t.Fatalf("Failed to clean up test data: %v", err)
// 			}
// 			if result.DeletedCount == 0 {
// 				t.Logf("No documents deleted for email %s", sampleDoctor.Email)
// 			}
// 		}
// 		log.Printf("Delete Doctor() successed")

// 	}

// }

var TestUsePatientsData = []Patient{
	{
		Name:         "John Doe",
		Email:        "johndoe@example.com",
		Phone:        "555-0100",
		Birthday:     "1980-04-12",
		PhotoSticker: "SomeBase64OrHexEncodedPhoto",
		// Messages:     map[string]string{"2024-03-11 17:00:00": "Hello", "2024-03-10 07:00:00": "World"},
		// Medications:  []string{["Med1", ], "Med2"},
	},
	{
		Name:         "Jane Doe",
		Email:        "janedoe@example.com",
		Phone:        "555-0200",
		Birthday:     "1985-05-15",
		PhotoSticker: "SomeBase64OrHexEncodedPhoto",
		// Messages:     map[string]string{"2024-03-12 09:00:00": "Good morning", "2024-03-11 20:00:00": "Good night"},
		// Medications:  []string{"Med3", "Med4"},
	},
	{
		Name:         "Alice Johnson",
		Email:        "alicej@example.com",
		Phone:        "555-0300",
		Birthday:     "1990-07-22",
		PhotoSticker: "SomeBase64OrHexEncodedPhoto",
		// Messages:     map[string]string{"2024-03-13 12:00:00": "Lunch?", "2024-03-12 18:00:00": "Dinner"},
		// Medications:  []string{"Med5", "Med6"},
	},
	{
		Name:         "Bob Smith",
		Email:        "bobsmith@example.com",
		Phone:        "555-0400",
		Birthday:     "1975-02-28",
		PhotoSticker: "SomeBase64OrHexEncodedPhoto",
		// Messages:     map[string]string{"2024-03-14 10:00:00": "Meeting at 10", "2024-03-13 15:00:00": "Coffee break"},
		// Medications:  []string{"Med7", "Med8"},
	},
	{
		Name:         "Charlie Brown",
		Email:        "charlieb@example.com",
		Phone:        "555-0500",
		Birthday:     "1988-09-17",
		PhotoSticker: "SomeBase64OrHexEncodedPhoto",
		// Messages:     map[string]string{"2024-03-15 08:00:00": "Workout?", "2024-03-14 22:00:00": "Late snack"},
		// Medications:  []string{"Med9", "Med10"},
	},
}
