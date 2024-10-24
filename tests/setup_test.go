package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/quick-dishes/database"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func SetupDatabase() {
	database.ConnectDB()
}
