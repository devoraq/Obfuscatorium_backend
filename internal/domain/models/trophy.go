package models

import (
	"time"

	"github.com/google/uuid"
)

type Trophy struct {
	ID          uuid.UUID
	Name        string
	Description string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
