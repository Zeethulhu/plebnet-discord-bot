package discord

import (
    "github.com/Zeethulhu/plebnet-discord-bot/internal/config"
)

func StartServer(opts config.Options) {
        logger.Println("📦 Loading configuration...")
        cfg, err := config.Load(opts)
        if err != nil {
                logger.Fatal(err)
        }

	logger.Println("🚀 Starting bot...")
	Start(cfg.DiscordToken, cfg.EventsChannel, cfg.NatsAddress, cfg.NatsTopic)
}
