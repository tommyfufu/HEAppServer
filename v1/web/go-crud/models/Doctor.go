package models

import (
	"context"
	"go-crud/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Doctor struct {
	ID         string `bson:"_id,omitempty"` // MongoDB ID
	Name       string `bson:"name"`
	Email      string `bson:"email"`
	Phone      string `bson:"phone"`
	Department string `bson:"department"`
}

func CreateDoctor(db *mongo.Client, d Doctor) error {
	collection := db.Database(config.MongodbDatabase).Collection("doctors")
	_, err := collection.InsertOne(context.TODO(), d)
	if err != nil {
		log.Printf("Error creating new doctor in %v: %v", config.MongodbDatabase, err)
	}
	return nil
}

func UpdateDoctor(db *mongo.Client, id string, d *Doctor) error {
	collection := db.Database(config.MongodbDatabase).Collection("doctors")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": *d}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Error updating doctor in %v: %v", config.MongodbDatabase, err)
		return err
	}
	return nil
}

func GetDoctor(db *mongo.Client, email string) (*Doctor, error) {
	collection := db.Database(config.MongodbDatabase).Collection("doctors")
	var doctor Doctor
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&doctor)
	if err != nil {
		log.Printf("Error finding doctor by email %v in %v: %v", email, config.MongodbDatabase, err)
		return nil, err
	}
	return &doctor, nil
}

func DeleteDoctor(db *mongo.Client, id string) error {
	collection := db.Database(config.MongodbDatabase).Collection("doctors")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error deleting doctor with ID %v in %v: %v", id, config.MongodbDatabase, err)
		return err
	}
	return nil
}
