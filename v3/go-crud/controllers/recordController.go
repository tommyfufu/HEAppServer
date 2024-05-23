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
			UserID   string `json:"userId"`
			GameID   int    `json:"gameId"`
			GameTime string `json:"gameTime"`
			Score    int    `json:"score"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		uid, err := primitive.ObjectIDFromHex(input.UserID)
		if err != nil {
			http.Error(w, "Invalid user ID format", http.StatusBadRequest)
			return
		}

		loc, _ := time.LoadLocation("Asia/Taipei")
		formattedTime := time.Now().In(loc).Format("2006-01-02 15:04")
		log.Printf("formattedTime: %s", formattedTime)

		record := models.GameRecord{
			UserID:       uid,
			GameID:       input.GameID,
			GameDateTime: formattedTime,
			GameTime:     input.GameTime,
			Score:        input.Score,
		}

		collection := db.Database(config.MongodbDatabase).Collection("records")
		_, err = collection.InsertOne(r.Context(), record)
		if err != nil {
			http.Error(w, "Failed to create record: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(record)
	}
}

// GetRecordsByUserID fetches all game records for a specific user.
func GetRecordsByUserID(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(vars["userId"]) // Ensure the key matches your route definition
		if err != nil {
			log.Printf("Error converting user ID from hex: %v, received ID: %s", err, vars["userId"])
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
