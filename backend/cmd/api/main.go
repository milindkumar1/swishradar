package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://*"},
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

	// API v1 routes
	r.Route("/api/v1", func(r chi.Router) {
		// ESPN sync routes
		r.Route("/espn", func(r chi.Router) {
			r.Get("/league", handleGetLeague)
			r.Get("/teams", handleGetTeams)
			r.Get("/waiver-wire", handleGetWaiverWire)
			r.Post("/sync", handleSyncLeague)
		})

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
		port = "8080"
	}

	fmt.Printf("🚀 SwishRadar API starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
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
