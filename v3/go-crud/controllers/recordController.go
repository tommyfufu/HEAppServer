package controllers

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRecord(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			UserID   string `json:"user_id"`
			GameID   int    `json:"game_id"`
			GameTime string `json:"game_time"`
			Score    int    `json:"score"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the UserID and convert to ObjectID
		uid, err := primitive.ObjectIDFromHex(input.UserID)
		if err != nil {
			log.Printf("Error converting UserID to ObjectID: %v", err)
			http.Error(w, "Invalid user ID format", http.StatusBadRequest)
			return
		}

		record := models.GameRecord{
			UserID:       uid,
			GameID:       input.GameID,
			GameDateTime: time.Now(),
			GameTime:     input.GameTime,
			Score:        input.Score,
		}

		log.Printf("Attempting to create record with UserID: %s", record.UserID.Hex())

		collection := db.Database(config.MongodbDatabase).Collection("records")
		result, err := collection.InsertOne(r.Context(), record)
		if err != nil {
			log.Printf("Failed to insert record: %v", err)
			http.Error(w, "Failed to create record: "+err.Error(), http.StatusInternalServerError)
			return
		}

		record.ID = result.InsertedID.(primitive.ObjectID)
		log.Printf("Record created with ID: %s", record.ID.Hex())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(record)
	}
}

// GetRecordsByUserID fetches all game records for a specific user.
func GetRecordsByUserID(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(vars["user_id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		collection := db.Database(config.MongodbDatabase).Collection("records")
		filter := bson.M{"user_id": userID}
		cursor, err := collection.Find(r.Context(), filter)
		if err != nil {
			http.Error(w, "Failed to fetch records", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(r.Context())

		var records []models.GameRecord
		if err = cursor.All(r.Context(), &records); err != nil {
			http.Error(w, "Failed to parse records", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(records)
	}
}

// DeleteRecord handles the deletion of a record
func DeleteRecord(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		recordID := vars["id"]
		if err := models.DeleteRecord(db, recordID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
