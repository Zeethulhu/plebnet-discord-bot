package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Help struct{}

func (h *Help) Name() string { return "help" }

func (h *Help) Description() string { return "Show available commands" }

func (h *Help) Execute(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var builder strings.Builder
	builder.WriteString("**Available commands:**\n")
	for _, cmd := range All() {
		builder.WriteString("!" + cmd.Name() + " - " + cmd.Description() + "\n")
	}

	_, err := s.ChannelMessageSend(m.ChannelID, builder.String())

	if err != nil {
		logger.Printf("❌ Failed to send message: %v", err)
		return
	}

	logger.Println("✅ `!help` command invoked.")

}

func init() {
	Register(&Help{})
}
