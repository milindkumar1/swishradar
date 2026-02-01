# SwishRadar - Start Script
# Run this to start all services

Write-Host "Starting SwishRadar Services..." -ForegroundColor Cyan
Write-Host ""

# Check if services are already running
$existing = Get-Process python,node,go -ErrorAction SilentlyContinue | Where-Object {$_.Path -like "*nbafantasy*"}
if ($existing) {
    Write-Host "Stopping existing services..." -ForegroundColor Yellow
    $existing | Stop-Process -Force
    Start-Sleep -Seconds 2
}

Write-Host "[1/3] ESPN Service..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\espn-service'; .\venv\Scripts\python.exe app.py"
Start-Sleep -Seconds 3

Write-Host "[2/3] Go Backend..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\backend'; go run cmd/api/main.go"
Start-Sleep -Seconds 2

Write-Host "[3/3] Next.js Frontend..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\frontend'; npm run dev"

Write-Host ""
Write-Host "All services starting in separate windows!" -ForegroundColor Cyan
Write-Host ""
Write-Host "Access URLs:" -ForegroundColor Yellow
Write-Host "  Frontend:     http://localhost:3000/dashboard" -ForegroundColor White
Write-Host "  ESPN Login:   http://localhost:5001" -ForegroundColor White
Write-Host "  Backend API:  http://localhost:8081/health" -ForegroundColor White
Write-Host ""
Write-Host "Close the terminal windows to stop each service" -ForegroundColor Gray
