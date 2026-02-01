package models

import "time"

// League represents an ESPN Fantasy Basketball league
type League struct {
	ID           string    `json:"id" db:"id"`
	Season       int       `json:"season" db:"season"`
	Name         string    `json:"name" db:"name"`
	SettingsJSON string    `json:"settings_json" db:"settings_json"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Team represents a fantasy team in a league
type Team struct {
	ID         int       `json:"id" db:"id"`
	LeagueID   string    `json:"league_id" db:"league_id"`
	OwnerID    *int      `json:"owner_id" db:"owner_id"`
	TeamName   string    `json:"team_name" db:"team_name"`
	RosterJSON string    `json:"roster_json" db:"roster_json"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// User represents a platform user
type User struct {
	ID          int       `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	DisplayName string    `json:"display_name" db:"display_name"`
	ESPNSWID    *string   `json:"-" db:"espn_swid"` // Never expose in JSON
	ESPNS2      *string   `json:"-" db:"espn_s2"`   // Never expose in JSON
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
