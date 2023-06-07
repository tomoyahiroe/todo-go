package controllers

import (
	"todo-go/services"
	"todo-go/types"
)

func Read() ([]*types.Todo, error) {
	todos, err := services.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// func Create(c *gin.Context) {
// 	var todo_create struct {
// 		Task        string `json:"task"`
// 		Description string `json:"description"`
// 	}
// 	if err := c.ShouldBindJSON(&todo_create); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"task": todo_create.Task, "description": todo_create.Description})

// }
