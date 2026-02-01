# Deploy SwishRadar

Quick deployment guide for all services.

## Prerequisites

- Railway account (free): https://railway.app
- Vercel account (free): https://vercel.com
- GitHub account with swishradar repo

## Step 1: Deploy ESPN Service to Railway

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login
railway login

# In espn-service folder
cd espn-service
railway init
railway up

# Get the deployed URL (e.g., https://espn-service-production.up.railway.app)
railway domain
```

**Environment variables to set in Railway dashboard:**
- `PORT` = `5001`

## Step 2: Deploy Go Backend to Railway

```bash
# In backend folder
cd ../backend
railway init
railway up

# Get the deployed URL
railway domain
```

**Environment variables to set in Railway dashboard:**
- `PORT` = `8081`
- `ESPN_SERVICE_URL` = `<your espn service URL from step 1>`

## Step 3: Deploy Frontend to Vercel

```bash
# Install Vercel CLI
npm install -g vercel

# In frontend folder
cd ../frontend

# Deploy
vercel --prod
```

**Environment variables to set in Vercel dashboard:**
- `NEXT_PUBLIC_API_URL` = `<your backend URL from step 2>`

## Step 4: Login to ESPN

1. Go to your deployed ESPN service URL
2. Click "Login with ESPN (Easy Way)"
3. Complete login in browser
4. Your credentials are saved

## Step 5: Access Your Dashboard

Go to your Vercel URL + `/dashboard`

---

## Alternative: Deploy All at Once via GitHub

### Option A: Railway (Backend + ESPN Service)

1. Push your code to GitHub
2. Go to Railway dashboard
3. Click "New Project" → "Deploy from GitHub repo"
4. Select `swishradar` repo
5. Railway will detect both services and deploy them
6. Set environment variables in Railway dashboard

### Option B: Vercel (Frontend)

1. Go to Vercel dashboard
2. Click "Import Project"
3. Select your `swishradar` repo
4. Set Root Directory to `frontend`
5. Add environment variable `NEXT_PUBLIC_API_URL`
6. Deploy

---

## Quick Commands

```bash
# Deploy ESPN Service
cd espn-service && railway up

# Deploy Backend
cd backend && railway up

# Deploy Frontend  
cd frontend && vercel --prod
```

---

## Free Tier Limits

**Railway:**
- $5 free credit per month
- Enough for 2 services running 24/7

**Vercel:**
- Unlimited deployments
- 100GB bandwidth per month
- Perfect for frontend

---

## Production URLs

After deployment, update these in your README:

- **Frontend**: `https://your-app.vercel.app`
- **Backend**: `https://your-backend.up.railway.app`
- **ESPN Service**: `https://your-espn.up.railway.app`
