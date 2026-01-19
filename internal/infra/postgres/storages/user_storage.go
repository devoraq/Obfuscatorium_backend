package storages

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	DB *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{DB: db}
}

func (s *UserStorage) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT id, email, username, password_hash, avatar, bio, role, created_at, updated_at
			  FROM users WHERE id = $1`

	var user models.User
	err := s.DB.GetContext(ctx, &user, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (s *UserStorage) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, email, username, password_hash, avatar, bio, role, created_at, updated_at
              FROM users WHERE username = $1`

	var user models.User
	err := s.DB.GetContext(ctx, &user, query, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (s *UserStorage) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users (email, username, password_hash)
	          VALUES ($1, $2, $3)
	          RETURNING id, email, username, password_hash, avatar, bio, role, created_at, updated_at`

	var createdUser models.User
	err := s.DB.GetContext(ctx, &createdUser, query,
		user.Email,
		user.Username,
		user.Password,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &createdUser, nil
}

func (s *UserStorage) Update(ctx context.Context, id uuid.UUID, updates map[string]any) (*models.User, error) {
	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	builder := psql.Update("users").
		Where(squirrel.Eq{"id": id}).
		Set("updated_at", squirrel.Expr("NOW()")).
		SetMap(updates).
		Suffix("RETURNING id, email, username, password_hash, avatar, bio, role, created_at, updated_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var updatedUser models.User
	err = s.DB.GetContext(ctx, &updatedUser, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db execution failed: %w", err)
	}

	return &updatedUser, nil
}

func (s *UserStorage) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
}
