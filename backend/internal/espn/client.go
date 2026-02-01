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
		Name                string          `json:"name"`
		ScoringSettings     json.RawMessage `json:"scoringSettings"`
		RosterSettings      json.RawMessage `json:"rosterSettings"`
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
	ID               int    `json:"id"`
	FullName         string `json:"fullName"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	DefaultPositionId int   `json:"defaultPositionId"`
	ProTeamId        int    `json:"proTeamId"`
	Injured          bool   `json:"injured"`
	InjuryStatus     string `json:"injuryStatus"`
}

type Member struct {
	ID              string `json:"id"`
	DisplayName     string `json:"displayName"`
	IsLeagueManager bool   `json:"isLeagueManager"`
}

// GetLeague fetches league information from ESPN
func (c *Client) GetLeague() (*League, error) {
	url := fmt.Sprintf(
		"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s?view=mTeam&view=mRoster&view=mSettings",
		c.Season,
		c.LeagueID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add ESPN authentication cookies
	req.AddCookie(&http.Cookie{Name: "SWID", Value: c.SWID})
	req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.S2})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch league: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ESPN API returned status %d: %s", resp.StatusCode, string(body))
	}

	var league League
	if err := json.NewDecoder(resp.Body).Decode(&league); err != nil {
		return nil, fmt.Errorf("failed to decode league data: %w", err)
	}

	return &league, nil
}

// GetFreeAgents fetches available free agents
func (c *Client) GetFreeAgents(limit int) ([]Player, error) {
	url := fmt.Sprintf(
		"https://fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%s?view=kona_player_info",
		c.Season,
		c.LeagueID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{Name: "SWID", Value: c.SWID})
	req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.S2})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch free agents: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ESPN API returned status %d", resp.StatusCode)
	}

	var data struct {
		Players []struct {
			Player Player `json:"player"`
		} `json:"players"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode free agents: %w", err)
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
