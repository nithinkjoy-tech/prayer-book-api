package controllers

import (
	"fmt"
	"net/http"
	"prayer-book/models"
	"prayer-book/services"

	"github.com/gin-gonic/gin"
)

func GetPrayers(c *gin.Context) {
	fmt.Println("get prayer controller")

	prayer, err := services.GetPrayersService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch prayers"})
		return
	}

	c.JSON(http.StatusOK, prayer)
}

func GetPrayerById() {
	fmt.Println("get prayer by ID controller")
}

func InsertPrayer(c *gin.Context) {
	var prayer models.Prayer
	if err := c.ShouldBindJSON(&prayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.InsertPrayer(prayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prayer successfully created",
		"id": id})
	fmt.Println("create prayer controller")
}
