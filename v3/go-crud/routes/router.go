package routes

import (
	"go-crud/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	// "yourapp/controllers"
)

// SetupRoutes configures the application routes
func SetupRoutes(db *mongo.Client) *mux.Router {
	r := mux.NewRouter()

	// Doctor routes
	// r.HandleFunc("/doctor", controllers.CreateDoctor(webUserDB)).Methods("POST")
	// r.HandleFunc("/doctor/{id}", controllers.GetDoctor(webUserDB)).Methods("GET")
	// r.HandleFunc("/doctor/{id}", controllers.UpdateDoctor(webUserDB)).Methods("PUT")
	// r.HandleFunc("/doctor/{id}", controllers.DeleteDoctor(webUserDB)).Methods("DELETE")

	// Patient routes
	r.HandleFunc("/patient", controllers.CreatePatient(db)).Methods("POST")
	r.HandleFunc("/patients", controllers.GetAllPatients(db)).Methods("GET") // For reading all patients
	r.HandleFunc("/patient/{id}", controllers.GetPatient(db)).Methods("GET")
	r.HandleFunc("/patient/{id}", controllers.UpdatePatient(db)).Methods("PUT")
	r.HandleFunc("/patient/{id}/medication", controllers.UpdatePatientMedication(db)).Methods("PATCH")
	r.HandleFunc("/patient/{id}", controllers.DeletePatient(db)).Methods("DELETE")

	// Record routes
	r.HandleFunc("/record", controllers.CreateRecord(db)).Methods("POST")
	r.HandleFunc("/record/{userid}", controllers.GetRecordsByUserID(db)).Methods("GET")
	r.HandleFunc("/record/{id}", controllers.DeleteRecord(db)).Methods("DELETE")

	return r
}
