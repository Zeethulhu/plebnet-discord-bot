package pubsub

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

type ServerEvent struct {
	LogOn     bool   `json:"log_on"`
	LogOff    bool   `json:"log_off"`
	Player    string `json:"player"`
	Timestamp string `json:"timestamp"`
}

func StartNATSListener(nc *nats.Conn, discord *discordgo.Session, channelID string) {
	_, err := nc.Subscribe("arcadia.belsco", func(msg *nats.Msg) {
		handleServerEvent(msg, discord, channelID)
	})

	if err != nil {
		log.Fatalf("❌ Failed to subscribe to NATS subject: %v", err)
	}

	log.Println("📡 NATS listener running on subject 'arcadia.belsco'")
}

func handleServerEvent(msg *nats.Msg, discord *discordgo.Session, channelID string) {
	var event ServerEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("❌ Failed to parse event: %v", err)
		return
	}

	t, _ := time.Parse(time.RFC3339Nano, event.Timestamp)
	tStr := t.Local().Format("Jan 2 15:04:05")

	if event.LogOn {
		discord.ChannelMessageSend(channelID, fmt.Sprintf("✅ Player logged in: @%s at %s", event.Player, tStr))
	} else if event.LogOff {
		discord.ChannelMessageSend(channelID, fmt.Sprintf("🚪 Player logged out: @%s at %s", event.Player, tStr))
	} else {
		discord.ChannelMessageSend(channelID, fmt.Sprintf("⚠️ Unrecognized player event for @%s at %s", event.Player, tStr))
	}
}
