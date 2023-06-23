package repositories

import (
	"fmt"
	"todo-go/types"
)

func GetAllTodos() ([]*types.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return nil, fmt.Errorf("Error while connecting database[GetAllTodos() in repositories/todo.go]\n%v", err)
	}
	rows, err := db.Query("SELECT * FROM todos")
	defer db.Close()
	if err != nil {
		return nil, fmt.Errorf("Error database query[GetAllTodos() in repositories/todo.go]\n%v", err)
	}
	todos := make([]*types.Todo, 0)
	for rows.Next() {
		row := types.Todo{}
		err := rows.Scan(&row.ID, &row.Task, &row.Description, &row.Created_at, &row.Updated_at)
		if err != nil {
			return nil, fmt.Errorf("Error while scanning rows[GetAllTodos() in repositories/todo.go]\n%v", err)
		}
		todos = append(todos, &row)
	}

	return todos, nil
}

func CreateTodo(task string, description string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("Error while connecting database[CreateTodo() in repositories/todo.go]\n%v", err)
	}
	if _, err := db.Exec(
		fmt.Sprintf("INSERT INTO todos (task, description) VALUES (\"%s\", \"%s\");",
			task,
			description),
	); err != nil {
		return fmt.Errorf("Error while inserting data[CreateTodo() in repositories/todo.go]\n%v", err)
	}
	return nil

}
