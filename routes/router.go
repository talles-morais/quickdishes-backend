package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/controllers/order"
	"github.com/talles-morais/quick-dishes/controllers/restaurant"
	"github.com/talles-morais/quick-dishes/services"
)

func Router() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	router.GET("/ws", func(ctx *gin.Context) {
		services.HandleWebSocket(ctx)
	})

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"hello": "world"}) })
	router.POST("/login", restaurant.LoginRestaurant)
	router.POST("/signup", restaurant.CreateRestaurant)
	router.GET("/restaurant", restaurant.Restaurant)
	router.POST("/logout", restaurant.Logout)
	
	router.POST("/order", order.CreateOrder)
	router.GET("/order/:id", order.GetOrderById)
	router.GET("/orders", order.GetAllOrders)
	router.GET("/orders/:id", order.GetAllOrdersByRestaurant)
	router.PUT("/order/:id", order.UpdateOrder)
	router.DELETE("/order/:id", order.DeleteOrder)

	go services.StartBroadcast()
	router.Run()
}
