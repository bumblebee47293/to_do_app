package models

import "github.com/bumblebee47293/to_do_app/database"

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetAllTodos() []Todo {
	var todos []Todo
	database.DB.Find(&todos)
	return todos
}

func CreateTodo(todo *Todo) {
	database.DB.Create(todo)
}

// Additional functions for update, delete...
