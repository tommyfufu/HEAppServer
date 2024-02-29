package routes

import (
	"database/sql"
	"go-crud/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	// "yourapp/controllers"
)

// SetupRoutes configures the application routes
func SetupRoutes(appUserDB *sql.DB, webUserDB *mongo.Client) *mux.Router {
	r := mux.NewRouter()

	// Doctor routes
	r.HandleFunc("/doctor", controllers.CreateDoctor(webUserDB)).Methods("POST")
	r.HandleFunc("/doctor/{id}", controllers.GetDoctor(webUserDB)).Methods("GET")
	r.HandleFunc("/doctor/{id}", controllers.UpdateDoctor(webUserDB)).Methods("PUT")
	r.HandleFunc("/doctor/{id}", controllers.DeleteDoctor(webUserDB)).Methods("DELETE")

	// Patient routes
	r.HandleFunc("/patient", controllers.CreatePatient(appUserDB, webUserDB)).Methods("POST")
	r.HandleFunc("/patients", controllers.GetAllPatients(webUserDB)).Methods("GET") // For reading all patients
	r.HandleFunc("/patient/{id}", controllers.GetPatient(webUserDB)).Methods("GET")
	r.HandleFunc("/patient/{id}", controllers.UpdatePatient(appUserDB, webUserDB)).Methods("PUT")
	r.HandleFunc("/patient/{id}", controllers.DeletePatient(appUserDB, webUserDB)).Methods("DELETE")

	return r
}
