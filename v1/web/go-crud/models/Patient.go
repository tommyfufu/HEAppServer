package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// You might need additional imports depending on your logic, such as for handling BSON
)

// Patient represents the structure of a patient's data
type Patient struct {
	ID       string `bson:"_id,omitempty"` // MongoDB ID
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Birthday string `bson:"birthday"`
	// Add other fields as necessary
}

// CreatePatient creates a new patient record in MongoDB
func CreatePatient(db *mongo.Client, p Patient) error {
	collection := db.Database("HEWEBdb").Collection("patients")

	_, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Printf("Error creating new patient in MongoDB: %v", err)
		return err
	}

	return nil
}

// GetAllPatients retrieves all patient records from MongoDB
func GetAllPatients(db *mongo.Client) ([]Patient, error) {
	var patients []Patient

	collection := db.Database("HEWEBdb").Collection("patients")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Printf("Error retrieving patients from MongoDB: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var patient Patient
		if err := cursor.Decode(&patient); err != nil {
			log.Printf("Error decoding patient data: %v", err)
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error with patient cursor: %v", err)
		return nil, err
	}

	return patients, nil
}

// UpdatePatient updates a patient's record in MongoDB
// Note: Implement the logic based on your application's needs

// DeletePatient deletes a patient's record from MongoDB
// Note: Implement the logic based on your application's needs
