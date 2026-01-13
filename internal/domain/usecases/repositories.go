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

type TeamRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Team, error)
	Create(ctx context.Context, team *models.Team) error
	Update(ctx context.Context, team *models.Team) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ContestRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Contest, error)
	Create(ctx context.Context, contest *models.Contest) error
	Update(ctx context.Context, contest *models.Contest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type JudgesRepositoryI interface {
	GetByID(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) (*models.User, error)
	AddJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
	DeleteJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
}
