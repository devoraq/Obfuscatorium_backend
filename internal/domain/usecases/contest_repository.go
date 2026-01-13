package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
)

type ContestRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Contest, error)
	Create(ctx context.Context, contest *models.Contest) error
	Update(ctx context.Context, contest *models.Contest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

