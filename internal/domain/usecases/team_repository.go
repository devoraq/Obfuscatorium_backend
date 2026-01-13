package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
)

type TeamRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Team, error)
	Create(ctx context.Context, team *models.Team) error
	Update(ctx context.Context, team *models.Team) error
	Delete(ctx context.Context, id uuid.UUID) error
}
