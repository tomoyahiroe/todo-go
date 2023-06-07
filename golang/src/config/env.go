package config

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(filepath.Join("config/", "local.env"))
	if err != nil {
		return fmt.Errorf("cannot laod env file: %v", err)
	}
	return nil
}
