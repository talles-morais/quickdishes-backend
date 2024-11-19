package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
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
