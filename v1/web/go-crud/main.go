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
	corsOpts := handlers.AllowedOrigins([]string{"*"})
	corsHandler := handlers.CORS(corsOpts)

	// HTTP server
	server := &http.Server{
		Handler:      corsHandler(r),
		Addr:         "127.0.0.1:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Starting server on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
