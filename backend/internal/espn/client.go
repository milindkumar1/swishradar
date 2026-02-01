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
	LeagueID string
	Season   int
	SWID     string
	S2       string
	client   *http.Client
}

// NewClient creates a new ESPN API client
func NewClient(leagueID string, season int, swid, s2 string) *Client {
	return &Client{
		LeagueID: leagueID,
		Season:   season,
		SWID:     swid,
		S2:       s2,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// League represents the ESPN league data
type League struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Season   int    `json:"seasonId"`
	Settings struct {
		Name            string          `json:"name"`
		ScoringSettings json.RawMessage `json:"scoringSettings"`
		RosterSettings  json.RawMessage `json:"rosterSettings"`
	} `json:"settings"`
	Teams   []Team   `json:"teams"`
	Members []Member `json:"members"`
}

type Team struct {
	ID           int    `json:"id"`
	Abbrev       string `json:"abbrev"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Nickname     string `json:"nickname"`
	PrimaryOwner string `json:"primaryOwner"`
	Roster       struct {
		Entries []RosterEntry `json:"entries"`
	} `json:"roster"`
	Record struct {
		Overall struct {
			Wins   int `json:"wins"`
			Losses int `json:"losses"`
			Ties   int `json:"ties"`
		} `json:"overall"`
	} `json:"record"`
}

type RosterEntry struct {
	PlayerPoolEntry struct {
		Player Player `json:"player"`
	} `json:"playerPoolEntry"`
	LineupSlotId int `json:"lineupSlotId"`
}

type Player struct {
	ID                int    `json:"id"`
	FullName          string `json:"fullName"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	DefaultPositionId int    `json:"defaultPositionId"`
	ProTeamId         int    `json:"proTeamId"`
	Injured           bool   `json:"injured"`
	InjuryStatus      string `json:"injuryStatus"`
}

type Member struct {
	ID              string `json:"id"`
	DisplayName     string `json:"displayName"`
	IsLeagueManager bool   `json:"isLeagueManager"`
}

// GetLeague fetches league information from ESPN
func (c *Client) GetLeague() (*League, error) {
	// Try current season first, then fall back to previous season
	seasons := []int{2025, 2024, 2026}

	var lastErr error
	for _, season := range seasons {
		url := fmt.Sprintf(
			"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s?view=mTeam&view=mRoster&view=mSettings&view=mMatchup",
			season,
			c.LeagueID,
		)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Add ESPN authentication cookies - try both cookie names
		// ESPN sometimes uses different cookie names
		req.AddCookie(&http.Cookie{
			Name:  "SWID",
			Value: c.SWID,
		})
		req.AddCookie(&http.Cookie{
			Name:  "swid",
			Value: c.SWID,
		})
		req.AddCookie(&http.Cookie{
			Name:  "espn_s2",
			Value: c.S2,
		})

		// Add headers to mimic browser exactly
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Referer", "https://fantasy.espn.com/basketball/")
		req.Header.Set("Origin", "https://fantasy.espn.com")

		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch league: %w", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			bodyStr := string(body)
			if len(bodyStr) > 300 {
				bodyStr = bodyStr[:300]
			}
			lastErr = fmt.Errorf("ESPN API returned status %d for season %d: %s", resp.StatusCode, season, bodyStr)
			continue
		}

		// Read and check response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		// Check if response is JSON
		if len(body) == 0 {
			lastErr = fmt.Errorf("ESPN returned empty response for season %d", season)
			continue
		}

		if body[0] != '{' && body[0] != '[' {
			bodyStr := string(body)
			if len(bodyStr) > 300 {
				bodyStr = bodyStr[:300]
			}
			lastErr = fmt.Errorf("ESPN returned non-JSON (season %d, first char: %c): %s", season, body[0], bodyStr)
			continue
		}

		var league League
		if err := json.Unmarshal(body, &league); err != nil {
			lastErr = fmt.Errorf("failed to decode league data (season %d): %w", season, err)
			continue
		}

		// Success!
		return &league, nil
	}

	return nil, fmt.Errorf("failed to fetch league from all seasons: %v", lastErr)
}

// GetFreeAgents fetches available free agents
func (c *Client) GetFreeAgents(limit int) ([]Player, error) {
	// Try current season first, then fall back
	seasons := []int{2025, 2024, 2026}

	var lastErr error
	for _, season := range seasons {
		url := fmt.Sprintf(
			"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s?view=kona_player_info",
			season,
			c.LeagueID,
		)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		req.AddCookie(&http.Cookie{Name: "SWID", Value: c.SWID})
		req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.S2})
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		req.Header.Set("Accept", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch free agents: %w", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("ESPN API returned status %d for season %d", resp.StatusCode, season)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		var data struct {
			Players []struct {
				Player Player `json:"player"`
			} `json:"players"`
		}

		if err := json.Unmarshal(body, &data); err != nil {
			lastErr = fmt.Errorf("failed to decode free agents (season %d): %w", season, err)
			continue
		}

		players := make([]Player, 0, len(data.Players))
		for _, p := range data.Players {
			players = append(players, p.Player)
			if len(players) >= limit {
				break
			}
		}

		return players, nil
	}

	return nil, fmt.Errorf("failed to fetch free agents from all seasons: %v", lastErr)
}
