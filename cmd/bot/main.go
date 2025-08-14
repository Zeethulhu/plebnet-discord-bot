package main

import (
    "flag"

    "github.com/Zeethulhu/plebnet-discord-bot/internal/config"
    "github.com/Zeethulhu/plebnet-discord-bot/internal/discord"
    "github.com/Zeethulhu/plebnet-discord-bot/internal/utils"
)

var logger = utils.NewLogger("Main")

func main() {
    cfgFile := flag.String("config", "", "path to config file")
    token := flag.String("token", "", "discord bot token")
    channel := flag.String("channel", "", "discord events channel")
    natsAddr := flag.String("nats", "", "nats server address")
    natsTopic := flag.String("topic", "", "nats subject")
    flag.Parse()

    opts := config.Options{
        ConfigFile:    *cfgFile,
        DiscordToken:  *token,
        EventsChannel: *channel,
        NatsAddress:   *natsAddr,
        NatsTopic:     *natsTopic,
    }

    logger.Println(getStartupMessage())
    discord.StartServer(opts)
}

func getStartupMessage() string {
	return "Bot started."
}
