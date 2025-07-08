package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Echo struct{}

func (e *Echo) Name() string { return "echo" }

func (e *Echo) Description() string { return "Echo back provided text" }

func (e *Echo) Execute(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(args) == 0 {
		if _, err := s.ChannelMessageSend(m.ChannelID, "🔇 Nothing to echo."); err != nil {
			logger.Printf("❌ Failed to send message: %v", err)
		}
		return
	}
	echoMsg := strings.Join(args, " ")
	msg, err := s.ChannelMessageSend(m.ChannelID, echoMsg)
	if err != nil {
		logger.Printf("❌ Failed to send message: %v", err)
		return
	}
	logger.Printf("✅ Message sent: %s", msg.ID)
}

func init() {
	Register(&Echo{})
}
