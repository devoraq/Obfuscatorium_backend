package storages

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
