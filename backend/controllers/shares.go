// controllers/shares.go
package controllers

import (
	"net/http"

	"peerpay/backend/models"

	"github.com/gin-gonic/gin"
)

func BuyShares(c *gin.Context) {
	var share models.Share
	if err := c.ShouldBindJSON(&share); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&share).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy shares"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shares bought successfully"})
}

func SellShares(c *gin.Context) {
	var share models.Share
	if err := c.ShouldBindJSON(&share); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Delete(&share).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sell shares"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shares sold successfully"})
}

func GetShares(c *gin.Context) {
	var shares []models.Share
	userID := c.Param("userID")

	if err := models.DB.Where("user_id = ?", userID).Find(&shares).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shares not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shares": shares})
}
