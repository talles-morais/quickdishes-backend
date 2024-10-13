package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"hello": "world"}) })
	router.Run()
}
