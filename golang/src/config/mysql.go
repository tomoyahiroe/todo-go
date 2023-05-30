package config

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

const ErrDBConnectionFailed = "DB connection failed"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/todo-db")
	if err != nil {
		return nil, errors.New(ErrDBConnectionFailed)
	}
	// defer db.Close()

	return db, nil
}
