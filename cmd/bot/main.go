package main

import (
	"github.com/Zeethulhu/plebnet-discord-bot/internal/discord"
	"github.com/Zeethulhu/plebnet-discord-bot/internal/utils"
)

var logger = utils.NewLogger("Main")

func main() {
	discord.StartServer()
	logger.Println(getStartupMessage())
}

func getStartupMessage() string {
	return "Bot started."
}
