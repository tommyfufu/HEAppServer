package main

import (
	"context"
	"log"
	"net/http"
	"time"

	config "go-crud/config"
	"go-crud/models"
	"go-crud/routes"

	"github.com/gorilla/handlers"
)

func main() {

	// delcare Database
	HEDB := config.ConnectMongoDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer HEDB.Disconnect(ctx)

	// Initialize indexes for Patients collection
	if err := models.InitPatientIndexes(HEDB); err != nil {
		log.Fatalf("Failed to initialize patient indexes: %v", err)
	}

	// Initialize indexes for Patients collection
	if err := models.InitRecordIndexes(HEDB); err != nil {
		log.Fatalf("Failed to initialize patient indexes: %v", err)
	}

	// setup router
	r := routes.SetupRoutes(HEDB)
	// Setup CORS
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:4200", "http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)

	// HTTP server
	server := &http.Server{
		Handler:      corsOpts(r),
		Addr:         "140.113.174.44:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Starting server on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
