# plebnet-discord-bot
Private Discord Bot written in Go for a small community of friends.

## Configuration

The bot can load settings from a configuration file, environment variables and
command line flags.  Environment variables continue to work for container
deployments, but for local use you can copy `.env.example` to `.env`:

```
DISCORD_TOKEN="..."               # Discord bot token
DISCORD_EVENTS_CHANNEL="..."      # (optional) default channel ID for server event messages
DISCORD_NATS_ADDRESS="nats://127.0.0.1:4222"  # NATS server address
DISCORD_NATS_TOPIC="enshrouded"   # NATS subject to subscribe to
```

An example YAML configuration file is provided as `config.example.yaml` and can
be renamed to `config.yaml` or referenced via the `--config` flag when starting
the bot. A `games` list can also be included where each entry specifies the
Discord channel, NATS topic, and Steam RSS feed for a game:

```yaml
# config.yaml
discord_token: "TOKEN"
# Optional default channel. Each game entry can override discord_channel.
events_channel: "123456789012345678"
nats_address: "nats://127.0.0.1:4222"
nats_topic: "enshrouded-logs"

games:
  - name: "Enshrouded"
    discord_channel: "123456789012345678"
    nats_topic: "enshrouded-logs"
    steam_rss: "https://store.steampowered.com/feeds/news.xml?appid=1203620"
  - name: "Valheim"
    discord_channel: "234567890123456789"
    nats_topic: "valheim-logs"
    steam_rss: "https://store.steampowered.com/feeds/news.xml?appid=892970"
```

Message templates for each game can be placed under `internal/config/messages/<game>.yaml`,
but they are optional. If a template file is missing or unreadable the bot will
fall back to basic join/leave notifications.

### Usage

Run the bot using Go directly or with a built binary. Command line flags override
both environment variables and the config file.

```bash
go run ./cmd/bot --config config.yaml

# Enable verbose logging
go run ./cmd/bot --config config.yaml --verbose

# Override specific options
go run ./cmd/bot --token "$DISCORD_TOKEN" --channel "$DISCORD_CHANNEL"
```
