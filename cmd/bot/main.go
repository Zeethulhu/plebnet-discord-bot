package main

import (
	"github.com/Zeethulhu/plebnet-discord-bot/internal/discord"
)

func main() {
	discord.StartServer()
	logger.Println(getStartupMessage())
}

func getStartupMessage() string {
	return "Bot started."
}
