package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/services"
)

func CreateOrder(ctx *gin.Context) {
	if err := services.CreateNewOrder(ctx); err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func DeleteOrder(ctx *gin.Context) {
	if err := services.DeleteOrder(ctx); err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func GetOrderById(ctx *gin.Context) {
	order, err := services.GetOrderById(ctx)
	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func GetAllOrders(ctx *gin.Context) {
	orders, err := services.GetAllOrders(ctx)
	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
