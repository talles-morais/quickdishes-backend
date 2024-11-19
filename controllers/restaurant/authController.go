package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/services"
	"github.com/talles-morais/quick-dishes/utils"
)

func CreateRestaurant(ctx *gin.Context) {
	if err := services.CreateNewRestaurant(ctx); err != nil {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Restaurant created successfully"})
}

func LoginRestaurant(ctx *gin.Context) {
	token, err := services.AuthenticateRestaurant(ctx)
	if err != nil {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	utils.SetJwtCookie(ctx, token)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Restaurant(ctx *gin.Context) {
	restaurant, err := services.GetAuthenticatedRestaurant(ctx)
	if err != nil {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}

func Logout(ctx *gin.Context) {
	utils.ClearJwtCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
