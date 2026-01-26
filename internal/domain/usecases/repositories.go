package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/google/uuid"
)

type UserRepositoryI interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, id uuid.UUID, updates map[string]any) (*models.User, error)
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
	List(ctx context.Context) ([]*models.Contest, error)

	ChangeStatus(ctx context.Context, id uuid.UUID, newStatus models.ContestStatus) error
	RegisterUser(ctx context.Context, contestID uuid.UUID, userID uuid.UUID) error
	RegisterTeam(ctx context.Context, contestID uuid.UUID, teamID uuid.UUID) error
	Unregister(ctx context.Context, contestID uuid.UUID, participantID uuid.UUID) error
	ListTeamParticipants(ctx context.Context, contestID uuid.UUID) ([]*models.Team, error)
	ListUserParticipants(ctx context.Context, contestID uuid.UUID) ([]*models.User, error)
}

type JudgesRepositoryI interface {
	GetByID(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) (*models.User, error)
	AddJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
	DeleteJudge(ctx context.Context, userID uuid.UUID, contestID uuid.UUID) error
}
