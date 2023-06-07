package main

import (
	_ "encoding/json"
	_ "fmt"
	_ "io"
	"net/http"
	"todo-go/controllers"

	_ "context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		responce, err := controllers.Read()
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(200, responce)
	})
	r.POST("/create", func(c *gin.Context) {
		var todo_create struct {
			Task        string `json:"task"`
			Description string `json:"description"`
		}
		if err := c.ShouldBindJSON(&todo_create); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"task": todo_create.Task, "description": todo_create.Description})

	})
	r.Run()
}
