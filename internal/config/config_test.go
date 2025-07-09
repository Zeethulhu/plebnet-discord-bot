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
	if err := os.Setenv("DISCORD_EVENTS_CHANNEL", "channel"); err != nil {
		t.Fatalf("failed to set DISCORD_EVENTS_CHANNEL: %v", err)
	}

	reset()
	c, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.NatsAddress != DefaultNatsAddress {
		t.Errorf("expected NatsAddress %q, got %q", DefaultNatsAddress, c.NatsAddress)
	}
	if c.NatsTopic != DefaultNatsTopic {
		t.Errorf("expected NatsTopic %q, got %q", DefaultNatsTopic, c.NatsTopic)
	}
}

func TestMissingVariables(t *testing.T) {
	os.Clearenv()
	createEnvFile(t, "")

	reset()
	_, err := Load()
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
