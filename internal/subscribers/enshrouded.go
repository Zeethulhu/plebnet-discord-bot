package subscribers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/messagepicker"
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

// ServerEvent represents an Enshrouded login/logout event.
type ServerEvent struct {
	LogOn     bool   `json:"log_on"`
	LogOff    bool   `json:"log_off"`
	Player    string `json:"player"`
	Timestamp string `json:"timestamp"`
}

// EnshroudedLoginHandler handles login/logout events published to NATS.
type EnshroudedLoginHandler struct {
	ChannelID   string
	SubjectName string
	Manager     *messagepicker.MessageManager
}

// NewEnshroudedLoginHandler creates the handler and registers it.
func NewEnshroudedLoginHandler(channelID, subject string, manager *messagepicker.MessageManager) *EnshroudedLoginHandler {
	h := &EnshroudedLoginHandler{ChannelID: channelID, SubjectName: subject, Manager: manager}
	Register(h)
	return h
}

func (h *EnshroudedLoginHandler) Subject() string { return h.SubjectName }

func (h *EnshroudedLoginHandler) Handle(msg *nats.Msg, discord *discordgo.Session) {
	var event ServerEvent
	logger.Println("Received Enshrouded event")
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		logger.Printf("❌ Failed to parse event: %v", err)
		return
	}

	t, err := time.Parse(time.RFC3339Nano, event.Timestamp)
	if err != nil {
		logger.Printf("❌ Failed to parse timestamp: %v", err)
		t = time.Now()
	}
	tStr := t.Local().Format("Jan 2 15:04:05")

	var msgStr string
	if event.LogOn {
		msgStr, err = h.Manager.Pick("join", event.Player)
		if err != nil {
			logger.Println("Message error:", err)
			return
		}
		_, err = discord.ChannelMessageSend(h.ChannelID, fmt.Sprintf("Player joined. %s", msgStr))
		if err != nil {
			logger.Printf("❌ Failed to send message: %v", err)
		}
	} else if event.LogOff {
		msgStr, err = h.Manager.Pick("leave", event.Player)
		if err != nil {
			logger.Println("Message error:", err)
			return
		}
		_, err = discord.ChannelMessageSend(h.ChannelID, fmt.Sprintf("Player left. %s", msgStr))
		if err != nil {
			logger.Printf("❌ Failed to send message: %v", err)
		}
	} else {
		_, err := discord.ChannelMessageSend(h.ChannelID, fmt.Sprintf("⚠️ Unrecognized player event for @%s at %s", event.Player, tStr))
		if err != nil {
			logger.Printf("❌ Failed to send message: %v", err)
		}
	}
}
