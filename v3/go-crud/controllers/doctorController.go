package controllers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctor models.Doctor
		err := json.NewDecoder(r.Body).Decode(&doctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = models.CreateDoctor(db, doctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(doctor)
	}
}

func GetDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// doctorEmail := vars["id"]
		doctorEmail := vars["email"]
		doctor, err := models.GetDoctor(db, doctorEmail)
		if err != nil {
			http.Error(w, "Doctor not found: "+err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctor)
	}
}

func UpdateDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		doctorId := vars["id"]

		var doctor models.Doctor
		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.UpdateDoctor(db, doctorId, &doctor); err != nil {
			http.Error(w, "Failed to update doctor: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctor)
	}
}

func DeleteDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		doctorId := vars["id"]

		if err := models.DeleteDoctor(db, doctorId); err != nil {
			http.Error(w, "Failed to delete doctor: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
