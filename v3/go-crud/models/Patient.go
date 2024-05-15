package models

import (
	"context"
	"go-crud/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Patient struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Email           string             `bson:"email"`
	Phone           string             `bson:"phone"`
	Birthday        string             `bson:"birthday"`
	Gender          string             `bson:"gender"`
	AsusvivowatchSN string             `bson:"asusvivowatchsn"` //Asus Vivowatch Serial Number
	PhotoSticker    string             `bson:"photosticker"`
	Messages        map[string]string  `bson:"messages"`
	Medications     []MedicationType   `bson:"medication"`
}

type MedicationType struct {
	Name      string `bson:"name"`
	Dosage    string `bson:"dosage"`    // Number of pills
	Frequency string `bson:"frequency"` // Times per day
}

func InitPatientIndexes(db *mongo.Client) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Index for the email field
	emailIndexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // Index in ascending order
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, emailIndexModel)
	if err != nil {
		log.Printf("Error creating unique index on email in patients collection: %v", err)
		return err
	}

	// Index for the phone field
	phoneIndexModel := mongo.IndexModel{
		Keys:    bson.M{"phone": 1}, // Index in ascending order
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(ctx, phoneIndexModel)
	if err != nil {
		log.Printf("Error creating unique index on phone in patients collection: %v", err)
		return err
	}

	return nil
}

func CreatePatient(db *mongo.Client, p Patient) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, p)
	if err != nil {
		log.Printf("Error creating new patient in MongoDB: %v", err)
		return err
	}

	return nil
}

func GetPatient(db *mongo.Client, id string) (*Patient, error) {
	var patient Patient
	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID %s to ObjectID: %v", id, err)
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&patient)
	if err != nil {
		log.Printf("Error finding patient with ID %s: %v", id, err)
		return nil, err
	}

	return &patient, nil
}

func GetAllPatients(db *mongo.Client) ([]Patient, error) {
	var patients []Patient

	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error retrieving patients from MongoDB: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": p}
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("Error updating patient with ID %s: %v", id, err)
		return err
	}

	return nil
}

func UpdatePatientMedication(db *mongo.Client, id string, medications []MedicationType) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to ObjectID: %v", err)
		return err
	}

	update := bson.M{"$set": bson.M{"medication": medications}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error updating patient with ID %s: %v", id, err)
		return err
	}

	return nil
}

func DeletePatient(db *mongo.Client, id string) error {
	collection := db.Database(config.MongodbDatabase).Collection("patients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Error deleting patient with ID %s: %v", id, err)
		return err
	}

	return nil
}
