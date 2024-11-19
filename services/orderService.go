package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
	"gorm.io/gorm"
)

func CreateNewOrder(ctx *gin.Context) *AppError {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		return &AppError{Status: http.StatusBadRequest, Message: err.Error()}
	}

	if err := models.ValidateOrder(&order); err != nil {
		return &AppError{Status: http.StatusBadRequest, Message: err.Error()}
	}

	if err := database.DB.Create(&order).Error; err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: "Failed to create order"}
	}
	return nil
}

func DeleteOrder(ctx *gin.Context) *AppError {
	var order models.Order
	orderId := ctx.Param("id")
	result := database.DB.Where("order_id = ?", orderId).Delete(&order)

	if result.RowsAffected == 0 {
		return &AppError{Status: http.StatusNotFound, Message: "Order not found"}
	}

	if result.Error != nil {
		return &AppError{Status: http.StatusBadRequest, Message: result.Error.Error()}
	}
	return nil
}

func GetOrderById(ctx *gin.Context) (*models.Order, *AppError) {
	var order models.Order
	orderId := ctx.Param("id")

	if err := database.DB.Preload("Products").First(&order, "order_id = ?", orderId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &AppError{
				Status:  http.StatusNotFound,
				Message: "Order not found",
			}
		}
		return nil, &AppError{
			Status: http.StatusInternalServerError, 
			Message: err.Error(),
		}
	}

	return &order, nil
}
