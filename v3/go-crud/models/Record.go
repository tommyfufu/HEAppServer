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

type GameRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id" json:"userId"`
	GameID       int                `bson:"game_id" json:"gameId"`
	GameDateTime string             `bson:"game_date_time" json:"gameDateTime"`
	GameTime     string             `bson:"game_time" json:"gameTime"`
	Score        int                `bson:"score" json:"score"`
}

// InitRecordIndexes initializes indexes for the records collection
func InitRecordIndexes(db *mongo.Client) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Creating an index for the UserID field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}}, // Index in ascending order
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Error creating index on user_id in records collection: %v", err)
		return err
	}

	return nil
}

// CreateRecord inserts a new record into the records collection
func CreateRecord(db *mongo.Client, record GameRecord) error {
	collection := db.Database(config.MongodbDatabase).Collection("records")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, record); err != nil {
		log.Printf("Error creating new record in MongoDB: %v", err)
		return err
	}

	return nil
}

// GetRecordsByUserID retrieves all records for a specific user from the records collection
func GetRecordsByUserID(db *mongo.Client, userID primitive.ObjectID) ([]GameRecord, error) {
	var records []GameRecord
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
		var record GameRecord
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

// DeleteRecord removes a record from the records collection based on its ID
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
