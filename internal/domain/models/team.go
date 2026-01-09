package models

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID          uuid.UUID
	Name        string
	Description string
	CaptainID   int
	IsPrivate   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
