package main

import (
	"github.com/bumblebee47293/to_do_app/config"
	"github.com/bumblebee47293/to_do_app/database"
	"github.com/bumblebee47293/to_do_app/models"
	"github.com/bumblebee47293/to_do_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDatabase()

	// Auto-migrate the Todo model
	database.DB.AutoMigrate(&models.Todo{})

	router := gin.Default()
	routes.InitializeRoutes(router)

	router.Run(":8080")

	DB.AutoMigrate(&models.Todo{})
}
