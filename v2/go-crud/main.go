package main

import (
	"context"
	"log"
	"net/http"

	config "go-crud/config"
	routes "go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	webUserDB := config.ConnectMongoDB()
	defer webUserDB.Disconnect(context.TODO())

	// Setup router with Gin
	router := gin.Default()

	// Enable detailed logging
	router.Use(gin.Logger())

	// Setup CORS
	router.Use(ginCors())

	// Setup API routes
	routes.SetupRoutes(router, webUserDB)

	// Start the server
	log.Println("Starting server on http://127.0.0.1:8090")
	log.Fatal(router.Run(":8090"))
}

func ginCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
