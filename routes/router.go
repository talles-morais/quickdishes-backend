package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/controllers"
)

func Router() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"hello": "world"}) })
	router.POST("/login", controllers.LoginRestaurant)
	router.POST("/signup", controllers.CreateRestaurant)
	router.Run()
}
