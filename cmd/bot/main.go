package main

import (
	"fmt"

	"github.com/Zeethulhu/plebnet-discord-bot/internal/discord"
)

func main() {
	discord.StartServer()
	fmt.Println(getStartupMessage())
}

func getStartupMessage() string {
	return "Bot started."
}
