# SwishRadar# SwishRadar



Fantasy basketball analytics platform for ESPN leagues.Fantasy basketball analytics platform for ESPN leagues.



## Quick Start## Quick Start



### 1. Start All Services### 1. Start All Services



```powershell```powershell

# Terminal 1 - ESPN Service# Terminal 1 - ESPN Service

cd espn-servicecd espn-service

.\venv\Scripts\python.exe app.py.\venv\Scripts\python.exe app.py



# Terminal 2 - Backend API  # Terminal 2 - Backend API  

cd backendcd backend

go run cmd/api/main.gogo run cmd/api/main.go



# Terminal 3 - Frontend# Terminal 3 - Frontend

cd frontendcd frontend

npm run devnpm run dev

``````



### 2. Login to ESPN### 2. Login to ESPN



Go to http://localhost:5001 and click "Login with ESPN (Easy Way)"Go to http://localhost:5001 and click "Login with ESPN (Easy Way)"



### 3. View Dashboard### 3. View Dashboard



Go to http://localhost:3000/dashboardGo to http://localhost:3000/dashboard



## Services## Services



- **Frontend**: http://localhost:3000 (Next.js)- **Frontend**: http://localhost:3000 (Next.js)

- **Backend API**: http://localhost:8081 (Go)- **Backend API**: http://localhost:8081 (Go)

- **ESPN Service**: http://localhost:5001 (Python/Flask)- **ESPN Service**: http://localhost:5001 (Python/Flask)



## Features## Features



- Real-time league standings- Real-time league standings

- Team rosters and player stats- Team rosters and player stats

- Free agent recommendations- Free agent recommendations

- Auto-login with ESPN (no cookie copying needed)- Auto-login with ESPN (no cookie copying needed)



## Tech Stack## Tech Stack



- Next.js 14 (React, TypeScript, Tailwind)- Next.js 14 (React, TypeScript, Tailwind)

- Go 1.21+ (Chi router)- Go 1.21+ (Chi router)

- Python 3.13 (Flask, espn-api v0.45.1)- Python 3.13 (Flask, espn-api v0.45.1)



## Stop Services## Stop Services



```powershell```powershell

Get-Process python,node,go | Where-Object {$_.Path -like "*nbafantasy*"} | Stop-ProcessGet-Process python,node,go | Where-Object {$_.Path -like "*nbafantasy*"} | Stop-Process

``````


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
