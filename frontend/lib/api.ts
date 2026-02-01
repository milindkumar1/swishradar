// Use relative URLs for Vercel serverless functions in production
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || (typeof window !== 'undefined' ? '' : 'http://localhost:8081')

export interface LeagueInfo {
  id: number
  name: string
  year: number
  size: number
  current_week: number
}

export interface TeamOwner {
  displayName: string
  firstName: string
  lastName: string
  id: string
}

export interface Player {
  name: string
  position: string
  proTeam: string
  injured: boolean
}

export interface Team {
  id: number
  name: string
  owners: TeamOwner[]
  wins: number
  losses: number
  roster: Player[]
}

export interface Standing {
  rank: number
  team_name: string
  owners: TeamOwner[]
  wins: number
  losses: number
  points_for: number
  points_against: number
}

export interface FreeAgent {
  name: string
  position: string
  proTeam: string
  avg_points: number
  total_points: number
}

class ApiClient {
  private baseUrl: string

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl
  }

  private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`
    const response = await fetch(url, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    })

    if (!response.ok) {
      throw new Error(`API Error: ${response.statusText}`)
    }

    return response.json()
  }

  // ESPN endpoints
  async getESPNHealth() {
    return this.request('/api/espn/health')
  }

  async getLeague(): Promise<LeagueInfo> {
    return this.request('/api/espn/league')
  }

  async getTeams(): Promise<{ teams: Team[] }> {
    return this.request('/api/espn/teams')
  }

  async getStandings(): Promise<{ standings: Standing[] }> {
    return this.request('/api/espn/standings')
  }

  async getFreeAgents(limit: number = 50): Promise<{ players: FreeAgent[] }> {
    return this.request(`/api/espn/free-agents?limit=${limit}`)
  }

  // Analytics endpoints (coming soon)
  async getStreamingRecommendations() {
    return this.request('/api/v1/analytics/streaming')
  }

  async calculateTrade(team1Players: number[], team2Players: number[]) {
    return this.request('/api/v1/analytics/trade', {
      method: 'POST',
      body: JSON.stringify({ team1_players: team1Players, team2_players: team2Players }),
    })
  }

  async getPowerRankings() {
    return this.request('/api/v1/analytics/power-rankings')
  }

  async getMatchupPrediction(week: number) {
    return this.request(`/api/v1/analytics/matchup/${week}`)
  }

  // Player endpoints
  async getPlayers() {
    return this.request('/api/v1/players')
  }

  async getPlayer(id: number) {
    return this.request(`/api/v1/players/${id}`)
  }

  async getPlayerStats(id: number) {
    return this.request(`/api/v1/players/${id}/stats`)
  }
}

export const apiClient = new ApiClient(API_BASE_URL)
export const api = apiClient
