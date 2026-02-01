Write-Host "🏀 Starting ESPN Service..." -ForegroundColor Cyan

# Activate virtual environment
& .\venv\Scripts\Activate.ps1

# Start the service
python app.py
