package timers

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Task represents a scheduled task to run periodically.
type Task interface {
	Name() string
	Interval() time.Duration
	Run(ctx context.Context, s *discordgo.Session)
}

var registry []Task

// Register adds a task to the registry.
func Register(t Task) { registry = append(registry, t) }

// All returns all registered tasks.
func All() []Task { return registry }
