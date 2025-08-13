package config

import (
	"os"
	"sync"
	"testing"
)

func reset() {
	once = sync.Once{}
	cfg = Config{}
	loadErr = nil
}

func createEnvFile(t *testing.T, content string) {
	if err := os.WriteFile(".env", []byte(content), 0644); err != nil {
		t.Fatalf("failed to create env file: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(".env") })

}

func TestDefaultValues(t *testing.T) {
	os.Clearenv()
	createEnvFile(t, "")
	if err := os.Setenv("DISCORD_TOKEN", "token"); err != nil {
		t.Fatalf("failed to set DISCORD_TOKEN: %v", err)
	}

	reset()
	c, err := Load(Options{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.NatsAddress != DefaultNatsAddress {
		t.Errorf("expected NatsAddress %q, got %q", DefaultNatsAddress, c.NatsAddress)
	}
	if c.NatsTopic != DefaultNatsTopic {
		t.Errorf("expected NatsTopic %q, got %q", DefaultNatsTopic, c.NatsTopic)
	}
	if c.EventsChannel != "" {
		t.Errorf("expected EventsChannel to be empty, got %q", c.EventsChannel)
	}
}

func TestMissingVariables(t *testing.T) {
	os.Clearenv()
	createEnvFile(t, "")

	reset()
	_, err := Load(Options{})
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestConfigFile(t *testing.T) {
	os.Clearenv()
	createEnvFile(t, "")

	content := []byte("discord_token: t\nnats_address: a\nnats_topic: top")
	if err := os.WriteFile("config.yaml", content, 0644); err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove("config.yaml") })

	reset()
	cfg, err := Load(Options{ConfigFile: "config.yaml"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.DiscordToken != "t" || cfg.NatsAddress != "a" || cfg.NatsTopic != "top" || cfg.EventsChannel != "" {
		t.Fatalf("config not loaded from file: %+v", cfg)
	}
}

func TestGamesFromConfigFile(t *testing.T) {
	os.Clearenv()
	createEnvFile(t, "")

	content := []byte("discord_token: t\ngames:\n  - name: g1\n    discord_channel: dc\n    nats_topic: nt\n    steam_rss: sr")
	if err := os.WriteFile("config.yaml", content, 0644); err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove("config.yaml") })

	reset()
	c, err := Load(Options{ConfigFile: "config.yaml"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.EventsChannel != "" {
		t.Fatalf("expected EventsChannel to be empty, got %q", c.EventsChannel)
	}
	if len(c.Games) != 1 {
		t.Fatalf("expected 1 game, got %d", len(c.Games))
	}
	g := c.Games[0]
	if g.Name != "g1" || g.DiscordChannel != "dc" || g.NatsTopic != "nt" || g.SteamRSS != "sr" {
		t.Fatalf("unexpected game config: %+v", g)
	}
}
