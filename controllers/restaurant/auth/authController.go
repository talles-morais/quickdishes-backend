package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
	"github.com/talles-morais/quick-dishes/utils"
)

const secretKey = "segredo"

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

	newRestaurant := models.Restaurant{
		CNPJ:     restaurant.CNPJ,
		Name:     restaurant.Name,
		Email:    restaurant.Email,
		Password: encryptedPassword,
		Address:  restaurant.Address,
		Phone:    restaurant.Phone,
	}

	database.DB.Create(&newRestaurant)
	ctx.JSON(http.StatusCreated, newRestaurant)
}

func LoginRestaurant(ctx *gin.Context) {
	var loginData struct {
		CNPJ     string `json:"cnpj" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// bind to struct
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// search on DB
	var restaurant models.Restaurant
	if err := database.DB.Where("cnpj = ?", loginData.CNPJ).First(&restaurant).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid cnpj or password",
		})
		return
	}

	// verify hashed password
	if err := utils.VerifyPassword(restaurant.Password, loginData.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	expTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnpj": restaurant.CNPJ,
		"exp":  expTime, // 1 day
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not login",
		})
		return
	}

	ctx.SetCookie(
		"jwt",        // name
		token,        // value
		int(expTime), // maxAge
		"/",          // path
		"localhost",  // domain
		false,        // secure
		true,         // httpOnly
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}

func Restaurant(ctx *gin.Context) {
	cookie, err := ctx.Cookie("jwt")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "could not parse claims",
		})
		return
	}

	cnpj := (*claims)["cnpj"].(string)

	var restaurant models.Restaurant
	if err := database.DB.Where("cnpj = ?", cnpj).First(&restaurant).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "restaurant not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"claims": restaurant,
	})
}
