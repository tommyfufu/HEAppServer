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

type Record struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	GameID       int                `bson:"game_id"`
	GameDateTime string             `bson:"game_date_time"`
	GameTime     string             `bson:"game_time"`
	Score        int                `bson:"score"`
}

// Initializes indexes for the records collection
func InitRecordIndexes(db *mongo.Client) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Creating an index for the UserID field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"user_id": 1}, // Index in ascending order
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Error creating index on user_id in records collection: %v", err)
		return err
	}

	return nil
}

func CreateRecord(db *mongo.Client, record Record) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, record); err != nil {
		log.Printf("Error creating new record in MongoDB: %v", err)
		return err
	}

	if err := EnsureMaxRecords(db, record.UserID); err != nil {
		log.Printf("Error enforcing max records limit: %v", err)
		return err
	}

	return nil
}

func GetRecordsByUserID(db *mongo.Client, userID primitive.ObjectID) ([]Record, error) {
	var records []Record
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error retrieving records for user with ID %s: %v", userID.Hex(), err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var record Record
		if err := cursor.Decode(&record); err != nil {
			log.Printf("Error decoding record: %v", err)
			return nil, err
		}
		records = append(records, record)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error processing records cursor: %v", err)
		return nil, err
	}

	return records, nil
}

func DeleteRecord(db *mongo.Client, id string) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to ObjectID: %v", err)
		return err
	}

	if _, err := collection.DeleteOne(ctx, bson.M{"_id": objID}); err != nil {
		log.Printf("Error deleting record with ID %s: %v", id, err)
		return err
	}

	return nil
}

func EnsureMaxRecords(db *mongo.Client, userID primitive.ObjectID) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Count the number of records for the user
	count, err := collection.CountDocuments(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error counting records for user with ID %s: %v", userID.Hex(), err)
		return err
	}

	// If the number of records exceeds 20, delete the oldest record
	if count > 20 {
		collection.FindOneAndDelete(ctx, bson.M{"user_id": userID}, options.FindOneAndDelete().SetSort(bson.M{"game_date_time": 1}))
	}

	return nil
}
