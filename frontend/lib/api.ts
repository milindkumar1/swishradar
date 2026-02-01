const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export interface Player {
  id: number
  name: string
  position: string
  team: string
  espn_id?: number
}

export interface StreamingRecommendation {
  player: Player
  score: number
  games_this_week: number
  projected_value: number
  trend_delta: number
  reason: string
}

export interface TradeAnalysis {
  fairness_score: number
  team1_value_change: number
  team2_value_change: number
  winner: string
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

  // League endpoints
  async getLeague() {
    return this.request('/api/v1/espn/league')
  }

  async getTeams() {
    return this.request('/api/v1/espn/teams')
  }

  async syncLeague() {
    return this.request('/api/v1/espn/sync', { method: 'POST' })
  }

  // Analytics endpoints
  async getStreamingRecommendations(): Promise<StreamingRecommendation[]> {
    return this.request('/api/v1/analytics/streaming')
  }

  async calculateTrade(team1Players: number[], team2Players: number[]): Promise<TradeAnalysis> {
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
