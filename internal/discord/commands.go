package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandMap hold all available commands
var CommandMap = map[string]func([]string, *discordgo.Session, *discordgo.MessageCreate){
	"ping": pingCommand,
	"echo": echoCommand,
	"help": helpCommand,
}

func pingCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "🏓 Pong!")
	if err != nil {
		// Handle the error, e.g. log it
		logger.Printf("❌ Failed to send message: %v", err)
		return
	}
	logger.Printf("✅ Message sent: %s", msg.ID)
}

func echoCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(args) == 0 {
		_, err := s.ChannelMessageSend(m.ChannelID, "🔇 Nothing to echo.")
		if err != nil {
			// Handle the error, e.g. log it
			logger.Printf("❌ Failed to send message: %v", err)
			return
		}
		return
	}
	echo_msg := strings.Join(args, " ")
	msg, err := s.ChannelMessageSend(m.ChannelID, echo_msg)
	if err != nil {
		// Handle the error, e.g. log it
		logger.Printf("❌ Failed to send message: %v", err)
		return
	}
	logger.Printf("✅ Message sent: %s", msg.ID)
}

func helpCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "**Available commands:** \n!ping\n!echo [text]\n!help")
	if err != nil {
		// Handle the error, e.g. log it
		logger.Printf("❌ Failed to send message: %v", err)
		return
	}
	logger.Printf("✅ Message sent: %s", msg.ID)
}
