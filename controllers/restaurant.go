package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
	"github.com/talles-morais/quick-dishes/utils"
)

func CreateRestaurant(ctx *gin.Context) {
	var restaurant models.Restaurant

	// bind to struct
	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validate the restaurant info
	if err := models.ValidateRestaurant(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// hash the password
	encryptedPassword, err := utils.HashPassword(restaurant.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error encrypting password",
		})
		return
	}

	newRestaurant := models.Restaurant {
		CNPJ: restaurant.CNPJ,
		Name: restaurant.Name,
		Email: restaurant.Email,
		Password: encryptedPassword,
		Address: restaurant.Address,
		Phone: restaurant.Phone,
	}

	database.DB.Create(&newRestaurant)
	ctx.JSON(http.StatusCreated, newRestaurant)
}
