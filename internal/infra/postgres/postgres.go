package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/devoraq/Obfuscatorium_backend/internal/config"
	"github.com/jmoiron/sqlx"
)

// Postgres представляет подключение к базе данных
type Postgres struct {
	DB *sqlx.DB
}

// New создает новое подключение к PostgreSQL
func New(cfg config.DatabaseConfig) (*Postgres, error) {
	// Создание подключения
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{DB: db}, nil
}

// Close закрывает подключение к БД
func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

// Health проверяет доступность БД
func (p *Postgres) Health(ctx context.Context) error {
	return p.DB.Ping()
}

// GetConn возвращает соединение (для использования в storages)
func (p *Postgres) GetConn() *sqlx.DB {
	return p.DB
}
