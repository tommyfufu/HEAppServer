package controllers

import (
	"database/sql"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func GetPatient(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func GetAllPatients(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func UpdatePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}

func DeletePatient(mysqldb *sql.DB, db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation of creating a doctor in MongoDB
	}
}
