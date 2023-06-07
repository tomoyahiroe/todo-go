package services

import (
	"fmt"
	"todo-go/repositories"
	"todo-go/types"
)

func GetTodos() ([]*types.Todo, error) {
	todos, err := repositories.GetAllTodos()
	if err != nil {
		return nil, fmt.Errorf("error while getting all todos: %v", err)
	}
	return todos, nil
}
