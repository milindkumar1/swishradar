# 🏀 SwishRadar Discord Bot

Discord bot integration for SwishRadar fantasy basketball analytics.

## Features

- `/matchup` - Get current week's matchup prediction
- `/streaming` - Top waiver wire recommendations
- `/powerrankings` - League power rankings
- `/player <name>` - Player stats and trends
- Daily scheduled reports (9 AM)

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Create a Discord application and bot at https://discord.com/developers/applications

3. Copy `.env.example` to `.env` and add your credentials:
```bash
cp .env.example .env
```

4. Invite the bot to your server using the OAuth2 URL generator

5. Run the bot:
```bash
go run main.go
```

## Deployment

### Railway (Free Tier)
```bash
railway login
railway init
railway up
```

### Fly.io (Free Tier)
```bash
fly launch
fly deploy
```

## Commands

All commands are available as slash commands in Discord.

**Prefix Commands:**
- `!ping` - Check if bot is online

**Slash Commands:**
- `/matchup` - This week's matchup
- `/streaming` - Waiver wire picks
- `/powerrankings` - Team rankings
- `/player <name>` - Player info

## Scheduled Reports

The bot automatically posts daily reports at 9 AM (configurable in `.env`):
- Top streaming picks for the day
- Matchup insights
- Injury updates
- Trending players
