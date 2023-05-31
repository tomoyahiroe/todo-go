package controllers

import (
	"todo-go/services"
	"todo-go/types"
)

func Read() []*types.Todo {
	todos := services.GetTodos()
	return todos
}
