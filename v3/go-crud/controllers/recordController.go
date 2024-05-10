package controllers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateRecord handles the creation of a new record
func CreateRecord(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := models.CreateRecord(db, record); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// GetRecordsByUserID fetches all records for a specific user
func GetRecordsByUserID(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(vars["userid"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		records, err := models.GetRecordsByUserID(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
