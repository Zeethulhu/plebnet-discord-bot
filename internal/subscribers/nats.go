package subscribers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

// StartListeners subscribes all registered handlers to their subjects.
func StartListeners(nc *nats.Conn, discord *discordgo.Session) {
	for _, h := range All() {
		handler := h
		_, err := nc.Subscribe(handler.Subject(), func(msg *nats.Msg) {
			handler.Handle(msg, discord)
		})
		if err != nil {
			logger.Fatalf("❌ Failed to subscribe to NATS subject: %v", err)
		}
		logger.Printf("📡 NATS listener running on subject '%s'", handler.Subject())
	}
}
