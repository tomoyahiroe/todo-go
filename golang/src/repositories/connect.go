package repositories

import (
	"database/sql"
	"fmt"
	"os"
	"todo-go/config"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	if err := config.LoadEnv(); err != nil {
		return nil, fmt.Errorf("Error while loading env[connectDB in repositories/connect.go]\n%v", err)
	}
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		return nil, fmt.Errorf("Error while connecting database[connectDB() in repositories/connect.go]\n%v", err)
	}
	return db, nil
}
