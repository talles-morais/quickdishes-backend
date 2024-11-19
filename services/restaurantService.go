package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
	"github.com/talles-morais/quick-dishes/utils"
)

type AppError struct {
	Status  int
	Message string
}

func CreateNewRestaurant(ctx *gin.Context) *AppError {
	var restaurant models.Restaurant

	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		return &AppError{Status: http.StatusBadRequest, Message: err.Error()}
	}

	if err := models.ValidateRestaurant(&restaurant); err != nil {
		return &AppError{Status: http.StatusBadRequest, Message: err.Error()}
	}

	encryptedPassword, err := utils.HashPassword(restaurant.Password)
	if err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: "Error encrypting password"}
	}

	restaurant.Password = encryptedPassword
	if err := database.DB.Create(&restaurant).Error; err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: "Failed to create restaurant"}
	}

	return nil
}

func AuthenticateRestaurant(ctx *gin.Context) (string, *AppError) {
	var loginData struct {
		CNPJ     string `json:"cnpj" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		return "", &AppError{Status: http.StatusBadRequest, Message: err.Error()}
	}

	var restaurant models.Restaurant
	if err := database.DB.Where("cnpj = ?", loginData.CNPJ).First(&restaurant).Error; err != nil {
		return "", &AppError{Status: http.StatusUnauthorized, Message: "Invalid CNPJ or password"}
	}

	if err := utils.VerifyPassword(restaurant.Password, loginData.Password); err != nil {
		return "", &AppError{Status: http.StatusUnauthorized, Message: "Invalid password"}
	}

	token, err := utils.GenerateJwtToken(restaurant.CNPJ)
	if err != nil {
		return "", &AppError{Status: http.StatusInternalServerError, Message: "Failed to generate token"}
	}

	return token, nil
}

func GetAuthenticatedRestaurant(ctx *gin.Context) (map[string]interface{}, *AppError) {
	cnpj, err := utils.GetJwtClaim(ctx, "cnpj")
	if err != nil {
		return nil, &AppError{Status: http.StatusUnauthorized, Message: err.Error()}
	}

	var restaurant models.Restaurant
	if err := database.DB.Where("cnpj = ?", cnpj).First(&restaurant).Error; err != nil {
		return nil, &AppError{Status: http.StatusNotFound, Message: "Restaurant not found"}
	}

	return map[string]interface{}{
		"cnpj":  restaurant.CNPJ,
		"name":  restaurant.Name,
		"email": restaurant.Email,
		"phone": restaurant.Phone,
	}, nil
}
