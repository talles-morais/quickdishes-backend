package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/models"
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
	orderId := ctx.Param("id")

	if err := services.DeleteOrder(orderId); err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func GetOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")

	order, err := services.GetOrderById(orderId)
	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func GetAllOrders(ctx *gin.Context) {
	orders, err := services.GetAllOrders()
	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Message,
		})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	var updatedOrder models.Order

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	order, appErr := services.UpdateOrder(orderId, updatedOrder)
	if appErr != nil {
		ctx.JSON(appErr.Status, gin.H{
			"error": appErr.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
