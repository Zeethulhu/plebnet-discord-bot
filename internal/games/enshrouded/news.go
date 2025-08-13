package enshrouded

import "github.com/Zeethulhu/plebnet-discord-bot/internal/timers"

const steamNewsFeed = "https://store.steampowered.com/feeds/news/app/1203620/"

// NewNewsTimer creates a SteamNewsTimer for Enshrouded.
func NewNewsTimer(channelID, dbPath string) (*timers.SteamNewsTimer, error) {
	return timers.NewSteamNewsTimer(channelID, steamNewsFeed, dbPath)
}
