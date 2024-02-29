package main

import (
	"context"
	"log"
	"net/http"
	"time"

	config "go-crud/config"
	"go-crud/routes"
)

func main() {

	appUserDB := config.ConnectMysqlDB()
	defer appUserDB.Close()
	webUserDB := config.ConnectMongoDB()
	defer webUserDB.Disconnect(context.TODO())

	r := routes.SetupRoutes(appUserDB, webUserDB)

	// Setup your routes
	// routes.SetupRoutes(r, db)

	// Define an HTTP server
	server := &http.Server{
		Handler:      r,
		Addr:         "140.113.151.61:8090", // Adjust the address and port accordingly
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Starting server on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
