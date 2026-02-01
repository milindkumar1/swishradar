#!/bin/bash
# SwishRadar Development Startup Script

echo "🏀 Starting SwishRadar Development Environment..."
echo ""

# Check if .env files exist
check_env_files() {
    if [ ! -f backend/.env ]; then
        echo "⚠️  backend/.env not found. Copy from backend/.env.example"
        return 1
    fi
    if [ ! -f frontend/.env.local ]; then
        echo "⚠️  frontend/.env.local not found. Copy from frontend/.env.example"
        return 1
    fi
    return 0
}

# Check dependencies
check_dependencies() {
    if ! command -v go &> /dev/null; then
        echo "❌ Go is not installed. Please install Go 1.21+"
        return 1
    fi
    
    if ! command -v node &> /dev/null; then
        echo "❌ Node.js is not installed. Please install Node.js 18+"
        return 1
    fi
    
    echo "✅ Dependencies check passed"
    return 0
}

# Start backend
start_backend() {
    echo "🚀 Starting Backend API..."
    cd backend
    go run cmd/api/main.go &
    BACKEND_PID=$!
    cd ..
    echo "   Backend running on http://localhost:8080 (PID: $BACKEND_PID)"
}

# Start frontend
start_frontend() {
    echo "🚀 Starting Frontend..."
    cd frontend
    npm run dev &
    FRONTEND_PID=$!
    cd ..
    echo "   Frontend running on http://localhost:3000 (PID: $FRONTEND_PID)"
}

# Main
if ! check_dependencies; then
    exit 1
fi

if ! check_env_files; then
    exit 1
fi

start_backend
sleep 2
start_frontend

echo ""
echo "✨ SwishRadar is running!"
echo "   Backend:  http://localhost:8080"
echo "   Frontend: http://localhost:3000"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for Ctrl+C
trap "echo ''; echo '🛑 Shutting down...'; kill $BACKEND_PID $FRONTEND_PID; exit" INT
wait
