package games

import (
	"strings"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/messagepicker"
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

// GameNATSHandler defines the behavior of a game-specific NATS message handler.
type GameNATSHandler interface {
	// Subject returns the NATS subject this handler subscribes to.
	Subject() string
	// Handle processes a NATS message using the provided Discord session.
	Handle(msg *nats.Msg, s *discordgo.Session)
}

// HandlerFactory creates a NATS handler for a game.
type HandlerFactory func(channelID, subject string, manager *messagepicker.Manager) GameNATSHandler

var natsHandlers = map[string]HandlerFactory{}

// RegisterNATSHandler registers a factory for the given game name.
func RegisterNATSHandler(name string, f HandlerFactory) {
	natsHandlers[strings.ToLower(name)] = f
}

// NewNATSHandler instantiates a handler for the specified game.
// The boolean return indicates whether a factory was registered for that game.
func NewNATSHandler(name, channelID, subject string, manager *messagepicker.Manager) (GameNATSHandler, bool) {
	f, ok := natsHandlers[strings.ToLower(name)]
	if !ok {
		return nil, false
	}
	return f(channelID, subject, manager), true
}
