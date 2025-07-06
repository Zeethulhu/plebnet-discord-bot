package pubsub

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/messagepicker"
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

type ServerEvent struct {
	LogOn     bool   `json:"log_on"`
	LogOff    bool   `json:"log_off"`
	Player    string `json:"player"`
	Timestamp string `json:"timestamp"`
}

func StartNATSListener(nc *nats.Conn, discord *discordgo.Session, channelID, subject string, manager *messagepicker.MessageManager) {
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		handleServerEvent(msg, discord, channelID, manager)
	})

	if err != nil {
		log.Fatalf("❌ Failed to subscribe to NATS subject: %v", err)
	}

	log.Printf("📡 NATS listener running on subject '%s'", subject)
}

func handleServerEvent(msg *nats.Msg, discord *discordgo.Session, channelID string, manager *messagepicker.MessageManager) {
	var event ServerEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("❌ Failed to parse event: %v", err)
		return
	}

	t, err := time.Parse(time.RFC3339Nano, event.Timestamp)
	if err != nil {
		log.Printf("❌ Failed to parse timestamp: %v", err)
		t = time.Now()
	}
	tStr := t.Local().Format("Jan 2 15:04:05")

	var msgStr string

	if event.LogOn {
		msgStr, err = manager.Pick("join", event.Player)
		if err != nil {
			log.Println("Message error:", err)
			return
		}
		_, err = discord.ChannelMessageSend(channelID, fmt.Sprintf("Player joined. %s", msgStr))
		if err != nil {
			// Handle the error, e.g. log it
			logger.Printf("❌ Failed to send message: %v", err)
			return
		}
	} else if event.LogOff {
		msgStr, err = manager.Pick("leave", event.Player)
		if err != nil {
			log.Println("Message error:", err)
			return
		}
		_, err = discord.ChannelMessageSend(channelID, fmt.Sprintf("Player left. %s", msgStr))
		if err != nil {
			// Handle the error, e.g. log it
			logger.Printf("❌ Failed to send message: %v", err)
			return
		}
	} else {
		_, err := discord.ChannelMessageSend(channelID, fmt.Sprintf("⚠️  Unrecognized player event for @%s at %s", event.Player, tStr))
		if err != nil {
			// Handle the error, e.g. log it
			logger.Printf("❌ Failed to send message: %v", err)
			return
		}
	}
}
