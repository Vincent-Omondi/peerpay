// controllers/profile.go
package controllers

import (
	"log" // Add logging
	"net/http"

	"peerpay/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the JWT token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Cast userID to uint
	profile.UserID = userID.(uint)

	log.Printf("Received profile creation request for user_id: %d", profile.UserID)

	if err := models.DB.Create(&profile).Error; err != nil {
		log.Printf("Error creating profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile created successfully"})
}

func EditProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&profile).Error; err != nil {
		log.Printf("Error updating profile: %v", err) // Add logging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func GetProfile(c *gin.Context) {
	var profile models.Profile
	userID := c.Param("userID")

	if err := models.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		log.Printf("Profile not found: %v", err) // Add logging
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}
