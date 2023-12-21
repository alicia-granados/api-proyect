package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBName     string
	DBUser     string
	DBHost     string
	DBPort     string
	DBPassword string
}

func LoadEnvironment() (*AppConfig, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %w", err)
	}

	envFilePath := filepath.Join(dir, ".env")
	if err := godotenv.Load(envFilePath); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &AppConfig{
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
	fmt.Println(config)

	if config.DBName == "" {
		return nil, fmt.Errorf("DBName is required")
	}
	if config.DBUser == "" {
		return nil, fmt.Errorf("DBUser is required")
	}

	if config.DBHost == "" {
		return nil, fmt.Errorf("DBHost is required")
	}

	if config.DBPort == "" {
		return nil, fmt.Errorf("DBPort is required")
	}

	return config, nil
}

func DSN() string {
	config, err := LoadEnvironment()
	if err != nil {
		log.Fatalln("incomplete configuration:", err.Error())
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
}
