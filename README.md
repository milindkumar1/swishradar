# 🏀 SwishRadar

**The ultimate fantasy basketball analytics platform for ESPN leagues**

SwishRadar provides advanced analytics, waiver wire recommendations, trade calculators, and power rankings for your ESPN Fantasy Basketball league — all powered by real-time data and backtested algorithms.

---

## 🎯 Features

### Core Analytics
- **🔄 Trade Calculator** - Evaluate trade fairness with multi-category analysis and positional scarcity
- **📊 Waiver Wire Engine** - Smart streaming recommendations based on games played, trends, and opportunity
- **🏆 Power Rankings** - Monte Carlo simulations for playoff odds and weekly matchup predictions
- **📈 Backtesting** - Validate algorithm accuracy against historical data

### Integrations
- **ESPN League Sync** - Real-time roster, matchup, and waiver wire data
- **Discord Bot** - Daily insights, matchup previews, and command-based queries
- **Custom Domain** - Hosted on your own domain for free

---

## 🏗️ Architecture

```
┌─────────────────────────────┐
│    Next.js Frontend         │
│    (Vercel/Supabase)        │
└──────────┬──────────────────┘
           │
┌──────────▼──────────────────┐
│    Go Backend API           │
│    (Railway/Fly.io)         │
└──────────┬──────────────────┘
           │
┌──────────▼──────────────────┐
│    Supabase PostgreSQL      │
│    - League data            │
│    - Player stats           │
│    - Backtesting results    │
└─────────────────────────────┘
```

### Data Sources (All Free!)
- ESPN Fantasy API (unofficial)
- NBA Stats API (official, no key required)
- Basketball Reference (optional scraping)

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- Supabase account (free tier)
- ESPN Fantasy Basketball league

### Backend Setup

```bash
cd backend
go mod download
cp .env.example .env
# Add your Supabase credentials and ESPN cookies
go run cmd/api/main.go
```

### Frontend Setup

```bash
cd frontend
npm install
cp .env.example .env.local
# Add your Supabase URL and API key
npm run dev
```

### Discord Bot Setup

```bash
cd discord-bot
go mod download
cp .env.example .env
# Add Discord bot token and Supabase credentials
go run main.go
```

---

## 📊 Database Schema

See `supabase/migrations/` for the complete schema:

- **users** - User accounts and ESPN credentials
- **leagues** - League configurations and settings
- **teams** - Rosters and team data
- **players** - NBA player information
- **player_stats_daily** - Historical performance data
- **weekly_streaming_results** - Backtesting results

---

## 🗺️ Roadmap

### Phase 1: Foundation (Weeks 1-2)
- [x] Project setup
- [ ] ESPN authentication
- [ ] League sync
- [ ] Database schema

### Phase 2: Data Ingestion (Weeks 3-4)
- [ ] NBA stats API integration
- [ ] Waiver wire algorithm MVP
- [ ] Weekly sync jobs

### Phase 3: Analytics (Weeks 5-6)
- [ ] Backtesting engine
- [ ] Trade calculator
- [ ] Power rankings

### Phase 4: Integrations (Weeks 7-8)
- [ ] Discord bot
- [ ] Frontend polish
- [ ] Deploy to production

---

## 🛠️ Tech Stack

**Backend**
- Go 1.21+ (API, data ingestion, algorithms)
- Chi router
- PostgreSQL (via Supabase)

**Frontend**
- Next.js 14 (App Router)
- TypeScript
- Tailwind CSS
- Supabase client

**Discord Bot**
- Go (discordgo)
- Scheduled jobs via cron

**Infrastructure**
- Supabase (PostgreSQL + Edge Functions)
- Vercel (Frontend hosting)
- Railway/Fly.io (Backend + Bot)

---

## 📝 Environment Variables

### Backend (.env)
```
SUPABASE_URL=
SUPABASE_KEY=
ESPN_SWID=
ESPN_S2=
PORT=8080
```

### Frontend (.env.local)
```
NEXT_PUBLIC_SUPABASE_URL=
NEXT_PUBLIC_SUPABASE_ANON_KEY=
NEXT_PUBLIC_API_URL=
```

### Discord Bot (.env)
```
DISCORD_TOKEN=
SUPABASE_URL=
SUPABASE_KEY=
```

---

## 🤝 Contributing

This is a personal project, but feel free to fork and adapt for your own league!

---

## 📜 License

MIT License - use freely for your fantasy basketball domination 🏆

---

**Built with ❤️ for fantasy basketball degenerates**
