'use client'

import { useEffect, useState } from 'react'
import { api } from '@/lib/api'
import type { LeagueInfo, Team, Standing, FreeAgent } from '@/lib/api'

export default function Dashboard() {
  const [league, setLeague] = useState<LeagueInfo | null>(null)
  const [teams, setTeams] = useState<Team[]>([])
  const [standings, setStandings] = useState<Standing[]>([])
  const [freeAgents, setFreeAgents] = useState<FreeAgent[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [activeTab, setActiveTab] = useState<'standings' | 'teams' | 'free-agents'>('standings')

  useEffect(() => {
    loadDashboardData()
  }, [])

  const loadDashboardData = async () => {
    try {
      setLoading(true)
      setError(null)

      // Load all data in parallel
      const [leagueData, teamsData, standingsData, freeAgentsData] = await Promise.all([
        api.getLeague(),
        api.getTeams(),
        api.getStandings(),
        api.getFreeAgents(10),
      ])

      setLeague(leagueData)
      setTeams(teamsData.teams)
      setStandings(standingsData.standings)
      setFreeAgents(freeAgentsData.players)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load dashboard data')
      console.error('Dashboard error:', err)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-blue-500 mx-auto mb-4"></div>
          <p className="text-white text-xl">Loading your league data...</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800 flex items-center justify-center">
        <div className="bg-red-900/50 border border-red-500 rounded-lg p-8 max-w-md">
          <h2 className="text-2xl font-bold text-white mb-4">Error</h2>
          <p className="text-red-200 mb-4">{error}</p>
          <button
            onClick={loadDashboardData}
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded"
          >
            Retry
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800">
      <div className="container mx-auto px-4 py-8">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-4xl font-bold text-white mb-2">
            {league?.name || 'SwishRadar Dashboard'}
          </h1>
          <p className="text-gray-300">
            {league?.year} Season • Week {league?.current_week} • {league?.size} Teams
          </p>
        </div>

        {/* Tabs */}
        <div className="mb-6 border-b border-gray-700">
          <div className="flex gap-4">
            <button
              onClick={() => setActiveTab('standings')}
              className={`pb-3 px-4 font-semibold transition-colors ${
                activeTab === 'standings'
                  ? 'text-blue-400 border-b-2 border-blue-400'
                  : 'text-gray-400 hover:text-gray-300'
              }`}
            >
              Standings
            </button>
            <button
              onClick={() => setActiveTab('teams')}
              className={`pb-3 px-4 font-semibold transition-colors ${
                activeTab === 'teams'
                  ? 'text-blue-400 border-b-2 border-blue-400'
                  : 'text-gray-400 hover:text-gray-300'
              }`}
            >
              Teams
            </button>
            <button
              onClick={() => setActiveTab('free-agents')}
              className={`pb-3 px-4 font-semibold transition-colors ${
                activeTab === 'free-agents'
                  ? 'text-blue-400 border-b-2 border-blue-400'
                  : 'text-gray-400 hover:text-gray-300'
              }`}
            >
              Free Agents
            </button>
          </div>
        </div>

        {/* Content */}
        {activeTab === 'standings' && (
          <div className="bg-gray-800 rounded-lg overflow-hidden">
            <table className="w-full">
              <thead className="bg-gray-900">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Rank
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Team
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Owner
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Record
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    PF
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    PA
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-700">
                {standings.map((standing) => (
                  <tr key={standing.rank} className="hover:bg-gray-750">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-white">
                      {standing.rank}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-white font-semibold">
                      {standing.team_name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {standing.owners[0]?.displayName || 'Unknown'}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {standing.wins}-{standing.losses}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {standing.points_for.toFixed(1)}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {standing.points_against.toFixed(1)}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        {activeTab === 'teams' && (
          <div className="grid md:grid-cols-2 gap-6">
            {teams.map((team) => (
              <div key={team.id} className="bg-gray-800 rounded-lg p-6">
                <div className="mb-4">
                  <h3 className="text-xl font-bold text-white mb-1">{team.name}</h3>
                  <p className="text-gray-400 text-sm">
                    {team.owners[0]?.displayName || 'Unknown'} • {team.wins}-{team.losses}
                  </p>
                </div>
                <div className="space-y-2">
                  <h4 className="text-sm font-semibold text-gray-400 uppercase">Roster</h4>
                  {team.roster.slice(0, 5).map((player, idx) => (
                    <div key={idx} className="flex justify-between items-center text-sm">
                      <span className="text-white">
                        {player.name}
                        {player.injured && <span className="text-red-400 ml-2">INJ</span>}
                      </span>
                      <span className="text-gray-400">
                        {player.position} - {player.proTeam}
                      </span>
                    </div>
                  ))}
                  {team.roster.length > 5 && (
                    <p className="text-gray-500 text-xs italic">
                      +{team.roster.length - 5} more players
                    </p>
                  )}
                </div>
              </div>
            ))}
          </div>
        )}

        {activeTab === 'free-agents' && (
          <div className="bg-gray-800 rounded-lg overflow-hidden">
            <table className="w-full">
              <thead className="bg-gray-900">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Player
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Position
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Team
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Avg Points
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                    Total
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-700">
                {freeAgents.map((player, idx) => (
                  <tr key={idx} className="hover:bg-gray-750">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-white">
                      {player.name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {player.position}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {player.proTeam}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {player.avg_points.toFixed(2)}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
                      {player.total_points.toFixed(1)}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  )
}
