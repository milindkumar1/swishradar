package nba

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL   = "https://stats.nba.com/stats"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
)

// Client handles NBA Stats API requests
type Client struct {
	client *http.Client
}

// NewClient creates a new NBA Stats API client
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// PlayerGameLog represents a single game performance
type PlayerGameLog struct {
	PlayerID      int     `json:"Player_ID"`
	GameID        string  `json:"Game_ID"`
	GameDate      string  `json:"GAME_DATE"`
	Matchup       string  `json:"MATCHUP"`
	WL            string  `json:"WL"`
	Min           float64 `json:"MIN"`
	FGM           float64 `json:"FGM"`
	FGA           float64 `json:"FGA"`
	FG3M          float64 `json:"FG3M"`
	FG3A          float64 `json:"FG3A"`
	FTM           float64 `json:"FTM"`
	FTA           float64 `json:"FTA"`
	OREB          float64 `json:"OREB"`
	DREB          float64 `json:"DREB"`
	REB           float64 `json:"REB"`
	AST           float64 `json:"AST"`
	STL           float64 `json:"STL"`
	BLK           float64 `json:"BLK"`
	TOV           float64 `json:"TOV"`
	PF            float64 `json:"PF"`
	PTS           float64 `json:"PTS"`
	PlusMinus     float64 `json:"PLUS_MINUS"`
	FantasyPoints float64 `json:"FANTASY_PTS"`
}

// GetPlayerGameLog fetches game logs for a specific player
func (c *Client) GetPlayerGameLog(playerID string, season string) ([]PlayerGameLog, error) {
	url := fmt.Sprintf("%s/playergamelog", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// NBA Stats API requires specific headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://stats.nba.com/")

	q := req.URL.Query()
	q.Add("PlayerID", playerID)
	q.Add("Season", season)
	q.Add("SeasonType", "Regular Season")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch game log: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("NBA API error: %d - %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Parse the response structure (NBA API uses specific format)
	// This is a placeholder for the actual parsing logic
	return []PlayerGameLog{}, nil
}

// GetTeamSchedule fetches the schedule for a team
func (c *Client) GetTeamSchedule(teamID string, season string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/teamgamelog", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://stats.nba.com/")

	q := req.URL.Query()
	q.Add("TeamID", teamID)
	q.Add("Season", season)
	q.Add("SeasonType", "Regular Season")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedule: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NBA API error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return []interface{}{}, nil
}

// GetAllPlayers fetches the list of all active NBA players
func (c *Client) GetAllPlayers(season string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/commonallplayers", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://stats.nba.com/")

	q := req.URL.Query()
	q.Add("LeagueID", "00")
	q.Add("Season", season)
	q.Add("IsOnlyCurrentSeason", "1")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch players: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NBA API error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return []interface{}{}, nil
}
