package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	AvatarURL    string
	TeamID       int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
