package subscribers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

// Handler represents a subscription handler for a NATS subject.
type Handler interface {
	// Subject returns the NATS subject this handler subscribes to.
	Subject() string
	// Handle processes a NATS message.
	Handle(msg *nats.Msg, s *discordgo.Session)
}

var registry []Handler

// Register adds a handler to the registry.
func Register(h Handler) {
	registry = append(registry, h)
}

// All returns all registered handlers.
func All() []Handler {
	return registry
}
