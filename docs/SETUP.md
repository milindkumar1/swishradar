# SwishRadar Development Setup

Complete guide to get SwishRadar running locally.

## Prerequisites

- **Go 1.21+** - [Download](https://go.dev/dl/)
- **Node.js 18+** - [Download](https://nodejs.org/)
- **Supabase Account** - [Sign up](https://supabase.com)
- **ESPN Fantasy League** - Active basketball league
- **Discord Account** (optional) - For bot integration

## Step-by-Step Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/swishradar.git
cd swishradar
```

### 2. Set Up Supabase

1. Create a new project at [supabase.com](https://supabase.com)
2. Go to Settings > Database and copy your connection string
3. Go to Settings > API and copy:
   - Project URL
   - Anon public key
   - Service role key (keep secret!)

4. Run migrations:
```bash
# Install Supabase CLI
npm install -g supabase

# Link to your project
supabase link --project-ref your-project-ref

# Run migrations
supabase db push
```

Or manually run the SQL from `supabase/migrations/001_initial_schema.sql` in the SQL Editor.

### 3. Get ESPN Credentials

1. Log in to ESPN Fantasy Basketball
2. Open browser DevTools (F12)
3. Go to Application > Cookies
4. Find and copy:
   - `SWID` cookie value
   - `espn_s2` cookie value
5. Keep these secure!

### 4. Backend Setup

```bash
cd backend

# Install Go dependencies
go mod download

# Create environment file
cp .env.example .env

# Edit .env with your credentials:
# - Supabase connection info
# - ESPN SWID and espn_s2
# - Your league ID

# Run the API
go run cmd/api/main.go
```

The API should start on http://localhost:8080

### 5. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Create environment file
cp .env.example .env.local

# Edit .env.local with:
# - Supabase URL and anon key
# - Backend API URL (http://localhost:8080)

# Run development server
npm run dev
```

The frontend should start on http://localhost:3000

### 6. Discord Bot Setup (Optional)

```bash
cd discord-bot

# Install dependencies
go mod download

# Create environment file
cp .env.example .env

# Create a Discord application:
# 1. Go to https://discord.com/developers/applications
# 2. Create New Application
# 3. Go to Bot section and create a bot
# 4. Copy the bot token
# 5. Enable required intents (Guild Messages, Direct Messages)
# 6. Go to OAuth2 > URL Generator
#    - Select 'bot' and 'applications.commands'
#    - Select permissions: Send Messages, Read Messages
# 7. Use generated URL to invite bot to your server

# Edit .env with:
# - Discord bot token
# - Supabase URL and key
# - Backend API URL

# Run the bot
go run main.go
```

## Verification

### Test Backend
```bash
curl http://localhost:8080/health
# Should return: {"status": "healthy"}
```

### Test Frontend
Open http://localhost:3000 in your browser

### Test Discord Bot
In your Discord server, try:
- `/matchup`
- `/streaming`
- `/powerrankings`

## Common Issues

### Backend won't start
- Check Go version: `go version`
- Verify .env file exists and has correct values
- Check if port 8080 is already in use

### Frontend errors
- Delete `node_modules` and `package-lock.json`, then `npm install`
- Check Node version: `node --version`
- Verify .env.local exists

### Database connection fails
- Verify Supabase credentials
- Check if your IP is allowed in Supabase settings
- Ensure migrations have been run

### ESPN API returns 401
- Your SWID or espn_s2 cookies may be expired
- Re-login to ESPN and get fresh cookies
- Make sure cookies aren't URL-encoded

## Next Steps

1. **Sync Your League**
   - Use the `/api/v1/espn/sync` endpoint
   - Or build the sync UI in the frontend

2. **Test Features**
   - Check waiver wire recommendations
   - Try the trade calculator
   - View power rankings

3. **Set Up Cron Jobs**
   - Schedule daily ESPN syncs
   - Schedule NBA stats updates
   - Configure Discord daily reports

4. **Deploy to Production**
   - See deployment guide in README.md

## Development Tips

- Use `air` for hot reloading Go code: `go install github.com/cosmtrek/air@latest`
- Frontend has hot reload built-in with Next.js
- Check logs in terminal for debugging
- Use Supabase Studio (web UI) to inspect database

## Getting Help

- Check [SPECIFICATIONS.md](docs/SPECIFICATIONS.md) for architecture details
- See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines
- Open an issue on GitHub for bugs or questions

---

Happy coding! 🏀
