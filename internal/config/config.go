package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config содержит конфигурацию приложения
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	SMTP     SMTPConfig
}

// ServerConfig содержит настройки HTTP сервера
type ServerConfig struct {
	Address        string
	AllowedOrigins string
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
	Mode     string
}

// DatabaseConfig содержит настройки подключения к БД
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// Load загружает конфигурацию из переменных окружения
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env not found, falling back to environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Address:        os.Getenv("SERVER_ADDRESS"),
			AllowedOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}, nil
}

// DSN возвращает строку подключения к PostgreSQL
func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode,
	)
}
