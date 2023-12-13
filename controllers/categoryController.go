package controllers

import (
	"net/http"

	"github.com/gabriel-kimutai/noiter/initializers"
	"github.com/gabriel-kimutai/noiter/models"
	"github.com/gin-gonic/gin"
)

func CategoryCreate(ctx *gin.Context) {
	var data struct {
		Name        string
		Description string
	}

	ctx.Bind(&data)

	category := models.Category{Name: data.Name, Description: data.Description}

	initializers.DB.Create(&category)

	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
		"message":  "category created",
	})
}

// db.Model(&User{}).Preload("CreditCards").Find(&users)

func CategoryFetchAll(ctx *gin.Context) {
	categories := []models.Category{}

	initializers.DB.Model(&models.Category{}).Preload("Dishes").Find(&categories)

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func CategoryFetchOne(ctx *gin.Context) {
	name := ctx.Param("name")
	category := models.Category{}
	// initializers.DB.Where(&models.Category{Name: name}).First(&category)
	initializers.DB.Model(&models.Category{}).Where("name = ?", name).Preload("Dishes").First(&category)
	
	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
