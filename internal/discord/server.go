package discord

import (
	"os"
	"strings"

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

	natsAddr := os.Getenv("DISCORD_NATS_ADDRESS")
	if natsAddr == "" {
		logger.Fatal("DISCORD_NATS_ADDRESS is not set")
	}

	natsTopic := os.Getenv("DISCORD_NATS_TOPIC")
	if natsTopic == "" {
		logger.Fatal("DISCORD_NATS_TOPIC is not set")
	}

	natsAddr = strings.TrimSpace(natsAddr)
	natsTopic = strings.TrimSpace(natsTopic)

	logger.Println("🚀 Starting bot...")
	Start(token, eventsChannel, natsAddr, natsTopic)
}
