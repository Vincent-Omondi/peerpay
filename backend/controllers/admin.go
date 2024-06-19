// controllers/admin.go
package controllers

import (
	"net/http"

	"peerpay/backend/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	if err := models.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func GetDormantAccounts(c *gin.Context) {
	var dormantUsers []models.User
	// Logic to find dormant accounts
	// e.g., if err := models.DB.Where("last_login < ?", someTime).Find(&dormantUsers).Error; err != nil {
	c.JSON(http.StatusOK, gin.H{"dormant_users": dormantUsers})
}
