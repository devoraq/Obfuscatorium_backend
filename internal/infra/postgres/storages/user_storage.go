package storages

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/exceptions"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserStorage struct {
	DB *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{DB: db}
}

func (s *UserStorage) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT id, email, username, password_hash, avatar, bio, role, created_at, updated_at
			  FROM user_service.users WHERE id = $1`

	var user models.User
	err := s.DB.GetContext(ctx, &user, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrUserNotFound
		}
		return nil, exceptions.ErrDatabaseError
	}

	return &user, nil
}

func (s *UserStorage) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, email, username, password_hash, avatar, bio, role, created_at, updated_at
              FROM user_service.users WHERE username = $1`

	var user models.User
	err := s.DB.GetContext(ctx, &user, query, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrUserNotFound
		}
		return nil, exceptions.ErrDatabaseError
	}

	return &user, nil
}

func (s *UserStorage) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO user_service.users (email, username, password_hash)
	          VALUES ($1, $2, $3)
	          RETURNING id, email, username, password_hash, avatar, bio, role, created_at, updated_at`

	var createdUser models.User
	err := s.DB.GetContext(ctx, &createdUser, query,
		user.Email,
		user.Username,
		user.Password,
	)

	if err != nil {
		// Проверяем на нарушение уникального ограничения
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			// PostgreSQL error code 23505 = unique_violation
			if pgErr.Code == "23505" {
				// Определяем, какое поле нарушило уникальность
				if strings.Contains(pgErr.Constraint, "username") || strings.Contains(pgErr.Detail, "username") {
					return nil, exceptions.ErrUserAlreadyExists
				}
				if strings.Contains(pgErr.Constraint, "email") || strings.Contains(pgErr.Detail, "email") {
					return nil, exceptions.ErrUserAlreadyExists
				}
				return nil, exceptions.ErrUserAlreadyExists
			}
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &createdUser, nil
}

func (s *UserStorage) Update(ctx context.Context, id uuid.UUID, updates map[string]any) (*models.User, error) {
	if len(updates) == 0 {
		return nil, exceptions.ErrNoFieldsToUpdate
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	builder := psql.Update("user_service.users").
		Where(squirrel.Eq{"id": id}).
		Set("updated_at", squirrel.Expr("NOW()")).
		SetMap(updates).
		Suffix("RETURNING id, email, username, password_hash, avatar, bio, role, created_at, updated_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, exceptions.ErrQueryBuildFailed
	}

	var updatedUser models.User
	err = s.DB.GetContext(ctx, &updatedUser, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrUserNotFound
		}
		// Проверяем на нарушение уникального ограничения
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, exceptions.ErrUserAlreadyExists
			}
		}
		return nil, exceptions.ErrQueryExecutionFailed
	}

	return &updatedUser, nil
}

func (s *UserStorage) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM user_service.users WHERE id = $1`

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
