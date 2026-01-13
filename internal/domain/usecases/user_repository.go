package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
)


type UserRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}


