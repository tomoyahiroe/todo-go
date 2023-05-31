package main

import (
	"todo-go/controllers"

	_ "context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		responce := controllers.Read()
		c.JSON(200, responce)
	})
	r.Run()
}
