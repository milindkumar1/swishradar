package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/milindkumar1/swishradar/internal/espn"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🏀 SwishRadar API v1.0"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now(),
		})
	})

	// ESPN API routes
	r.Get("/api/league", getLeagueHandler)
	r.Get("/api/league/teams", getTeamsHandler)
	r.Get("/api/league/free-agents", getFreeAgentsHandler)

	// API v1 routes (future endpoints)
	r.Route("/api/v1", func(r chi.Router) {
		// Analytics routes
		r.Route("/analytics", func(r chi.Router) {
			r.Get("/streaming", handleGetStreamingRecommendations)
			r.Post("/trade", handleCalculateTrade)
			r.Get("/power-rankings", handleGetPowerRankings)
			r.Get("/matchup/{week}", handleGetMatchupPrediction)
		})

		// Player routes
		r.Route("/players", func(r chi.Router) {
			r.Get("/", handleGetPlayers)
			r.Get("/{id}", handleGetPlayer)
			r.Get("/{id}/stats", handleGetPlayerStats)
		})

		// Backtesting routes
		r.Route("/backtest", func(r chi.Router) {
			r.Post("/run", handleRunBacktest)
			r.Get("/results", handleGetBacktestResults)
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("🚀 SwishRadar API starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// ESPN League Handlers
func getLeagueHandler(w http.ResponseWriter, r *http.Request) {
	client := createESPNClient()

	league, err := client.GetLeague()
	if err != nil {
		log.Printf("Error fetching league: %v", err)
		http.Error(w, fmt.Sprintf("Failed to fetch league: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(league)
}

func getTeamsHandler(w http.ResponseWriter, r *http.Request) {
	client := createESPNClient()

	league, err := client.GetLeague()
	if err != nil {
		log.Printf("Error fetching teams: %v", err)
		http.Error(w, fmt.Sprintf("Failed to fetch teams: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"teams": league.Teams,
		"count": len(league.Teams),
	})
}

func getFreeAgentsHandler(w http.ResponseWriter, r *http.Request) {
	client := createESPNClient()

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	players, err := client.GetFreeAgents(limit)
	if err != nil {
		log.Printf("Error fetching free agents: %v", err)
		http.Error(w, fmt.Sprintf("Failed to fetch free agents: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"players": players,
		"count":   len(players),
	})
}

func createESPNClient() *espn.Client {
	leagueID := os.Getenv("ESPN_LEAGUE_ID")
	swid := os.Getenv("ESPN_SWID")
	s2 := os.Getenv("ESPN_S2")
	
	// Get current season (2025 for now)
	season := 2025
	
	return espn.NewClient(leagueID, season, swid, s2)
}

// Placeholder handlers - to be implemented
func handleGetLeague(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "League data endpoint - coming soon"}`))
}

func handleGetTeams(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Teams endpoint - coming soon"}`))
}

func handleGetWaiverWire(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Waiver wire endpoint - coming soon"}`))
}

func handleSyncLeague(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Sync endpoint - coming soon"}`))
}

func handleGetStreamingRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Streaming recommendations - coming soon"}`))
}

func handleCalculateTrade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Trade calculator - coming soon"}`))
}

func handleGetPowerRankings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Power rankings - coming soon"}`))
}

func handleGetMatchupPrediction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Matchup prediction - coming soon"}`))
}

func handleGetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Players list - coming soon"}`))
}

func handleGetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Player details - coming soon"}`))
}

func handleGetPlayerStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Player stats - coming soon"}`))
}

func handleRunBacktest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Backtest runner - coming soon"}`))
}

func handleGetBacktestResults(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Backtest results - coming soon"}`))
}
