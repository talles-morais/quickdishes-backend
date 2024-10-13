package main

import (
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/routes"
)

func main() {
	database.ConnectDB()
	routes.Router()
}
