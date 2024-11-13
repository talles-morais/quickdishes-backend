package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/controllers/restaurant/auth"
)

func Router() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"hello": "world"}) })
	router.POST("/login", auth.LoginRestaurant)
	router.POST("/signup", auth.CreateRestaurant)
	router.GET("/restaurant", auth.Restaurant)
	router.Run()
}
