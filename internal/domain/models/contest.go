package models

import (
	"time"

	"github.com/google/uuid"
)

type ContestStatus string

const (
	StatusUnspecified ContestStatus = "unspecified"
	StatusDraft       ContestStatus = "draft"
	StatusActive      ContestStatus = "active"
	StatusFinished    ContestStatus = "finished"
	StatusCancelled   ContestStatus = "cancelled"
)

type ContestType string

const (
	TypeIndividual ContestType = "individual"
	TypeTeam       ContestType = "team"
)

type Contest struct {
	ID                uuid.UUID     `db:"id"`
	Name              string        `db:"name"`
	Description       string        `db:"description"`
	Status            ContestStatus `db:"status"`
	Type              ContestType   `db:"type"`
	StartDate         time.Time     `db:"start_date"`
	EndDate           time.Time     `db:"end_date"`
	RegistrationStart time.Time     `db:"registration_start"`
	RegistrationEnd   time.Time     `db:"registration_end"`
	MaxParticipants   int           `db:"max_participants"`
	MaxTeams          int           `db:"max_teams"`
	MinTeamSize       int           `db:"min_team_size"`
	MaxTeamSize       int           `db:"max_team_size"`
	CreatedAt         time.Time     `db:"created_at"`
	UpdatedAt         time.Time     `db:"updated_at"`
}
