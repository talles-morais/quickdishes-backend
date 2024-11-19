package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/controllers/order"
	"github.com/talles-morais/quick-dishes/controllers/restaurant"
)

func Router() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"hello": "world"}) })
	router.POST("/login", restaurant.LoginRestaurant)
	router.POST("/signup", restaurant.CreateRestaurant)
	router.GET("/restaurant", restaurant.Restaurant)
	router.POST("/logout", restaurant.Logout)

	router.POST("/order", order.CreateOrder)
	router.GET("/order/:id", order.GetOrderById)
	router.GET("/orders", order.GetAllOrders)
	router.PUT("/order/:id", order.UpdateOrder)
	router.DELETE("/order/:id", order.DeleteOrder)
	router.Run()
}
