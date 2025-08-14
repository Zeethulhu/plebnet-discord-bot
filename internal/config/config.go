package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Default values for optional configuration variables.
const (
	DefaultNatsAddress = "nats://127.0.0.1:4222"
	DefaultNatsTopic   = "enshrouded-logs"
)

// Options specifies sources of configuration that override defaults.
type Options struct {
	ConfigFile    string
	DiscordToken  string
	EventsChannel string
	NatsAddress   string
	NatsTopic     string
	Games         []GameConfig
}

// Config holds the merged configuration values.
type Config struct {
	DiscordToken  string
	EventsChannel string
	NatsAddress   string
	// NatsTopic is the default NATS subject used when a game does not
	// specify its own nats_topic.
	NatsTopic string
	Games     []GameConfig
}

// GameConfig holds configuration for a single game.
type GameConfig struct {
	Name           string `yaml:"name"`
	DiscordChannel string `yaml:"discord_channel"`
	// NatsTopic overrides the root NatsTopic for this specific game.
	NatsTopic string `yaml:"nats_topic"`
	SteamRSS  string `yaml:"steam_rss"`
}

var (
	cfg     Config
	once    sync.Once
	loadErr error
)

// Load reads the environment once and returns the configuration. If required
// variables are missing it returns an error.
func Load(opts Options) (Config, error) {
	once.Do(func() {
		_ = godotenv.Load()

		// defaults
		cfg.NatsAddress = DefaultNatsAddress
		cfg.NatsTopic = DefaultNatsTopic

		// config file
		if opts.ConfigFile != "" {
			data, err := os.ReadFile(opts.ConfigFile)
			if err != nil {
				loadErr = fmt.Errorf("error reading config file: %w", err)
				return
			}
			var fc struct {
				DiscordToken  string       `yaml:"discord_token"`
				EventsChannel string       `yaml:"events_channel"`
				NatsAddress   string       `yaml:"nats_address"`
				NatsTopic     string       `yaml:"nats_topic"`
				Games         []GameConfig `yaml:"games"`
			}
			if err := yaml.Unmarshal(data, &fc); err != nil {
				loadErr = fmt.Errorf("error parsing config file: %w", err)
				return
			}
			if fc.DiscordToken != "" {
				cfg.DiscordToken = fc.DiscordToken
			}
			if fc.EventsChannel != "" {
				cfg.EventsChannel = fc.EventsChannel
			}
			if fc.NatsAddress != "" {
				cfg.NatsAddress = fc.NatsAddress
			}
			if fc.NatsTopic != "" {
				cfg.NatsTopic = fc.NatsTopic
			}
			if len(fc.Games) > 0 {
				cfg.Games = fc.Games
			}
		}

		// environment variables override file
		if v := strings.TrimSpace(os.Getenv("DISCORD_TOKEN")); v != "" {
			cfg.DiscordToken = v
		}
		if v := strings.TrimSpace(os.Getenv("DISCORD_EVENTS_CHANNEL")); v != "" {
			cfg.EventsChannel = v
		}
		if v := strings.TrimSpace(os.Getenv("DISCORD_NATS_ADDRESS")); v != "" {
			cfg.NatsAddress = v
		}
		if v := strings.TrimSpace(os.Getenv("DISCORD_NATS_TOPIC")); v != "" {
			cfg.NatsTopic = v
		}

		// command line flags override env
		if opts.DiscordToken != "" {
			cfg.DiscordToken = opts.DiscordToken
		}
		if opts.EventsChannel != "" {
			cfg.EventsChannel = opts.EventsChannel
		}
		if opts.NatsAddress != "" {
			cfg.NatsAddress = opts.NatsAddress
		}
		if opts.NatsTopic != "" {
			cfg.NatsTopic = opts.NatsTopic
		}
		if len(opts.Games) > 0 {
			cfg.Games = opts.Games
		}

		// ensure defaults for optional values
		if cfg.NatsAddress == "" {
			cfg.NatsAddress = DefaultNatsAddress
		}
		if cfg.NatsTopic == "" {
			cfg.NatsTopic = DefaultNatsTopic
		}

		if cfg.DiscordToken == "" {
			loadErr = fmt.Errorf("discord_token is not set")
			return
		}
	})

	return cfg, loadErr
}
