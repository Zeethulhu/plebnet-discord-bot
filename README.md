# plebnet-discord-bot
Private Discord Bot written in Go for a small community of friends.

## Configuration

The bot reads its configuration from environment variables. Copy `.env.example`
to `.env` and fill in the values:

```
DISCORD_TOKEN="..."               # Discord bot token
DISCORD_EVENTS_CHANNEL="..."      # Channel ID for server event messages
DISCORD_NATS_ADDRESS="nats://127.0.0.1:4222"  # NATS server address
DISCORD_NATS_TOPIC="enshrouded"   # NATS subject to subscribe to
```
