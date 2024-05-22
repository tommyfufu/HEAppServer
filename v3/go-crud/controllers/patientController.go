package controllers

import (
	"context"
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

func CreatePatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var patient models.Patient
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.CreatePatient(db, patient); err != nil {
			// Check if the error is due to a duplicate key
			if mongo.IsDuplicateKeyError(err) {
				http.Error(w, "A patient with this email already exists.", http.StatusConflict)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(patient)
	}
}

func GetPatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var query bson.M
		var patient models.Patient
		collection := db.Database(config.MongodbDatabase).Collection("patients")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := r.URL.Query().Get("id")
		email := r.URL.Query().Get("email")
		phone := r.URL.Query().Get("phone")

		if id != "" {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			query = bson.M{"_id": objID}
		} else if email != "" {
			query = bson.M{"email": email}
		} else if phone != "" {
			query = bson.M{"phone": phone}
		} else {
			http.Error(w, "No valid identifier provided", http.StatusBadRequest)
			return
		}

		err := collection.FindOne(ctx, query).Decode(&patient)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				http.Error(w, "Patient not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(patient)
	}
}

func GetAllPatients(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patients, err := models.GetAllPatients(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(patients)
	}
}

func GetPatientMedication(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		patient, err := models.GetPatientMedication(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(patient.Medications)
	}
}

func GetPatientMessages(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		patient, err := models.GetPatient(db, id)
		if err != nil {
			http.Error(w, "Failed to fetch patient: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(patient.Messages)
	}
}

func UpdatePatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var patient models.Patient
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.UpdatePatient(db, id, patient); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(patient)
	}
}

type MedicationUpdate struct {
	Medication []models.MedicationType `json:"medication"`
}

func UpdatePatientMedication(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var medUpdate MedicationUpdate

		if err := json.NewDecoder(r.Body).Decode(&medUpdate); err != nil {
			log.Printf("Error decoding medication update: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Updating medication for patient %s with data: %v", id, medUpdate)
		loc, _ := time.LoadLocation("Asia/Taipei")
		formattedTime := time.Now().In(loc).Format("2006-01-02 15:04")
		log.Printf("formattedTime: %s", formattedTime)

		if err := models.UpdatePatientMedication(db, id, medUpdate.Medication); err != nil {
			log.Printf("Error updating medication: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		updatedPatient, err := models.GetPatient(db, id)
		if err != nil {
			log.Printf("Error fetching updated patient: %v", err)
			http.Error(w, "Failed to fetch updated patient", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(updatedPatient)
	}
}

func AddPatientMessage(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var message models.Message
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Set the current timestamp in a human-readable format (TW timezone)
		loc, _ := time.LoadLocation("Asia/Taipei")
		message.Date = time.Now().In(loc).Format("2006-01-02 15:04")

		update := bson.M{
			"$push": bson.M{"messages": message},
		}

		err := models.UpdatePatientField(db, id, update)
		if err != nil {
			http.Error(w, "Failed to update messages: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent) // No content to return, indicating successful operation
	}
}

func DeletePatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := models.DeletePatient(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNoContent)
	}
}
