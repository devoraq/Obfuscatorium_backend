package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/devoraq/Obfuscatorium_backend/internal/config"
	"github.com/jackc/pgx/v5"
)

// Postgres представляет подключение к базе данных
type Postgres struct {
	Conn *pgx.Conn
}

// New создает новое подключение к PostgreSQL
func New(cfg config.DatabaseConfig) (*Postgres, error) {
	// Создание подключения
	conn, err := pgx.Connect(context.Background(), cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Ping(ctx); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{Conn: conn}, nil
}

// Close закрывает подключение к БД
func (p *Postgres) Close() {
	if p.Conn != nil {
		p.Conn.Close(context.Background())
	}
}

// Health проверяет доступность БД
func (p *Postgres) Health(ctx context.Context) error {
	return p.Conn.Ping(ctx)
}

// GetConn возвращает соединение (для использования в storages)
func (p *Postgres) GetConn() *pgx.Conn {
	return p.Conn
}
