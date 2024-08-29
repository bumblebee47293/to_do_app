package main

import (
	"todo-app/config"
	"todo-app/database"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDatabase()

	router := gin.Default()
	routes.InitializeRoutes(router)

	router.Run(":8080")
}
