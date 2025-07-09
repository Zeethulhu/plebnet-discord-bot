package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

// Default values for optional configuration variables.
const (
	DefaultNatsAddress = "nats://127.0.0.1:4222"
	DefaultNatsTopic   = "enshrouded-logs"
)

// Config holds all configuration loaded from the environment.
type Config struct {
	DiscordToken  string
	EventsChannel string
	NatsAddress   string
	NatsTopic     string
}

var (
	cfg     Config
	once    sync.Once
	loadErr error
)

// Load reads the environment once and returns the configuration. If required
// variables are missing it returns an error.
func Load() (Config, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			loadErr = fmt.Errorf("error loading .env file: %w", err)
			return
		}

		cfg.DiscordToken = strings.TrimSpace(os.Getenv("DISCORD_TOKEN"))
		cfg.EventsChannel = strings.TrimSpace(os.Getenv("DISCORD_EVENTS_CHANNEL"))
		cfg.NatsAddress = strings.TrimSpace(getEnvDefault("DISCORD_NATS_ADDRESS", DefaultNatsAddress))
		cfg.NatsTopic = strings.TrimSpace(getEnvDefault("DISCORD_NATS_TOPIC", DefaultNatsTopic))

		if cfg.DiscordToken == "" {
			loadErr = fmt.Errorf("DISCORD_TOKEN is not set")
			return
		}
		if cfg.EventsChannel == "" {
			loadErr = fmt.Errorf("DISCORD_EVENTS_CHANNEL is not set")
			return
		}
	})

	return cfg, loadErr
}

func getEnvDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
