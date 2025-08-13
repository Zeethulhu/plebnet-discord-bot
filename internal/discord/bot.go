package discord

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/config"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/games"
	_ "github.com/Zeethulhu/plebnet-discord-bot/internal/games/enshrouded"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/messagepicker"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/subscribers"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/timers"
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

// Start launches the bot using the provided configuration.
func Start(cfg config.Config) {
	ctx, cancel := context.WithCancel(context.Background())

	// Connecting to NATS to subscribe to Server Events
	nc, err := nats.Connect(cfg.NatsAddress)
	if err != nil {
		logger.Fatal("❌ Failed to connect to NATS:", err)
	}

	// Create a new Discord session
	logger.Println("🤖 Creating Discord session...")
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		logger.Fatalf("❌ Error creating Discord session: %v", err)
	}

	// Register the message handler
	dg.AddHandler(MessageCreate)

	// Open the websocket connection to Discord
	logger.Println("📡 Connecting to Discord...")
	err = dg.Open()
	if err != nil {
		logger.Fatalf("❌ Cannot open the session: %v", err)
	}

	manager, err := messagepicker.NewManager("internal/config/messages.yaml", 3)
	if err != nil {
		logger.Fatalf("❌ Error loading messages: %v", err)
	}

	timersStarted := false

	// Register event handlers and start subscriptions for each configured game
	for _, g := range cfg.Games {
		channel := g.DiscordChannel
		if channel == "" {
			channel = cfg.EventsChannel
		}
		if channel == "" {
			logger.Printf("⚠️ Skipping game '%s': no Discord channel configured", g.Name)
			continue
		}

		started := false

		if g.NatsTopic != "" {
			if handler, ok := games.NewNATSHandler(g.Name, channel, g.NatsTopic, manager); ok {
				subscribers.Register(handler)
				logger.Printf("📡 NATS handler started for game '%s' on topic '%s'", g.Name, g.NatsTopic)
				started = true
			} else {
				logger.Printf("⚠️ No NATS handler registered for game '%s'", g.Name)
			}
		} else {
			logger.Printf("⚠️ Game '%s' missing NATS topic; NATS handler not started", g.Name)
		}

		if g.SteamRSS != "" {
			if _, err := timers.NewSteamNewsTimer(channel, g.SteamRSS, "steam_news.db"); err != nil {
				logger.Printf("❌ Failed to start Steam news timer for '%s': %v", g.Name, err)
			} else {
				timersStarted = true
				started = true
				logger.Printf("⏰ Steam news timer started for game '%s'", g.Name)
			}
		} else {
			logger.Printf("ℹ️ Game '%s' missing Steam RSS; timer not started", g.Name)
		}

		if started {
			logger.Printf("🎮 Game '%s' started", g.Name)
		} else {
			logger.Printf("⚠️ Game '%s' skipped: no handlers started", g.Name)
		}
	}

	go subscribers.StartListeners(nc, dg)
	logger.Println("NATS Event subscription routine started")

	if timersStarted {
		go timers.Start(ctx, dg)
	}

	logger.Println("✅ Bot is now running. Press CTRL+C to exit.")

	// Wait here until CTRL-C or other signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()

	logger.Println("Shutting down.")
	defer nc.Close()
	if err := dg.Close(); err != nil {
		logger.Printf("error closing Discord session: %v", err)
	}
}
