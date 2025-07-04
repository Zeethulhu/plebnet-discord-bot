package discord

import (
	"os"

	"github.com/joho/godotenv"
)

func StartServer() {
	// Load envrironment variables from .env file

	logger.Println("📦 Loading environment...")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("❌ Error loading .env file")
	}

	logger.Println("🔐 Reading token...")
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		logger.Fatal("DISCORD_TOKEN is not set")
	}

	eventsChannel := os.Getenv("DISCORD_EVENTS_CHANNEL")
	if eventsChannel == "" {
		logger.Fatal("DISCORD_EVENTS_CHANNEL is not set")
	}

	logger.Println("🚀 Starting bot...")
	Start(token, eventsChannel)
}
