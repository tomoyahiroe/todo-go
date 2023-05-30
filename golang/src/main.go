package main

import (
	"fmt"
	"todo-go/config"

	_ "context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		todos := getTodos()
		fmt.Println(todos)
		c.JSON(200, todos)
	})
	r.Run()
}

type todo struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func getTodos() []*todo {
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
	todos := make([]*todo, 0)
	for rows.Next() {
		var id int
		var task string
		var description string
		var created_at string
		var updated_at string
		if err := rows.Scan(&id, &task, &description, &created_at, &updated_at); err != nil {
			panic(fmt.Sprintf("unknown error:%v", err))
		}
		todos = append(todos, &todo{id, task, description, created_at, updated_at})
		fmt.Println(id, task, description, created_at, updated_at)
	}

	return todos
}
