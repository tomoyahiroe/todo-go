package services

import (
	"fmt"
	"todo-go/config"
	"todo-go/types"
)

func GetTodos() []*types.Todo {
	db, err := config.ConnectDB()
	if err != nil {
		if err.Error() == config.ErrDBConnectionFailed {
			panic("DB ConnectionFailed")
		}
		panic(fmt.Sprintf("unknown error:%v", err))
	}
	rows, err := db.Query("SELECT * FROM todos")
	defer db.Close()
	if err != nil {
		panic(fmt.Sprintf("unknown error:%v", err))
	}
	todos := make([]*types.Todo, 0)
	for rows.Next() {
		row := types.Todo{}
		err := rows.Scan(&row.ID, &row.Task, &row.Description, &row.Created_at, &row.Updated_at)
		if err != nil {
			panic(fmt.Sprintf("unknown error:%v", err))
		}
		todos = append(todos, &types.Todo{row.ID, row.Task, row.Description, row.Created_at, row.Updated_at})
	}

	return todos
}
