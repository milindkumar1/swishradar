# SwishRadar - Quick Reference

## Project Structure

```
swishradar/
├── backend/              # Go API server
│   ├── cmd/
│   │   └── api/         # Main application entry
│   ├── internal/
│   │   ├── database/    # Database connection
│   │   ├── espn/        # ESPN API client
│   │   ├── nba/         # NBA Stats API client
│   │   └── models/      # Data models
│   ├── .env.example
│   └── go.mod
│
├── frontend/            # Next.js web app
│   ├── app/            # Next.js 14 App Router
│   ├── lib/            # Utilities and API client
│   ├── .env.example
│   └── package.json
│
├── discord-bot/        # Discord integration
│   ├── main.go
│   ├── .env.example
│   └── go.mod
│
├── supabase/
│   └── migrations/     # Database migrations
│
└── docs/
    ├── SPECIFICATIONS.md   # Full technical spec
    └── SETUP.md           # Development setup guide
```

## Key URLs & Resources

### Development
- **Backend API:** http://localhost:8080
- **Frontend:** http://localhost:3000
- **Health Check:** http://localhost:8080/health

### External Services
- **Supabase Dashboard:** https://app.supabase.com
- **Discord Developers:** https://discord.com/developers/applications
- **ESPN API:** https://fantasy.espn.com/apis/v3/games/fba/...
- **NBA Stats API:** https://stats.nba.com/stats

## Essential Commands

### Backend
```bash
cd backend
go run cmd/api/main.go          # Start API
go test ./...                    # Run tests
go mod tidy                      # Clean dependencies
```

### Frontend
```bash
cd frontend
npm run dev                      # Start dev server
npm run build                    # Build for production
npm run lint                     # Lint code
```

### Discord Bot
```bash
cd discord-bot
go run main.go                   # Start bot
```

### Git
```bash
git add .
git commit -m "Your message"
git push origin master
```

## Environment Variables Checklist

### Backend (.env)
- [ ] SUPABASE_URL
- [ ] SUPABASE_KEY
- [ ] ESPN_SWID
- [ ] ESPN_S2
- [ ] ESPN_LEAGUE_ID
- [ ] PORT

### Frontend (.env.local)
- [ ] NEXT_PUBLIC_SUPABASE_URL
- [ ] NEXT_PUBLIC_SUPABASE_ANON_KEY
- [ ] NEXT_PUBLIC_API_URL

### Discord Bot (.env)
- [ ] DISCORD_TOKEN
- [ ] SUPABASE_URL
- [ ] SUPABASE_KEY
- [ ] API_URL

## API Endpoints

### ESPN Sync
- `GET /api/v1/espn/league` - League info
- `GET /api/v1/espn/teams` - All teams
- `GET /api/v1/espn/waiver-wire` - Available players
- `POST /api/v1/espn/sync` - Sync league data

### Analytics
- `GET /api/v1/analytics/streaming` - Waiver recommendations
- `POST /api/v1/analytics/trade` - Trade calculator
- `GET /api/v1/analytics/power-rankings` - Power rankings
- `GET /api/v1/analytics/matchup/{week}` - Matchup prediction

### Players
- `GET /api/v1/players` - All players
- `GET /api/v1/players/{id}` - Player details
- `GET /api/v1/players/{id}/stats` - Player stats

## Discord Commands

- `/matchup` - Current week's matchup
- `/streaming` - Top waiver wire picks
- `/powerrankings` - League power rankings
- `/player <name>` - Player statistics

## Development Workflow

### Daily Development
1. Pull latest changes: `git pull`
2. Start backend: `cd backend && go run cmd/api/main.go`
3. Start frontend: `cd frontend && npm run dev`
4. Make changes
5. Test locally
6. Commit and push

### Adding New Features
1. Create feature branch: `git checkout -b feature/your-feature`
2. Implement feature
3. Test thoroughly
4. Update documentation
5. Create pull request

### Database Changes
1. Write migration in `supabase/migrations/`
2. Test locally
3. Apply to production Supabase project

## Deployment Checklist

### Before Deploying
- [ ] All tests pass
- [ ] Environment variables configured
- [ ] Database migrations applied
- [ ] API endpoints tested
- [ ] Frontend builds successfully
- [ ] Discord bot connects

### Deployment Steps
1. **Backend:** Deploy to Railway/Fly.io
2. **Frontend:** Deploy to Vercel
3. **Database:** Migrations on Supabase
4. **Discord Bot:** Deploy to Railway/Fly.io
5. **Domain:** Point DNS to Vercel

## Troubleshooting

### Backend Issues
- Port already in use? Change PORT in .env
- Database connection fails? Check Supabase credentials
- ESPN API 401? Get fresh cookies

### Frontend Issues
- Build fails? Clear `.next/` and rebuild
- API not connecting? Check NEXT_PUBLIC_API_URL
- TypeScript errors? Run `npm install`

### Discord Bot Issues
- Bot offline? Check Discord token
- Commands not working? Re-register commands
- No responses? Check API_URL connection

## Getting ESPN Cookies

1. Open ESPN Fantasy Basketball in browser
2. Press F12 (DevTools)
3. Go to Application tab > Cookies
4. Find `SWID` and `espn_s2`
5. Copy full values (including {})
6. Add to .env file

## Next Steps

- [ ] Set up Supabase project
- [ ] Get ESPN credentials
- [ ] Configure environment variables
- [ ] Run initial sync
- [ ] Test all features
- [ ] Deploy to production
- [ ] Invite friends to league
- [ ] Dominate your fantasy league! 🏆

---

For detailed information, see:
- [SETUP.md](docs/SETUP.md) - Complete setup guide
- [SPECIFICATIONS.md](docs/SPECIFICATIONS.md) - Full technical spec
- [README.md](README.md) - Project overview
