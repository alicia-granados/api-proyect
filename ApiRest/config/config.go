package config

import (
	"fmt"
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
	ServerPort string
	ServerHost string
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
		ServerPort: os.Getenv("SERVER_PORT"),
		ServerHost: os.Getenv("SERVER_HOST"),
	}

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

	if config.ServerPort == "" {
		return nil, fmt.Errorf("ServerPort is required")
	}

	if config.ServerHost == "" {
		return nil, fmt.Errorf("ServerHost is required")
	}
	return config, nil
}

func DSNDatabase() (string, error) {
	config, err := LoadEnvironment()
	if err != nil {
		return "", fmt.Errorf("error loading environment configuration: %w", err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName), nil
}

func DSNServer() (string, error) {
	config, err := LoadEnvironment()
	if err != nil {
		return "", fmt.Errorf("error loading environment configuration: %w", err)
	}

	return fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort), nil
}
