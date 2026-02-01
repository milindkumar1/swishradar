-- SwishRadar Database Schema
-- Initial migration for core tables

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    espn_swid TEXT,
    espn_s2 TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Leagues table
CREATE TABLE IF NOT EXISTS leagues (
    id VARCHAR(50) PRIMARY KEY,
    season INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    settings_json JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Teams table
CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    league_id VARCHAR(50) NOT NULL REFERENCES leagues(id) ON DELETE CASCADE,
    owner_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    team_name VARCHAR(255) NOT NULL,
    roster_json JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(league_id, team_name)
);

-- Players table
CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    espn_id INTEGER UNIQUE,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(10) NOT NULL,
    team VARCHAR(10) NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Player stats (daily) table
CREATE TABLE IF NOT EXISTS player_stats_daily (
    id SERIAL PRIMARY KEY,
    player_id INTEGER NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    points DECIMAL(5,2) DEFAULT 0,
    rebounds DECIMAL(5,2) DEFAULT 0,
    assists DECIMAL(5,2) DEFAULT 0,
    steals DECIMAL(5,2) DEFAULT 0,
    blocks DECIMAL(5,2) DEFAULT 0,
    turnovers DECIMAL(5,2) DEFAULT 0,
    threes_made DECIMAL(5,2) DEFAULT 0,
    fgm DECIMAL(5,2) DEFAULT 0,
    fga DECIMAL(5,2) DEFAULT 0,
    ftm DECIMAL(5,2) DEFAULT 0,
    fta DECIMAL(5,2) DEFAULT 0,
    minutes DECIMAL(5,2) DEFAULT 0,
    fantasy_value DECIMAL(6,2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(player_id, date)
);

-- Weekly streaming results (for backtesting)
CREATE TABLE IF NOT EXISTS weekly_streaming_results (
    id SERIAL PRIMARY KEY,
    week INTEGER NOT NULL,
    season INTEGER NOT NULL,
    player_id INTEGER NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    actual_points DECIMAL(6,2),
    projected_points DECIMAL(6,2),
    model_rank INTEGER,
    actual_rank INTEGER,
    games_played INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(week, season, player_id)
);

-- Matchups table
CREATE TABLE IF NOT EXISTS matchups (
    id SERIAL PRIMARY KEY,
    league_id VARCHAR(50) NOT NULL REFERENCES leagues(id) ON DELETE CASCADE,
    week INTEGER NOT NULL,
    season INTEGER NOT NULL,
    team1_id INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    team2_id INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    team1_score DECIMAL(6,2),
    team2_score DECIMAL(6,2),
    predicted_winner INTEGER,
    actual_winner INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(league_id, week, season, team1_id, team2_id)
);

-- NBA team schedules
CREATE TABLE IF NOT EXISTS nba_team_schedules (
    id SERIAL PRIMARY KEY,
    team VARCHAR(10) NOT NULL,
    week INTEGER NOT NULL,
    season INTEGER NOT NULL,
    games_count INTEGER NOT NULL DEFAULT 0,
    game_dates JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(team, week, season)
);

-- Indexes for performance
CREATE INDEX idx_player_stats_player_date ON player_stats_daily(player_id, date DESC);
CREATE INDEX idx_player_stats_date ON player_stats_daily(date DESC);
CREATE INDEX idx_teams_league ON teams(league_id);
CREATE INDEX idx_matchups_league_week ON matchups(league_id, week, season);
CREATE INDEX idx_streaming_results_week ON weekly_streaming_results(week, season);
CREATE INDEX idx_players_active ON players(active) WHERE active = true;
CREATE INDEX idx_nba_schedules_week ON nba_team_schedules(week, season);

-- Updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Add updated_at triggers
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_leagues_updated_at BEFORE UPDATE ON leagues
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_teams_updated_at BEFORE UPDATE ON teams
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_players_updated_at BEFORE UPDATE ON players
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_nba_schedules_updated_at BEFORE UPDATE ON nba_team_schedules
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
