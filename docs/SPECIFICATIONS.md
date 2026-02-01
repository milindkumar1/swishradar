# 📋 SwishRadar Project Specifications

This document outlines the complete technical architecture and implementation plan for SwishRadar.

## Table of Contents
1. [System Architecture](#system-architecture)
2. [Data Sources](#data-sources)
3. [Database Schema](#database-schema)
4. [API Endpoints](#api-endpoints)
5. [Algorithms](#algorithms)
6. [Deployment](#deployment)
7. [Development Roadmap](#development-roadmap)

## System Architecture

### High-Level Overview
```
┌─────────────────────────────┐
│    Next.js Frontend         │
│    (Vercel/Supabase)        │
│    - Dashboard              │
│    - Trade Calculator       │
│    - Waiver Wire UI         │
└──────────┬──────────────────┘
           │ HTTPS/REST
┌──────────▼──────────────────┐
│    Go Backend API           │
│    (Railway/Fly.io)         │
│    - ESPN Sync              │
│    - NBA Stats Ingestion    │
│    - Analytics Engine       │
└──────────┬──────────────────┘
           │
┌──────────▼──────────────────┐
│    Supabase PostgreSQL      │
│    - League data            │
│    - Player stats           │
│    - Backtesting results    │
│    - User credentials       │
└──────────┬──────────────────┘
           │
┌──────────▼──────────────────┐
│    Discord Bot (Go)         │
│    (Railway/Fly.io)         │
│    - Slash commands         │
│    - Scheduled reports      │
└─────────────────────────────┘
```

## Data Sources

### 1. ESPN Fantasy API (Unofficial)
**Purpose:** League sync, rosters, matchups, waiver wire

**Authentication:** SWID + espn_s2 cookies

**Key Endpoints:**
- League info: `/apis/v3/games/fba/seasons/{year}/segments/0/leagues/{leagueId}`
- Free agents: Add `?view=kona_player_info`
- Matchups: Add `?view=mMatchup`

**Rate Limits:** ~100 requests/hour (conservative)

### 2. NBA Stats API
**Purpose:** Player statistics, game logs, schedules

**Base URL:** `https://stats.nba.com/stats`

**Key Endpoints:**
- `/playergamelog` - Individual game performances
- `/teamgamelog` - Team schedules
- `/commonallplayers` - All active players

**Headers Required:**
```
User-Agent: Mozilla/5.0...
Referer: https://stats.nba.com/
```

### 3. Basketball Reference (Optional)
**Purpose:** Historical data validation, injury history

**Method:** Web scraping with rate limiting

## Database Schema

See `supabase/migrations/001_initial_schema.sql` for complete schema.

### Core Tables

**users**
- Stores ESPN credentials (encrypted)
- User preferences

**leagues**
- League configuration
- Scoring settings (JSONB)

**teams**
- Team rosters (JSONB)
- Owner linkage

**players**
- NBA player master list
- ESPN ID mapping

**player_stats_daily**
- Daily box scores
- Pre-calculated fantasy values

**weekly_streaming_results**
- Backtesting data
- Model performance tracking

**nba_team_schedules**
- Games per week by team
- Critical for streaming algorithm

## API Endpoints

### ESPN Sync
- `GET /api/v1/espn/league` - Fetch league info
- `GET /api/v1/espn/teams` - Get all teams
- `GET /api/v1/espn/waiver-wire` - Available players
- `POST /api/v1/espn/sync` - Manual sync trigger

### Analytics
- `GET /api/v1/analytics/streaming` - Waiver recommendations
- `POST /api/v1/analytics/trade` - Trade evaluation
- `GET /api/v1/analytics/power-rankings` - Team rankings
- `GET /api/v1/analytics/matchup/{week}` - Matchup prediction

### Players
- `GET /api/v1/players` - List all players
- `GET /api/v1/players/{id}` - Player details
- `GET /api/v1/players/{id}/stats` - Player game log

### Backtesting
- `POST /api/v1/backtest/run` - Execute backtest
- `GET /api/v1/backtest/results` - Get results

## Algorithms

### 1. Streaming Recommendation Engine

**Input Variables:**
- `games_this_week` - Number of games player's team has
- `per_game_value` - Average fantasy points per game
- `trend_delta` - Performance change (L5 vs L15 games)
- `minutes_stability` - Consistency of playing time
- `opportunity_factor` - Role changes, injuries on team

**Formula:**
```
StreamingScore = 
  (games_this_week * per_game_value * 1.0)
  + (trend_delta * 0.5)
  + (minutes_stability * 0.3)
  + (opportunity_factor * 0.4)
```

**Weights tuned via backtesting**

### 2. Trade Calculator

**Metrics:**
- Z-score normalization across categories
- Positional scarcity adjustments
- Injury risk factors
- ROS projections vs season averages

**Output:**
- Fairness score (0-100)
- Value delta per team
- Category-by-category impact

### 3. Power Rankings

**Method:** Monte Carlo simulation
- Simulate remaining schedule 10,000 times
- Use player distributions (not just averages)
- Calculate playoff odds
- Weekly win probabilities

### 4. Backtesting Framework

**Process:**
1. For each historical week:
   - Run streaming algorithm with data available at that time
   - Compare recommended players to actual performance
   - Track correlation between score and actual value
2. Metrics:
   - Spearman correlation
   - Hit rate (top 10 recommendations)
   - Value added vs naive approach

## Deployment

### Frontend (Vercel)
```bash
cd frontend
vercel --prod
```

**Environment Variables:**
- `NEXT_PUBLIC_SUPABASE_URL`
- `NEXT_PUBLIC_SUPABASE_ANON_KEY`
- `NEXT_PUBLIC_API_URL`

### Backend (Railway)
```bash
cd backend
railway login
railway up
```

**Environment Variables:**
- `SUPABASE_URL`
- `SUPABASE_KEY`
- `ESPN_SWID`
- `ESPN_S2`
- `PORT`

### Discord Bot (Railway)
```bash
cd discord-bot
railway up
```

**Environment Variables:**
- `DISCORD_TOKEN`
- `SUPABASE_URL`
- `SUPABASE_KEY`
- `API_URL`

### Database (Supabase)
1. Create project at supabase.com
2. Run migrations from `supabase/migrations/`
3. Configure Row Level Security (RLS) policies
4. Set up cron jobs for scheduled syncs

## Development Roadmap

### Phase 1: Foundation (Weeks 1-2)
- [x] Project setup and repository
- [ ] ESPN authentication module
- [ ] Basic league sync
- [ ] Database setup on Supabase
- [ ] Frontend shell with routing

### Phase 2: Data Pipeline (Weeks 3-4)
- [ ] NBA Stats API integration
- [ ] Daily stat ingestion job
- [ ] Schedule fetching
- [ ] Player database population
- [ ] Basic frontend data display

### Phase 3: Analytics MVP (Weeks 5-6)
- [ ] Streaming recommendation algorithm v1
- [ ] Backtesting framework
- [ ] Model tuning
- [ ] Results visualization
- [ ] Trade calculator v1

### Phase 4: Integration (Weeks 7-8)
- [ ] Discord bot commands
- [ ] Scheduled reports
- [ ] Power rankings
- [ ] Matchup predictions
- [ ] Polish and bug fixes

### Phase 5: Enhancement (Ongoing)
- [ ] Mobile responsive design
- [ ] User accounts and multi-league support
- [ ] Advanced trade scenarios
- [ ] Injury impact analysis
- [ ] Rest-of-season projections
- [ ] Historical league analysis

## Performance Considerations

### Caching Strategy
- Cache ESPN data for 15 minutes
- Cache NBA stats for 1 hour (longer for historical)
- Cache streaming recommendations until next sync

### Rate Limiting
- ESPN: Max 4 requests/minute
- NBA API: Max 20 requests/minute
- Implement exponential backoff

### Database Optimization
- Index on (player_id, date) for quick stat lookups
- Partition player_stats_daily by month
- Use JSONB for flexible league settings

## Security

### Sensitive Data
- Encrypt ESPN cookies in database
- Use Supabase RLS for user data isolation
- Never expose SWID/espn_s2 in API responses
- Environment variables for all secrets

### API Security
- Rate limiting per IP
- API key authentication for Discord bot
- CORS restrictions on frontend

## Testing Strategy

### Unit Tests
- Algorithm correctness
- Data transformations
- API response parsing

### Integration Tests
- End-to-end API flows
- Database operations
- External API mocking

### Backtesting Validation
- Historical accuracy metrics
- Cross-validation across seasons
- Edge case handling

## Monitoring

### Metrics to Track
- API response times
- ESPN sync success rate
- Algorithm accuracy (vs actual outcomes)
- User engagement (Discord/Web)

### Logging
- Structured logging (JSON)
- Error tracking with context
- Performance profiling

---

**Last Updated:** 2026-02-01
**Version:** 0.1.0
**Status:** In Development
