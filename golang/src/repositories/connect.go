package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connectDB() (*sql.DB, error) {
	loadEnv()
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")+")")
	if err != nil {
		return nil, errors.New(errDBConnectionFailed)
	}
	return db, nil
}

const errDBConnectionFailed = "DB connection failed"

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("cannot laod env file: %v", err)
	}
}
