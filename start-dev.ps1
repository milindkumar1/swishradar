# SwishRadar Development Startup Script for Windows
# Run with: .\start-dev.ps1

Write-Host "🏀 Starting SwishRadar Development Environment..." -ForegroundColor Cyan
Write-Host ""

# Check if .env files exist
function Check-EnvFiles {
    $missingFiles = @()
    
    if (-not (Test-Path "backend\.env")) {
        $missingFiles += "backend\.env (copy from backend\.env.example)"
    }
    if (-not (Test-Path "frontend\.env.local")) {
        $missingFiles += "frontend\.env.local (copy from frontend\.env.example)"
    }
    
    if ($missingFiles.Count -gt 0) {
        Write-Host "⚠️  Missing environment files:" -ForegroundColor Yellow
        foreach ($file in $missingFiles) {
            Write-Host "   - $file" -ForegroundColor Yellow
        }
        return $false
    }
    return $true
}

# Check dependencies
function Check-Dependencies {
    $allGood = $true
    
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        Write-Host "❌ Go is not installed. Please install Go 1.21+" -ForegroundColor Red
        $allGood = $false
    }
    
    if (-not (Get-Command node -ErrorAction SilentlyContinue)) {
        Write-Host "❌ Node.js is not installed. Please install Node.js 18+" -ForegroundColor Red
        $allGood = $false
    }
    
    if ($allGood) {
        Write-Host "✅ Dependencies check passed" -ForegroundColor Green
    }
    return $allGood
}

# Main execution
if (-not (Check-Dependencies)) {
    exit 1
}

if (-not (Check-EnvFiles)) {
    exit 1
}

Write-Host ""
Write-Host "🚀 Starting Backend API..." -ForegroundColor Cyan
Start-Process -FilePath "powershell" -ArgumentList "-NoExit", "-Command", "cd backend; go run cmd/api/main.go"

Write-Host "⏳ Waiting for backend to start..." -ForegroundColor Yellow
Start-Sleep -Seconds 3

Write-Host "🚀 Starting Frontend..." -ForegroundColor Cyan
Start-Process -FilePath "powershell" -ArgumentList "-NoExit", "-Command", "cd frontend; npm run dev"

Write-Host ""
Write-Host "✨ SwishRadar is starting up!" -ForegroundColor Green
Write-Host "   Backend:  http://localhost:8080" -ForegroundColor White
Write-Host "   Frontend: http://localhost:3000" -ForegroundColor White
Write-Host ""
Write-Host "Both services are running in separate PowerShell windows." -ForegroundColor Yellow
Write-Host "Close those windows to stop the services." -ForegroundColor Yellow
