package controllers

import (
	"net/http"

	"github.com/gabriel-kimutai/noiter/initializers"
	"github.com/gabriel-kimutai/noiter/models"
	"github.com/gin-gonic/gin"
)

func TableCreate(ctx *gin.Context) {
	table := models.Table{}
	initializers.DB.Create(&table)

	ctx.JSON(http.StatusOK, gin.H{
		"table": table,
		"message": "table created",
	})
}

func TablesFetchAll(ctx *gin.Context) {
	tables := []models.Table{}
	initializers.DB.Find(&tables)

	ctx.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}