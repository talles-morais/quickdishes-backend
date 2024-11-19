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

func DeleteOrder(orderId string) *AppError {
	var order models.Order
	result := database.DB.Where("order_id = ?", orderId).Delete(&order)

	if result.RowsAffected == 0 {
		return &AppError{Status: http.StatusNotFound, Message: "Order not found"}
	}

	if result.Error != nil {
		return &AppError{Status: http.StatusBadRequest, Message: result.Error.Error()}
	}
	return nil
}

func GetOrderById(orderId string) (*models.Order, *AppError) {
	var order models.Order

	if err := database.DB.Preload("Products").First(&order, "order_id = ?", orderId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &AppError{
				Status:  http.StatusNotFound,
				Message: "Order not found",
			}
		}
		return nil, &AppError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &order, nil
}

func GetAllOrders() ([]models.Order, *AppError) {
	var orders []models.Order

	if err := database.DB.Preload("Products").Find(&orders).Error; err != nil {
		return nil, &AppError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return orders, nil
}

func UpdateOrder(orderId string, updatedOrder models.Order) (*models.Order, *AppError) {
	var order models.Order

	if err := database.DB.First(&order, "order_id = ?", orderId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &AppError{
				Status:  http.StatusNotFound,
				Message: "Order not found",
			}
		}
		return nil, &AppError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := database.DB.Model(&order).Updates(updatedOrder).Error; err != nil {
		return nil, &AppError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := database.DB.Model(&order).UpdateColumn("pickup", updatedOrder.Pickup).Error; err != nil {
		return nil, &AppError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &order, nil
}
