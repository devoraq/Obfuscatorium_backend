package models

import (
	"time"

	"github.com/google/uuid"
)

type Contest struct {
	ID          uuid.UUID
	Name        string
	Description string
	IsActive    bool
	IsTeam      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
