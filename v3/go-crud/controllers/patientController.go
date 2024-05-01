package controllers

import (
	"encoding/json"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

type MedicationUpdate struct {
	Medication []string `json:"medication"`
}

func UpdatePatientMedication(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var medUpdate MedicationUpdate
		// Decode the entire JSON body into the medUpdate struct
		if err := json.NewDecoder(r.Body).Decode(&medUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Updating medication for patient %s with data: %v", id, medUpdate)

		if err := models.UpdatePatientMedication(db, id, medUpdate.Medication); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		updatedPatient, err := models.GetPatient(db, id)
		if err != nil {
			http.Error(w, "Failed to fetch updated patient", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedPatient)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
