package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	Database   DatabaseConfig
}

type HTTPServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return Config{
		HTTPServer: HTTPServerConfig{
			Port: 8080,
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
}
