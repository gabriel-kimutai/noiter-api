package main

import (

	"github.com/gabriel-kimutai/noiter/initializers"
	"github.com/gabriel-kimutai/noiter/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Dish{},
		&models.Order{},
		&models.Table{},
		&models.Category{},
	)

	initializers.InitializeCategories()
	initializers.InitializeTables()

	if err != nil {
		panic("Migration failed")
	}
}
