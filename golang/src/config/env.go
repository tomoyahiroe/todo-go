package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("cannot laod env file: %v", err)
	}
}
