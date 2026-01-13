package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
)

type JudgesRepositoryI interface {
	GetByID(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) (*models.User, error)
	AddJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
	DeleteJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
}
