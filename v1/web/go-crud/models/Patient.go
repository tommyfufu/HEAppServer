package models

import (
	"context"
	"go-crud/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Patient struct {
	ID       string `bson:"_id,omitempty"` // MongoDB ID
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Birthday string `bson:"birthday"`
}

func CreatePatient(db *mongo.Client, p Patient) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")

	_, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Printf("Error creating new patient in MongoDB: %v", err)
		return err
	}

	return nil
}

func GetPatient(db *mongo.Client, id string) (*Patient, error) {
	var patient Patient
	collection := db.Database(config.MongodbDatabase).Collection("patients")

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&patient)
	if err != nil {
		log.Printf("Error finding patient with ID %s: %v", id, err)
		return nil, err
	}

	return &patient, nil
}

func GetAllPatients(db *mongo.Client) ([]Patient, error) {
	var patients []Patient

	collection := db.Database(config.MongodbDatabase).Collection("patients")
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

func UpdatePatient(db *mongo.Client, id string, p Patient) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")

	update := bson.M{
		"$set": p,
	}
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("Error updating patient with ID %s: %v", id, err)
		return err
	}

	return nil
}

func DeletePatient(db *mongo.Client, id string) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error deleting patient with ID %s: %v", id, err)
		return err
	}

	return nil
}
