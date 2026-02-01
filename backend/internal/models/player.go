package models

import "time"

// Player represents an NBA player
type Player struct {
	ID        int       `json:"id" db:"id"`
	ESPNID    *int      `json:"espn_id" db:"espn_id"`
	Name      string    `json:"name" db:"name"`
	Position  string    `json:"position" db:"position"`
	Team      string    `json:"team" db:"team"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// PlayerStats represents daily player statistics
type PlayerStats struct {
	ID           int       `json:"id" db:"id"`
	PlayerID     int       `json:"player_id" db:"player_id"`
	Date         time.Time `json:"date" db:"date"`
	Points       float64   `json:"points" db:"points"`
	Rebounds     float64   `json:"rebounds" db:"rebounds"`
	Assists      float64   `json:"assists" db:"assists"`
	Steals       float64   `json:"steals" db:"steals"`
	Blocks       float64   `json:"blocks" db:"blocks"`
	Turnovers    float64   `json:"turnovers" db:"turnovers"`
	ThreesMade   float64   `json:"threes_made" db:"threes_made"`
	FGM          float64   `json:"fgm" db:"fgm"`
	FGA          float64   `json:"fga" db:"fga"`
	FTM          float64   `json:"ftm" db:"ftm"`
	FTA          float64   `json:"fta" db:"fta"`
	Minutes      float64   `json:"minutes" db:"minutes"`
	FantasyValue float64   `json:"fantasy_value" db:"fantasy_value"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// StreamingRecommendation represents a waiver wire pickup suggestion
type StreamingRecommendation struct {
	Player            Player  `json:"player"`
	Score             float64 `json:"score"`
	GamesThisWeek     int     `json:"games_this_week"`
	ProjectedValue    float64 `json:"projected_value"`
	TrendDelta        float64 `json:"trend_delta"`
	MinutesStability  float64 `json:"minutes_stability"`
	OpportunityFactor float64 `json:"opportunity_factor"`
	Reason            string  `json:"reason"`
}
