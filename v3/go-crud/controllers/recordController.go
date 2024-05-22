package controllers

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRecord(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Controller CreateRecord")

		var record models.GameRecord
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Parse the user_id from the JSON body
		var data struct {
			UserID string `json:"user_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid or missing user ID", http.StatusBadRequest)
			return
		}

		uid, err := primitive.ObjectIDFromHex(data.UserID)
		if err != nil {
			http.Error(w, "Invalid user ID format", http.StatusBadRequest)
			return
		}
		record.UserID = uid
		record.GameDateTime = time.Now() // Setting current time as gameDateTime
		// fmt.Printf("record.UserId: %v, record.GameDateTime: %v", record.UserID, record.GameDateTime)
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
