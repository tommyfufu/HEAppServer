package controllers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func GetDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func UpdateDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func DeleteDoctor(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}
