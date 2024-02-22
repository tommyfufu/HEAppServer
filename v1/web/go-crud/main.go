package main

import (
	"context"
	"log"
	"net/http"
	"time"

	config "go-crud/config"
	"go-crud/routes"

	"github.com/gorilla/handlers"
)

func main() {

	// delcare Database
	appUserDB := config.ConnectMysqlDB()
	defer appUserDB.Close()
	webUserDB := config.ConnectMongoDB()
	defer webUserDB.Disconnect(context.TODO())

	// setup router
	r := routes.SetupRoutes(appUserDB, webUserDB)
	// Setup CORS
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:4200", "http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)

	// HTTP server
	server := &http.Server{
		Handler:      corsOpts(r),
		Addr:         "127.0.0.1:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Starting server on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
