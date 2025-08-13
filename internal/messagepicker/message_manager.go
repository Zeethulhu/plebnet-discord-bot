package messagepicker

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/config"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/utils"
	"gopkg.in/yaml.v3"
)

var logger = utils.NewLogger("Messages")

// MessagesByCategory holds categorized message lists from YAML
type MessagesByCategory map[string][]string

// GameManager handles category-aware message picking with recent tracking for a single game
type GameManager struct {
	messages   MessagesByCategory
	recent     map[string][]int
	recentSize int
	rng        *rand.Rand
}

// Manager manages message templates for multiple games
type Manager struct {
	games map[string]*GameManager
}

// NewManager loads game-specific messages from YAML files based on GameConfig.Name
func NewManager(dir string, games []config.GameConfig, recentSize int) (*Manager, error) {
	m := &Manager{games: make(map[string]*GameManager)}
	for _, g := range games {
		filename := strings.ToLower(g.Name) + ".yaml"
		path := filepath.Join(dir, filename)
		data, err := os.ReadFile(path)
		if err != nil {
			logger.Printf("⚠️ skipping messages for %s: %v", g.Name, err)
			continue
		}
		var messages MessagesByCategory
		if err := yaml.Unmarshal(data, &messages); err != nil {
			logger.Printf("⚠️ skipping messages for %s: %v", g.Name, err)
			continue
		}
		m.games[g.Name] = &GameManager{
			messages:   messages,
			recent:     make(map[string][]int),
			recentSize: recentSize,
			rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
		}
	}
	return m, nil
}

// ForGame retrieves the GameManager for the given game name
func (m *Manager) ForGame(name string) (*GameManager, bool) {
	gm, ok := m.games[name]
	return gm, ok
}

// Pick returns a message from the specified category with <player> substituted
func (gm *GameManager) Pick(category, player string) (string, error) {
	msgList, ok := gm.messages[category]
	if !ok || len(msgList) == 0 {
		return "", fmt.Errorf("no messages found for category '%s'", category)
	}

	var idx int
	attempts := 0
	for {
		idx = gm.rng.Intn(len(msgList))
		if !gm.wasRecentlyUsed(category, idx) || attempts > 10 {
			break
		}
		attempts++
	}

	// Track recent
	gm.recent[category] = append(gm.recent[category], idx)
	if len(gm.recent[category]) > gm.recentSize {
		gm.recent[category] = gm.recent[category][1:]
	}

	player = fmt.Sprintf("**%s**", player)
	msg := strings.ReplaceAll(msgList[idx], "<player>", player)
	return msg, nil
}

func (gm *GameManager) wasRecentlyUsed(category string, idx int) bool {
	for _, r := range gm.recent[category] {
		if r == idx {
			return true
		}
	}
	return false
}
