package discord

import (
	"strings"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/discord/commands"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.Bot {
		return
	}

	prefix := "!"
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// Parse command and arguments
	content := strings.TrimPrefix(m.Content, prefix)
	parts := utils.ParseArgs(content)
	if len(parts) == 0 {
		return
	}
	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	// Route command using registry
	if handler, ok := commands.Get(cmd); ok {
		handler.Execute(args, s, m)
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "❓ Unknown command. Try `!help`.")
		if err != nil {
			logger.Printf("❌ Failed to send message: %v", err)
			return
		}
	}
}
