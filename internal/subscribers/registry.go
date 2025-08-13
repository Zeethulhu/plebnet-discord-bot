package subscribers

import "github.com/Zeethulhu/plebnet-discord-bot/internal/games"

var registry []games.GameNATSHandler

// Register adds a handler to the registry.
func Register(h games.GameNATSHandler) {
	registry = append(registry, h)
}

// All returns all registered handlers.
func All() []games.GameNATSHandler {
	return registry
}
