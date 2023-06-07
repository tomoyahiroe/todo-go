package services

import (
	"todo-go/repositories"
	"todo-go/types"
)

func GetTodos() ([]*types.Todo, error) {
	todos, err := repositories.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
