package controllers

import (
	"github.com/bumblebee47293/to_do_app/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos := models.GetAllTodos()
	c.JSON(200, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.CreateTodo(&todo)
	c.JSON(201, todo)
}

// Additional handlers for update, delete...
