# Contributing to SwishRadar

We love your input! We want to make contributing to SwishRadar as easy and transparent as possible.

## Development Process

1. Fork the repo and create your branch from `master`
2. If you've added code that should be tested, add tests
3. Ensure the test suite passes
4. Make sure your code follows the existing style
5. Issue that pull request!

## Quick Start

### Backend Development
```bash
cd backend
go mod download
cp .env.example .env
# Add your credentials to .env
go run cmd/api/main.go
```

### Frontend Development
```bash
cd frontend
npm install
cp .env.example .env.local
# Add your credentials to .env.local
npm run dev
```

### Discord Bot Development
```bash
cd discord-bot
go mod download
cp .env.example .env
# Add your credentials to .env
go run main.go
```

## Code Style

### Go
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Add comments for exported functions
- Keep functions focused and small

### TypeScript/React
- Use TypeScript for all new code
- Follow ESLint rules
- Use functional components with hooks
- Keep components small and reusable

## Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

## Commit Messages

- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit first line to 72 characters
- Reference issues and pull requests after the first line

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
