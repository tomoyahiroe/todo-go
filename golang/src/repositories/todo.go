package repositories

import (
	"fmt"
	"todo-go/types"
)

func GetAllTodos() ([]*types.Todo, error) {
	db, err := connectDB()
	if err != nil {
		if err.Error() == errDBConnectionFailed {
			return nil, err
		}
		return nil, fmt.Errorf("unknown error:%v", err)
	}
	rows, err := db.Query("SELECT * FROM todos")
	defer db.Close()
	if err != nil {
		return nil, fmt.Errorf("database query error:%v", err)
	}
	todos := make([]*types.Todo, 0)
	for rows.Next() {
		row := types.Todo{}
		err := rows.Scan(&row.ID, &row.Task, &row.Description, &row.Created_at, &row.Updated_at)
		if err != nil {
			panic(fmt.Sprintf("unknown error:%v", err))
		}
		todos = append(todos, &row)
	}

	return todos, nil
}
