package timers

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Start launches all registered tasks.
func Start(ctx context.Context, s *discordgo.Session) {
	for _, t := range All() {
		task := t
		go func() {
			ticker := time.NewTicker(task.Interval())
			defer ticker.Stop()
			logger.Printf("⏰ Timer '%s' started, interval %s", task.Name(), task.Interval())
			for {
				select {
				case <-ctx.Done():
					logger.Printf("🛑 Timer '%s' stopped", task.Name())
					if err := task.Close(); err != nil {
						logger.Printf("❌ Error closing timer '%s': %v", task.Name(), err)
					}
					return
				case <-ticker.C:
					task.Run(ctx, s)
				}
			}
		}()
	}
}
