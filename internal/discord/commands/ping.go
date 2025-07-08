package commands

import "github.com/bwmarrin/discordgo"

type Ping struct{}

func (p *Ping) Name() string { return "ping" }

func (p *Ping) Description() string { return "Respond with Pong" }

func (p *Ping) Execute(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "🏓 Pong!")
	if err != nil {
		logger.Printf("\u274C Failed to send message: %v", err)
		return
	}
	logger.Printf("\u2705 Message sent: %s", msg.ID)
}

func init() {
	Register(&Ping{})
}
