package main

import (
	"net/http"

	"github.com/gabriel-kimutai/noiter/controllers"
	"github.com/gabriel-kimutai/noiter/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/dishes")
	})

	// Get and create dishes
	r.GET("/dishes", controllers.DishesFetchAll)
	r.POST("/dishes", controllers.DishCreate)

	// Get and create tables
	r.POST("/tables", controllers.TableCreate)
	r.GET("/tables", controllers.TablesFetchAll)

	// Get and create categories
	r.POST("/categories", controllers.CategoryCreate)
	r.GET("/categories", controllers.CategoryFetchAll)
	r.GET("/categories/:name", controllers.CategoryFetchOne)

	r.Run()
}