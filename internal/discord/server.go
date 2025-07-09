package discord

import (
	"github.com/Zeethulhu/plebnet-discord-bot/internal/config"
)

func StartServer() {
	logger.Println("📦 Loading environment...")
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("🚀 Starting bot...")
	Start(cfg.DiscordToken, cfg.EventsChannel, cfg.NatsAddress, cfg.NatsTopic)
}
