from flask import Flask, jsonify, request, redirect, render_template_string
from flask_cors import CORS
from espn_api.basketball import League
import os
import json
from dotenv import load_dotenv

load_dotenv()

app = Flask(__name__)
CORS(app)

LEAGUE_ID = int(os.getenv('LEAGUE_ID', 1356604871))
CREDS_FILE = 'espn_credentials.json'
IS_PRODUCTION = os.getenv('RAILWAY_ENVIRONMENT') is not None

def load_credentials():
    # In production, use environment variables
    if IS_PRODUCTION:
        espn_s2 = os.getenv('ESPN_S2')
        swid = os.getenv('SWID')
        if espn_s2 and swid:
            return {'espn_s2': espn_s2, 'swid': swid}
        return None
    
    # In development, use credentials file
    if os.path.exists(CREDS_FILE):
        try:
            with open(CREDS_FILE, 'r') as f:
                return json.load(f)
        except:
            pass
    return None

def save_credentials(espn_s2, swid):
    with open(CREDS_FILE, 'w') as f:
        json.dump({'espn_s2': espn_s2, 'swid': swid}, f)

def auto_login_espn():
    """Auto-login disabled in production. Use manual credentials."""
    if IS_PRODUCTION:
        return False, "Auto-login not available in production. Please set ESPN_S2 and SWID environment variables."
    
    return False, "Auto-login only works in local development with browser access."

def get_league():
    creds = load_credentials()
    if not creds:
        return None, "No credentials. Please login first."
    
    # Try current year and nearby years with the NEW ESPN API endpoint
    from datetime import datetime
    current_year = datetime.now().year
    years_to_try = [2025, 2026, current_year, 2024]
    
    for year in years_to_try:
        try:
            print(f"Trying year {year} with new ESPN API endpoint...")
            
            # The espn-api library should handle the new endpoint automatically
            # but let's make sure we're trying the right years
            league = League(
                league_id=LEAGUE_ID,
                year=year,
                espn_s2=creds['espn_s2'],
                swid=creds['swid']
            )
            print(f"SUCCESS! Connected to: {league.settings.name} (Year: {year})")
            return league, None
        except Exception as e:
            error_msg = str(e)
            print(f"Year {year} failed: {error_msg}")
            
            # If it's a 404, try next year. If it's 401/403, cookies are bad
            if "401" in error_msg or "403" in error_msg:
                print("Authentication error - cookies may be invalid")
    
    return None, "Could not connect. Try: 1) Make league public in settings, OR 2) Get fresh cookies while viewing your league page"

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
        .status.connected {
            background: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .status.disconnected {
            background: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .form-group { margin-bottom: 20px; }
        label {
            display: block;
            color: #333;
            font-weight: 600;
            margin-bottom: 8px;
            font-size: 14px;
        }
        input {
            width: 100%;
            padding: 12px;
            border: 2px solid #e1e4e8;
            border-radius: 8px;
            font-size: 14px;
            transition: border-color 0.3s;
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
        button:hover { transform: translateY(-2px); }
        .logout-btn {
            background: #dc3545;
            margin-top: 10px;
        }
        .instructions {
            background: #f8f9fa;
            border-radius: 10px;
            padding: 20px;
            margin-top: 20px;
        }
        .instructions h3 { color: #333; margin-bottom: 10px; font-size: 16px; }
        .instructions ol {
            margin-left: 20px;
            color: #666;
            font-size: 14px;
            line-height: 1.6;
        }
        .instructions li { margin-bottom: 8px; }
        .code {
            background: #2d3748;
            color: #48bb78;
            padding: 2px 6px;
            border-radius: 4px;
            font-family: monospace;
            font-size: 12px;
        }
        a { color: #667eea; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🏀 SwishRadar</h1>
        <p class="subtitle">ESPN Fantasy Basketball</p>
        
        {% if connected %}
        <div class="status connected">
            ✅ Connected to: <strong>{{ league_name }}</strong>
        </div>
        <button class="logout-btn" onclick="window.location.href='/logout'">Disconnect</button>
        {% else %}
        <div class="status disconnected">
            ❌ Not connected to ESPN
        </div>
        
        <form method="POST" action="/auto-login">
            <button type="submit" style="margin-bottom: 20px;">🔐 Login with ESPN (Easy Way)</button>
        </form>
        
        <details style="margin: 20px 0;">
            <summary style="cursor: pointer; color: #667eea; font-weight: 600;">Or paste cookies manually...</summary>
            <form method="POST" action="/login" style="margin-top: 20px;">
                <div class="form-group">
                    <label>ESPN SWID</label>
                    <input type="text" name="swid" placeholder="{XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" required>
                    <p class="help-text">Include the curly braces { }</p>
                </div>
                
                <div class="form-group">
                    <label>ESPN S2 Cookie</label>
                    <input type="text" name="espn_s2" placeholder="AEBxxxxxxx..." required>
                    <p class="help-text">Very long string from ESPN cookies</p>
                </div>
                
                <button type="submit">Connect to ESPN</button>
            </form>
        </details>
        
        <div class="instructions">
            <h3>📝 How to get your cookies:</h3>
            <ol>
                <li>Go to <a href="https://fantasy.espn.com/basketball/" target="_blank">fantasy.espn.com</a></li>
                <li>Log in and open your league</li>
                <li>Press <span class="code">F12</span> → <strong>Application</strong> tab</li>
                <li>Click <strong>Cookies</strong> → <span class="code">fantasy.espn.com</span></li>
                <li>Copy <span class="code">SWID</span> and <span class="code">espn_s2</span> values</li>
                <li>Paste them above and click Connect</li>
            </ol>
        </div>
        {% endif %}
    </div>
</body>
</html>
'''

@app.route('/')
def index():
    league, error = get_league()
    return render_template_string(
        LOGIN_HTML,
        connected=league is not None,
        league_name=league.settings.name if league else None
    )

@app.route('/auto-login', methods=['POST'])
def auto_login():
    """Automatically log in with browser automation"""
    success, message = auto_login_espn()
    
    if success:
        # Test connection
        league, error = get_league()
        if league:
            print(f"Auto-login successful! Connected to: {league.settings.name}")
        return redirect('/')
    else:
        return f"<h1>Login Failed</h1><p>{message}</p><a href='/'>Go Back</a>", 400

@app.route('/login', methods=['POST'])
def login():
    swid = request.form.get('swid', '').strip()
    espn_s2 = request.form.get('espn_s2', '').strip()
    
    if not swid or not espn_s2:
        return "Missing credentials", 400
    
    save_credentials(espn_s2, swid)
    print(f"Credentials saved!")
    
    league, error = get_league()
    if league:
        print(f"Connected to: {league.settings.name}")
    else:
        print(f"Failed: {error}")
    
    return redirect('/')

@app.route('/logout')
def logout():
    if os.path.exists(CREDS_FILE):
        os.remove(CREDS_FILE)
    print("Logged out")
    return redirect('/')

@app.route('/health')
def health():
    league, error = get_league()
    return jsonify({
        'status': 'healthy' if league else 'disconnected',
        'connected': league is not None
    })

@app.route('/api/league')
def get_league_info():
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    return jsonify({
        'id': league.league_id,
        'name': league.settings.name,
        'year': league.year,
        'size': len(league.teams),
        'current_week': league.current_week
    })

@app.route('/api/teams')
def get_teams():
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
            'owners': team.owners if hasattr(team, 'owners') else [getattr(team, 'owner', 'Unknown')],
            'wins': getattr(team, 'wins', 0),
            'losses': getattr(team, 'losses', 0),
            'roster': roster
        })
    
    return jsonify({'teams': teams})

@app.route('/api/free-agents')
def get_free_agents():
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    limit = request.args.get('limit', 50, type=int)
    free_agents = league.free_agents(size=limit)
    
    players = []
    for player in free_agents:
        players.append({
            'name': player.name,
            'position': player.position,
            'proTeam': player.proTeam,
            'avg_points': player.avg_points,
            'total_points': player.total_points
        })
    
    return jsonify({'players': players})

@app.route('/api/standings')
def get_standings():
    league, error = get_league()
    if not league:
        return jsonify({'error': error}), 401
    
    standings = []
    for team in sorted(league.teams, key=lambda x: getattr(x, 'standing', 99)):
        owners = team.owners if hasattr(team, 'owners') else [getattr(team, 'owner', 'Unknown')]
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

if __name__ == '__main__':
    port = int(os.getenv('PORT', 5001))
    print(f"\nSwishRadar ESPN Service")
    print(f"Port: {port}")
    print(f"Open: http://localhost:{port}\n")
    
    try:
        app.run(host='127.0.0.1', port=port, debug=False, use_reloader=False, threaded=True)
    except KeyboardInterrupt:
        print("\nShutting down...")
    except Exception as e:
        print(f"Error: {e}")
        import traceback
        traceback.print_exc()
        input("Press Enter to exit...")
