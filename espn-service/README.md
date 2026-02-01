# ESPN Service

Python microservice for accessing ESPN Fantasy Basketball API using the `espn-api` library.

## Setup

1. **Install Python 3.9+** (if not installed)
   ```powershell
   winget install Python.Python.3.11
   ```

2. **Create virtual environment**
   ```powershell
   cd espn-service
   python -m venv venv
   .\venv\Scripts\Activate.ps1
   ```

3. **Install dependencies**
   ```powershell
   pip install -r requirements.txt
   ```

4. **Configure environment**
   ```powershell
   copy .env.example .env
   notepad .env
   ```
   
   Add your ESPN credentials (same as before).

5. **Run the service**
   ```powershell
   python app.py
   ```

## Endpoints

- `GET /health` - Health check
- `GET /api/league` - League information
- `GET /api/teams` - All teams with rosters
- `GET /api/free-agents` - Available free agents (top 50)
- `GET /api/matchups` - Current week's matchups
- `GET /api/standings` - League standings

## Testing

```powershell
curl http://localhost:8082/health
curl http://localhost:8082/api/league
curl http://localhost:8082/api/teams
curl http://localhost:8082/api/free-agents
curl http://localhost:8082/api/matchups
curl http://localhost:8082/api/standings
```
