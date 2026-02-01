package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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

	// ESPN Service URL
	espnServiceURL := os.Getenv("ESPN_SERVICE_URL")
	if espnServiceURL == "" {
		espnServiceURL = "http://localhost:5001"
	}

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("SwishRadar API v1.0"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now(),
		})
	})

	// ESPN Service proxy routes
	r.Get("/api/espn/health", func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, espnServiceURL+"/health")
	})

	r.Get("/api/espn/league", func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, espnServiceURL+"/api/league")
	})

	r.Get("/api/espn/teams", func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, espnServiceURL+"/api/teams")
	})

	r.Get("/api/espn/free-agents", func(w http.ResponseWriter, r *http.Request) {
		url := espnServiceURL + "/api/free-agents"
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}
		proxyRequest(w, r, url)
	})

	r.Get("/api/espn/standings", func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, espnServiceURL+"/api/standings")
	})

	// API v1 routes (future analytics endpoints)
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

	fmt.Printf("SwishRadar API starting on port %s\n", port)
	fmt.Printf("ESPN Service: %s\n", espnServiceURL)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// Proxy helper function
func proxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
	resp, err := http.Get(targetURL)
	if err != nil {
		log.Printf("Error proxying request to %s: %v", targetURL, err)
		http.Error(w, fmt.Sprintf("Failed to connect to ESPN service: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copy status code
	w.WriteHeader(resp.StatusCode)

	// Copy body
	io.Copy(w, resp.Body)
}

// Placeholder handlers for future analytics features
func handleGetStreamingRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Streaming recommendations - coming soon"}`))
}

func handleCalculateTrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Trade calculator - coming soon"}`))
}

func handleGetPowerRankings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Power rankings - coming soon"}`))
}

func handleGetMatchupPrediction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Matchup prediction - coming soon"}`))
}

func handleGetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Players list - coming soon"}`))
}

func handleGetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Player details - coming soon"}`))
}

func handleGetPlayerStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Player stats - coming soon"}`))
}

func handleRunBacktest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Backtest runner - coming soon"}`))
}

func handleGetBacktestResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Backtest results - coming soon"}`))
}
