package controllers

import (
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePatient(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := models.CreatePatient(db, patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, patient)
	}
}

func GetPatient(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		patient, err := models.GetPatient(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

func GetAllPatients(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, err := models.GetAllPatients(db)
		log.Printf("GetAllPatients in controller")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, patients)
	}
}

func UpdatePatient(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := models.UpdatePatient(db, id, patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

type MedicationUpdate struct {
	Medication []string `json:"medication"`
}

func UpdatePatientMedication(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var medUpdate MedicationUpdate
		if err := c.ShouldBindJSON(&medUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := models.UpdatePatientMedication(db, id, medUpdate.Medication); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		updatedPatient, err := models.GetPatient(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated patient"})
			return
		}

		c.JSON(http.StatusOK, updatedPatient)
	}
}

func DeletePatient(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := models.DeletePatient(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
