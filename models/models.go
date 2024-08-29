package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetAllTodos(db *gorm.DB) []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func CreateTodo(db *gorm.DB, todo *Todo) {
	db.Create(todo)
}

// Additional functions for update, delete...
