package discord

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/messagepicker"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/pubsub"
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

func Start(token string, eventsChan string) {

	// Connecting to NATS to subscribe to Enshrouded Server Events
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to NATS:", err)
	}

	// Create a new Discord session
	logger.Println("🤖 Creating Discord session...")
	dg, err := discordgo.New("Bot " + token)
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
		log.Fatalf("❌ Error loading messages: %v", err)
	}

	// Start the Events subscription in a goroutine
	go pubsub.StartNATSListener(nc, dg, eventsChan, manager)
	logger.Println("NATS Event subscription routine started")

	logger.Println("✅ Bot is now running. Press CTRL+C to exit.")

	// Wait here until CTRL-C or other signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Println("Shutting down.")
	defer nc.Close()
	dg.Close()
}
