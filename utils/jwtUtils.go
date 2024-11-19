package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "segredo"

func GenerateJwtToken(cnpj string) (string, error) {
	expTime := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnpj": cnpj,
		"exp":  expTime,
	})

	return claims.SignedString([]byte(secretKey))
}

func GetJwtClaim(ctx *gin.Context, claimKey string) (string, error) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		return "", errors.New("authentication required")
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(*jwt.MapClaims)
	if value, ok := (*claims)[claimKey].(string); ok {
		return value, nil
	}

	return "", errors.New("invalid claims")
}

func SetJwtCookie(ctx *gin.Context, token string) {
	ctx.SetCookie("jwt", token, 3600*24, "/", "localhost", false, true)
}

func ClearJwtCookie(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)
}
