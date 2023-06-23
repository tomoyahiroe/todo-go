package main

import (
	_ "encoding/json"
	"fmt"
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
			c.JSON(500, gin.H{"error": string(err.Error())})
			return
		}
		c.JSON(200, responce)
	})
	r.POST("/create", func(c *gin.Context) {
		err := controllers.Create(c)
		if err != nil {
			err = fmt.Errorf("Error while creating todo[POST /create in main.go]\n%v", err)
			fmt.Println(err) // デバッグ用
			c.JSON(500,
				gin.H{
					"error": string(err.Error()),
				},
			)
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"message": "Todo created successfully"},
		)

	})
	r.Run()
}
