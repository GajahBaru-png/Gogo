package controllers

import (
	"gogo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemsInput struct {
	Name        string `json:"name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,gte=0"`
	Description string `json:"description"`
	AddedBy     string `json:"added_by"`
}

func Items(c *gin.Context) {
	var input ItemsInput
	var u models.User

	username := c.GetString("user")
	if username == "" {
		c.JSON(401, gin.H{"error": "User not found in token"})
		return
	}

	if err := models.DB.Where("username = ?", username).First(&u).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot input your data"})
		return
	}

	item := models.Items{Name: input.Name, Quantity: input.Quantity, Description: input.Description, AddedBy: u.Username}
	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Error"})
		return
	}

	models.DB.Create(&models.Activity{
		UserID: u.ID,
		Action: "Menambahkan Item",
		ItemID: item.ID,
	})

	if input.Quantity < 0 {
		c.JSON(400, gin.H{"error": "Quantity cannot be negative"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})

}

func FindItems(c *gin.Context) {
	var items []models.Items

	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})
}

type UpdateItemsInput struct {
	Name        string `json:"name"`
	Quantity    int    `json:"quantity" binding:"gte=0"`
	Description string `json:"description"`
	AddedBy     string `json:"added_by"`
}

func UpdateItems(c *gin.Context) {
	var item models.Items

	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input UpdateItemsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Males ah"})
		return
	}

	updatedItems := models.Items{Name: input.Name, Quantity: input.Quantity, Description: input.Description, AddedBy: input.AddedBy}

	models.DB.Model(&item).Updates(&updatedItems)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteItems(c *gin.Context) {
	ItemsID := c.Param("id")
	var item models.Items

	result := models.DB.First(&item, ItemsID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not Found"})
		return
	}

	models.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item delete succesfully"})
}
