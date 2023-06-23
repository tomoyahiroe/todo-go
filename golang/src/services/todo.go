package services

import (
	"fmt"
	"todo-go/repositories"
	"todo-go/types"
)

func GetTodos() ([]*types.Todo, error) {
	todos, err := repositories.GetAllTodos()
	if err != nil {
		return nil, fmt.Errorf("Error while getting todos[GetTodos() in services/todo.go]\n%v", err)
	}
	return todos, nil
}

func CreateTodo(task string, description string) error {
	if err := repositories.CreateTodo(task, description); err != nil {
		return fmt.Errorf("Error while creating todo[CreateTodo() in services/todo.go]\n%v", err)
	}
	return nil
}
