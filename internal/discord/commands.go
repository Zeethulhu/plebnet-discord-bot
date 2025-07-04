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
	s.ChannelMessageSend(m.ChannelID, "🏓 Pong!")
}

func echoCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "🔇 Nothing to echo.")
		return
	}
	msg := strings.Join(args, " ")
	s.ChannelMessageSend(m.ChannelID, msg)
}

func helpCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "**Available commands:** \n!ping\n!echo [text]\n!help")
}
