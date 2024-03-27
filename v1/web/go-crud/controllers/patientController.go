package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patient models.Patient
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.CreatePatient(db, patient); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(patient)
	}
}

func GetPatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		patient, err := models.GetPatient(db, id)
		if err != nil {
			http.Error(w, "Patient not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patients)
	}
}

func UpdatePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

type MedicationUpdate struct {
	Medication []string `json:"medication"`
}

func UpdatePatientMedication(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var medUpdate MedicationUpdate
		if err := json.NewDecoder(r.Body).Decode(&medUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Preprocessing on the front end, no need to process like below
		// Convert string ID to MongoDB ObjectID
		// objID, err := primitive.ObjectIDFromHex(id)
		// if err != nil {
		// 	http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		// 	return
		// }

		if err := models.UpdatePatientMedication(db, id, medUpdate.Medication); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(medUpdate.Medication)
	}
}

func DeletePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := models.DeletePatient(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent) // No content to return upon successful deletion
	}
}
