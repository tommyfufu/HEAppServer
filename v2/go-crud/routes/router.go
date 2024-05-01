package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configures the application routes
func SetupRoutes(r *gin.Engine, webUserDB *mongo.Client) {
	// Patient routes
	// api := r.Group("/heapp")
	// {
	patients := r.Group("/patients")
	{
		// Handle both with and without trailing slash
		patients.POST("", controllers.CreatePatient(webUserDB))
		patients.POST("/", controllers.CreatePatient(webUserDB))

		patients.GET("", controllers.GetAllPatients(webUserDB))
		patients.GET("/", controllers.GetAllPatients(webUserDB))

		patients.GET("/:id", controllers.GetPatient(webUserDB))
		patients.PUT("/:id", controllers.UpdatePatient(webUserDB))
		patients.PATCH("/:id/medication", controllers.UpdatePatientMedication(webUserDB))
		patients.DELETE("/:id", controllers.DeletePatient(webUserDB))
	}
	// }
}
