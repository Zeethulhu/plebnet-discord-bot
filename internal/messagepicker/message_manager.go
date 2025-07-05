package messagepicker

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// MessagesByCategory holds categorized message lists from YAML
type MessagesByCategory map[string][]string

// MessageManager handles category-aware message picking with recent tracking
type MessageManager struct {
	messages   MessagesByCategory
	recent     map[string][]int
	recentSize int
	rng        *rand.Rand
}

// NewManager loads messages from a YAML file
func NewManager(yamlPath string, recentSize int) (*MessageManager, error) {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	var messages MessagesByCategory
	if err := yaml.Unmarshal(data, &messages); err != nil {
		return nil, err
	}

	return &MessageManager{
		messages:   messages,
		recent:     make(map[string][]int),
		recentSize: recentSize,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

// Pick returns a message from the specified category with <player> substituted
func (mm *MessageManager) Pick(category, player string) (string, error) {
	msgList, ok := mm.messages[category]
	if !ok || len(msgList) == 0 {
		return "", fmt.Errorf("no messages found for category '%s'", category)
	}

	var idx int
	attempts := 0
	for {
		idx = mm.rng.Intn(len(msgList))
		if !mm.wasRecentlyUsed(category, idx) || attempts > 10 {
			break
		}
		attempts++
	}

	// Track recent
	mm.recent[category] = append(mm.recent[category], idx)
	if len(mm.recent[category]) > mm.recentSize {
		mm.recent[category] = mm.recent[category][1:]
	}

	player = fmt.Sprintf("**%s**", player)
	msg := strings.ReplaceAll(msgList[idx], "<player>", player)
	return msg, nil
}

func (mm *MessageManager) wasRecentlyUsed(category string, idx int) bool {
	for _, r := range mm.recent[category] {
		if r == idx {
			return true
		}
	}
	return false
}
