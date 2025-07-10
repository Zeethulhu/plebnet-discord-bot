package timers

import (
	"context"
	"database/sql"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmcdole/gofeed"
)

// EnshroudedNewsTimer posts recent Steam RSS items to Discord.
type EnshroudedNewsTimer struct {
	ChannelID string
	db        *sql.DB
	parser    *gofeed.Parser
}

// NewEnshroudedNewsTimer creates the timer and registers it.
func NewEnshroudedNewsTimer(channelID, dbPath string) (*EnshroudedNewsTimer, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS posts (guid TEXT PRIMARY KEY)`); err != nil {
		return nil, err
	}
	t := &EnshroudedNewsTimer{ChannelID: channelID, db: db, parser: gofeed.NewParser()}
	Register(t)
	return t, nil
}

func (t *EnshroudedNewsTimer) Name() string { return "enshrouded_news" }

func (t *EnshroudedNewsTimer) Interval() time.Duration { return time.Hour }

func (t *EnshroudedNewsTimer) Run(ctx context.Context, s *discordgo.Session) {
	feed, err := t.parser.ParseURL("https://store.steampowered.com/feeds/news/app/1203620/")
	if err != nil {
		logger.Printf("❌ Failed to fetch Enshrouded Steam feed: %v", err)
		return
	}
	cutoff := time.Now().Add(-24 * time.Hour)
	for _, item := range feed.Items {
		if item.PublishedParsed == nil || item.PublishedParsed.Before(cutoff) {
			continue
		}
		guid := item.GUID
		if guid == "" {
			guid = item.Link
		}
		var exists int
		if err := t.db.QueryRow(`SELECT 1 FROM posts WHERE guid = ?`, guid).Scan(&exists); err == nil {
			continue // already posted
		}
		if _, err := s.ChannelMessageSend(t.ChannelID, item.Link); err != nil {
			logger.Printf("❌ Failed to post Enshrouded Steam news: %v", err)
			continue
		}
		if _, err := t.db.Exec(`INSERT INTO posts(guid) VALUES(?)`, guid); err != nil {
			logger.Printf("❌ Failed to record Enshrouded Steam news GUID: %v", err)
		}
	}
}
