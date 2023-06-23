package controllers

import (
	"fmt"
	"todo-go/services"
	"todo-go/types"

	"github.com/gin-gonic/gin"
)

func Read() ([]*types.Todo, error) {
	todos, err := services.GetTodos()
	if err != nil {
		return nil, fmt.Errorf("Error while getting todos[Read() in controllers/todo.go]\n%v", err)
	}
	return todos, nil
}

func Create(c *gin.Context) error {
	var createTodo types.CreateTodo
	createTodo.Task = c.PostForm("task")
	createTodo.Description = c.PostForm("description")
	if err := services.CreateTodo(createTodo.Task, createTodo.Description); err != nil {
		return fmt.Errorf("Error while creating todo[Create() in controllers/todo.go]\n%v", err)
	}
	return nil
}
