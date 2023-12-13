package controllers

import (
	"net/http"

	"github.com/gabriel-kimutai/noiter/initializers"
	"github.com/gabriel-kimutai/noiter/models"
	"github.com/gin-gonic/gin"
)

func DishCreate(ctx *gin.Context) {
	// Get data off request body
	var data struct {
		Name        string
		Description string
		Price       float32
		Rating      float32
	}

	ctx.Bind(&data)

	// Insert record to database
	foodItem := models.Dish{Name: data.Name, Description: data.Description, Price: data.Price, Rating: data.Rating}
	initializers.DB.Create(&foodItem)

	// Return the record
	ctx.JSON(http.StatusOK, gin.H{
		"foodItems": foodItem,
	})
}

func DishesFetchAll(ctx *gin.Context) {
	foods := []models.Dish{}

	initializers.DB.Find(&foods)

	ctx.JSON(http.StatusOK, gin.H{
		"foodItems": foods,
	})
}

func DishesFetchCategory(ctx *gin.Context) {
	// category = ctx.Param("name")
	foods := []models.Dish{}

	initializers.DB.Find(&foods)

	ctx.JSON(http.StatusOK, gin.H{
		"foodItems": foods,
	})
}