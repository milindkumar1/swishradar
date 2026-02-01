# 🏀 SwishRadar - Vercel Deployment Guide

## Quick Deploy to Vercel

### 1. Install Vercel CLI
```bash
npm install -g vercel
```

### 2. Login to Vercel
```bash
vercel login
```

### 3. Deploy
```bash
vercel --prod
```

## Environment Variables Setup

After deploying, add these environment variables in your Vercel dashboard:

1. Go to your project on [vercel.com](https://vercel.com)
2. Click **Settings** → **Environment Variables**
3. Add the following:

| Variable | Value | Description |
|----------|-------|-------------|
| `ESPN_S2` | Your ESPN espn_s2 cookie | Long string starting with AEB... |
| `SWID` | Your ESPN SWID cookie | Format: `{XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}` |
| `ESPN_LEAGUE_ID` | `1356604871` | Your ESPN league ID |

### How to Get ESPN Cookies

1. Go to [fantasy.espn.com/basketball](https://fantasy.espn.com/basketball/) and log in
2. Press `F12` to open Developer Tools
3. Go to **Application** tab (Chrome) or **Storage** tab (Firefox)
4. Click **Cookies** → `https://fantasy.espn.com`
5. Find and copy:
   - `SWID` - includes the curly braces `{ }`
   - `espn_s2` - long string starting with `AEB`

## Project Structure

```
nbafantasy/
├── api/
│   ├── index.py          # ESPN API serverless function
│   └── requirements.txt  # Python dependencies
├── frontend/
│   ├── app/
│   │   ├── dashboard/
│   │   │   └── page.tsx  # Main dashboard
│   │   ├── layout.tsx
│   │   └── page.tsx      # Home page
│   ├── lib/
│   │   ├── api.ts        # API client
│   │   └── supabase.ts
│   └── package.json
├── supabase/             # Database (future)
└── vercel.json           # Vercel configuration
```

## API Routes

After deployment, your API will be available at:

- `https://your-domain.vercel.app/api` - Login page
- `https://your-domain.vercel.app/api/health` - Health check
- `https://your-domain.vercel.app/api/league` - League info
- `https://your-domain.vercel.app/api/teams` - All teams
- `https://your-domain.vercel.app/api/standings` - Standings
- `https://your-domain.vercel.app/api/free-agents` - Free agents

## First Time Setup

1. Deploy to Vercel: `vercel --prod`
2. Set environment variables in Vercel dashboard
3. Visit your deployed URL
4. The dashboard will automatically load your league data

## Troubleshooting

### "Not logged in" error
- Make sure `ESPN_S2` and `SWID` are set in Vercel environment variables
- Cookies must include the curly braces for SWID: `{xxx-xxx-xxx}`
- Redeploy after adding environment variables

### Dashboard not loading
- Check the Vercel deployment logs
- Verify your league ID is correct
- Make sure you're using the 2026 season (or update year in `api/index.py`)

### Python errors
- Vercel automatically installs dependencies from `api/requirements.txt`
- Check build logs for any package installation errors

## Local Development (Optional)

To run locally for development:

```bash
# Frontend
cd frontend
npm install
npm run dev

# API (Python)
cd api
pip install -r requirements.txt
python index.py
```

Set up `.env` file in the root:
```env
ESPN_S2=your_espn_s2_cookie
SWID={your-swid-cookie}
ESPN_LEAGUE_ID=1356604871
```

## What's Next?

After successful deployment:
- ✅ View your league dashboard
- ✅ Check team standings
- ✅ Browse free agents
- 🔜 Trade calculator
- 🔜 Waiver wire recommendations
- 🔜 Power rankings

---

**Note**: The old `backend/`, `espn-service/`, and `discord-bot/` directories are no longer needed. You can safely delete them. Everything now runs as Vercel serverless functions.
