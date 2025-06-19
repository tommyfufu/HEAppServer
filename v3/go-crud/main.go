package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "time"

    config "go-crud/config"
    "go-crud/models"
    "go-crud/routes"

    "github.com/gorilla/handlers"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on environment variables")
    }

    // Declare Database
    HEDB := config.ConnectMongoDB()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    defer HEDB.Disconnect(ctx)

    // Initialize indexes for Patients collection
    if err := models.InitPatientIndexes(HEDB); err != nil {
        log.Fatalf("Failed to initialize patient indexes: %v", err)
    }

    // Initialize indexes for Records collection
    if err := models.InitRecordIndexes(HEDB); err != nil {
        log.Fatalf("Failed to initialize record indexes: %v", err)
    }

    // Setup router
    r := routes.SetupRoutes(HEDB)

    // Setup CORS
    corsOpts := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://13.236.164.131:4200",}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
    )

    serverAddr := os.Getenv("SERVER_ADDR")
    if serverAddr == "" {
        serverAddr = "0.0.0.0:8090" // Default address if not set
    }

    // HTTP server
    server := &http.Server{
        Handler:      corsOpts(r),
        Addr:         serverAddr,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    // Start the server
    log.Println("Starting server on", server.Addr)
    log.Fatal(server.ListenAndServe())
}
