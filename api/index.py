from flask import Flask, jsonify, request, render_template_string
from espn_api.basketball import League
import os

app = Flask(__name__)

LEAGUE_ID = 1356604871

# Store credentials in memory (will reset on cold start)
credentials = {}

def get_creds():
    """Get credentials from memory or env vars"""
    if credentials:
        return credentials
    
    espn_s2 = os.getenv('ESPN_S2')
    swid = os.getenv('SWID')
    if espn_s2 and swid:
        return {'espn_s2': espn_s2, 'swid': swid}
    
    return None

def get_league():
    """Get ESPN league instance"""
    creds = get_creds()
    if not creds:
        return None, "Not logged in"
    
    try:
        league = League(
            league_id=LEAGUE_ID,
            year=2026,
            espn_s2=creds['espn_s2'],
            swid=creds['swid']
        )
        return league, None
    except:
        try:
            league = League(
                league_id=LEAGUE_ID,
                year=2025,
                espn_s2=creds['espn_s2'],
                swid=creds['swid']
            )
            return league, None
        except Exception as e:
            return None, str(e)

LOGIN_HTML = '''
<!DOCTYPE html>
<html>
<head>
    <title>SwishRadar - ESPN Login</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            padding: 40px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            max-width: 500px;
            width: 100%;
        }
        h1 { color: #333; margin-bottom: 10px; font-size: 32px; }
        .subtitle { color: #666; margin-bottom: 30px; font-size: 16px; }
        .status {
            padding: 15px;
            border-radius: 10px;
            margin-bottom: 20px;
            font-size: 14px;
        }
        .status.connected { background: #d4edda; color: #155724; border: 1px solid #c3e6cb; }
        .status.disconnected { background: #f8d7da; color: #721c24; border: 1px solid #f5c6cb; }
        .form-group { margin-bottom: 20px; }
        label { display: block; color: #333; font-weight: 600; margin-bottom: 8px; font-size: 14px; }
        input {
            width: 100%;
            padding: 12px;
            border: 2px solid #e1e4e8;
            border-radius: 8px;
            font-size: 14px;
            font-family: monospace;
        }
        input:focus { outline: none; border-color: #667eea; }
        .help-text { color: #666; font-size: 12px; margin-top: 5px; }
        button {
            width: 100%;
            padding: 14px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 8px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s;
        }
        button:hover { transform: translateY(-2px); box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4); }
        .logout-btn { background: #dc3545; margin-top: 10px; }
        .instructions {
            background: #f8f9fa;
            border-radius: 8px;
            padding: 15px;
            margin-bottom: 20px;
            font-size: 13px;
            line-height: 1.6;
        }
        .instructions ol { margin-left: 20px; margin-top: 10px; }
        .instructions li { margin-bottom: 8px; }
        code {
            background: #e9ecef;
            padding: 2px 6px;
            border-radius: 3px;
            font-size: 12px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🏀 SwishRadar</h1>
        <p class="subtitle">ESPN Fantasy Basketball</p>
        
        {% if connected %}
        <div class="status connected">
            ✅ Connected: <strong>{{ league_name }}</strong>
        </div>
        <button class="logout-btn" onclick="if(confirm('Disconnect from ESPN?')) window.location.href='/api/logout'">Disconnect</button>
        <div style="margin-top: 20px; text-align: center;">
            <a href="/dashboard" style="color: #667eea; text-decoration: none; font-weight: 600;">→ Go to Dashboard</a>
        </div>
        {% else %}
        <div class="status disconnected">
            ❌ Not connected to ESPN
        </div>
        
        <div class="instructions">
            <strong>How to get your ESPN cookies:</strong>
            <ol>
                <li>Go to <a href="https://fantasy.espn.com/basketball/" target="_blank">fantasy.espn.com</a> and log in</li>
                <li>Press <code>F12</code> to open Developer Tools</li>
                <li>Go to <strong>Application</strong> tab (Chrome) or <strong>Storage</strong> tab (Firefox)</li>
                <li>Click <strong>Cookies</strong> → <code>https://fantasy.espn.com</code></li>
                <li>Copy the <code>SWID</code> and <code>espn_s2</code> values</li>
            </ol>
        </div>
        
        <form method="POST">
            <div class="form-group">
                <label for="swid">SWID Cookie</label>
                <input type="text" id="swid" name="swid" placeholder="{XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" required>
                <p class="help-text">Include the curly braces { }</p>
            </div>
            
            <div class="form-group">
                <label for="espn_s2">espn_s2 Cookie</label>
                <input type="text" id="espn_s2" name="espn_s2" placeholder="AEBxxxxxxx..." required>
                <p class="help-text">Long string starting with AEB</p>
            </div>
            
            <button type="submit">Connect to ESPN</button>
        </form>
        {% endif %}
    </div>
</body>
</html>
'''

@app.route('/api', methods=['GET', 'POST'])
@app.route('/api/', methods=['GET', 'POST'])
def index():
    """Login page"""
    if request.method == 'POST':
        swid = request.form.get('swid', '').strip()
        espn_s2 = request.form.get('espn_s2', '').strip()
        
        if swid and espn_s2:
            credentials['swid'] = swid
            credentials['espn_s2'] = espn_s2
            return '<script>window.location.href="/api"</script>'
    
    league, error = get_league()
    connected = league is not None
    league_name = league.settings.name if league else None
    
    return render_template_string(LOGIN_HTML, connected=connected, league_name=league_name)

@app.route('/api/logout')
def logout():
    """Clear credentials"""
    credentials.clear()
    return '<script>window.location.href="/api"</script>'

@app.route('/api/health')
def health():
    """Health check"""
    league, _ = get_league()
    return jsonify({'status': 'healthy', 'connected': league is not None})

@app.route('/api/league')
def get_league_info():
    """Get league info"""
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    return jsonify({
        'id': league.league_id,
        'name': league.settings.name,
        'year': league.year,
        'current_week': league.current_week,
        'size': len(league.teams)
    })

@app.route('/api/teams')
def get_teams():
    """Get all teams"""
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    teams = []
    for team in league.teams:
        roster = []
        for player in team.roster:
            roster.append({
                'name': player.name,
                'position': getattr(player, 'position', 'N/A'),
                'proTeam': getattr(player, 'proTeam', 'N/A'),
                'injured': getattr(player, 'injured', False)
            })
        
        teams.append({
            'id': team.team_id,
            'name': team.team_name,
            'owners': team.owners if hasattr(team, 'owners') else [],
            'wins': getattr(team, 'wins', 0),
            'losses': getattr(team, 'losses', 0),
            'roster': roster
        })
    
    return jsonify({'teams': teams})

@app.route('/api/standings')
def get_standings():
    """Get standings"""
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    standings = []
    for team in sorted(league.teams, key=lambda x: getattr(x, 'standing', 99)):
        owners = team.owners if hasattr(team, 'owners') else []
        standings.append({
            'rank': getattr(team, 'standing', 0),
            'team_name': team.team_name,
            'owners': owners,
            'wins': getattr(team, 'wins', 0),
            'losses': getattr(team, 'losses', 0),
            'points_for': getattr(team, 'points_for', 0),
            'points_against': getattr(team, 'points_against', 0)
        })
    
    return jsonify({'standings': standings})

@app.route('/api/free-agents')
def get_free_agents():
    """Get free agents"""
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    limit = request.args.get('limit', 20, type=int)
    free_agents = league.free_agents(size=limit)
    
    players = []
    for player in free_agents[:limit]:
        players.append({
            'name': player.name,
            'position': getattr(player, 'position', 'N/A'),
            'proTeam': getattr(player, 'proTeam', 'N/A'),
            'avg_points': getattr(player, 'avg_points', 0),
            'total_points': getattr(player, 'total_points', 0)
        })
    
    return jsonify({'players': players})

# Vercel serverless handler
def handler(request, response):
    with app.request_context(request.environ):
        return app.full_dispatch_request()
