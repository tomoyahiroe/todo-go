package repositories

import (
	"database/sql"
	"errors"
	"os"
	"todo-go/config"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	config.LoadEnv()
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")+")")
	if err != nil {
		return nil, errors.New(errDBConnectionFailed)
	}
	return db, nil
}

const errDBConnectionFailed = "DB connection failed"
