package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles ESPN Fantasy API requests
type Client struct {
	SWID     string
	S2       string
	LeagueID string
	Season   int
	client   *http.Client
}

// NewClient creates a new ESPN API client
func NewClient(swid, s2, leagueID string, season int) *Client {
	return &Client{
		SWID:     swid,
		S2:       s2,
		LeagueID: leagueID,
		Season:   season,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// LeagueResponse represents the ESPN league response
type LeagueResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Season   int    `json:"seasonId"`
	Settings struct {
		ScoringSettings struct {
			ScoringItems []struct {
				StatID        int     `json:"statId"`
				Points        float64 `json:"points"`
				PointsOverride float64 `json:"pointsOverride"`
			} `json:"scoringItems"`
		} `json:"scoringSettings"`
	} `json:"settings"`
	Teams []struct {
		ID       int    `json:"id"`
		Location string `json:"location"`
		Nickname string `json:"nickname"`
		Roster   struct {
			Entries []struct {
				PlayerPoolEntry struct {
					Player struct {
						ID        int    `json:"id"`
						FullName  string `json:"fullName"`
						ProTeamID int    `json:"proTeamId"`
					} `json:"player"`
				} `json:"playerPoolEntry"`
			} `json:"entries"`
		} `json:"roster"`
	} `json:"teams"`
}

// GetLeague fetches league information from ESPN
func (c *Client) GetLeague() (*LeagueResponse, error) {
	url := fmt.Sprintf(
		"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s",
		c.Season,
		c.LeagueID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add ESPN authentication cookies
	req.AddCookie(&http.Cookie{Name: "swid", Value: c.SWID})
	req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.S2})

	// Add query parameters for more data
	q := req.URL.Query()
	q.Add("view", "mTeam")
	q.Add("view", "mRoster")
	q.Add("view", "mMatchup")
	q.Add("view", "mSettings")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch league: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ESPN API error: %d - %s", resp.StatusCode, string(body))
	}

	var leagueResp LeagueResponse
	if err := json.NewDecoder(resp.Body).Decode(&leagueResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &leagueResp, nil
}

// GetFreeAgents fetches available free agents
func (c *Client) GetFreeAgents(limit int) ([]interface{}, error) {
	url := fmt.Sprintf(
		"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s",
		c.Season,
		c.LeagueID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{Name: "swid", Value: c.SWID})
	req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.S2})

	q := req.URL.Query()
	q.Add("view", "kona_player_info")
	q.Add("scoringPeriodId", "0")
	if limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", limit))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch free agents: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ESPN API error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Extract free agents from response
	// This is a placeholder - actual structure may vary
	return []interface{}{}, nil
}
