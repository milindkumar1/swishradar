package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	apiURL string
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	apiURL = os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}

	// Create Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	// Register commands
	dg.AddHandler(messageCreate)
	dg.AddHandler(interactionCreate)

	// Open connection
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	// Register slash commands
	registerCommands(dg)

	// Setup cron jobs for scheduled reports
	setupCronJobs(dg)

	fmt.Println("🏀 SwishRadar Discord Bot is now running!")
	fmt.Println("Press CTRL-C to exit")

	// Wait for interrupt signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanup
	dg.Close()
}

func registerCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "matchup",
			Description: "Get current week's matchup prediction",
		},
		{
			Name:        "streaming",
			Description: "Get top waiver wire streaming recommendations",
		},
		{
			Name:        "powerrankings",
			Description: "Get current power rankings for your league",
		},
		{
			Name:        "player",
			Description: "Get player statistics and trends",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Player name",
					Required:    true,
				},
			},
		},
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Simple ping command
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "🏀 Pong! SwishRadar is online.")
	}
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "matchup":
		handleMatchupCommand(s, i)
	case "streaming":
		handleStreamingCommand(s, i)
	case "powerrankings":
		handlePowerRankingsCommand(s, i)
	case "player":
		handlePlayerCommand(s, i)
	}
}

func handleMatchupCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "📊 **This Week's Matchup**\n\n🔜 Coming soon! Matchup predictions will be available here.",
		},
	})
}

func handleStreamingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Fetch from API
	resp, err := http.Get(apiURL + "/api/v1/analytics/streaming")
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Error fetching streaming recommendations",
			},
		})
		return
	}
	defer resp.Body.Close()

	var recommendations []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&recommendations)

	content := "🔥 **Top Waiver Wire Pickups**\n\n"
	if len(recommendations) == 0 {
		content += "No recommendations available yet. Check back soon!"
	} else {
		// Format recommendations
		content += "Coming soon with real data!"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func handlePowerRankingsCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🏆 **League Power Rankings**\n\n🔜 Coming soon! Power rankings will be calculated here.",
		},
	})
}

func handlePlayerCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	playerName := options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("📈 **Player Stats: %s**\n\n🔜 Coming soon!", playerName),
		},
	})
}

func setupCronJobs(s *discordgo.Session) {
	c := cron.New()

	// Daily morning report at 9 AM
	c.AddFunc("0 9 * * *", func() {
		sendDailyReport(s)
	})

	c.Start()
}

func sendDailyReport(s *discordgo.Session) {
	// This would send to configured channels
	// For now, just log it
	log.Println("📅 Daily report scheduled - implementation coming soon")
}
