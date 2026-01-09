package main

import (
	"log"
	
	"github.com/devoraq/Obfuscatorium_backend/internal/config"
	"github.com/devoraq/Obfuscatorium_backend/internal/infra/postgres"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к БД
	db, err := postgres.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")
}