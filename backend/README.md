# Backend API

Go backend service for Music Migration Tool.

## Setup

```bash
# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env with your credentials
# nano .env

# Run the server
go run cmd/api/main.go
```

## API Endpoints

### Health Check
```
GET /health
```

### Authentication
```
GET  /api/auth/spotify           - Initiate Spotify OAuth
GET  /api/auth/spotify/callback  - Spotify OAuth callback
GET  /api/auth/apple              - Initiate Apple Music OAuth
GET  /api/auth/apple/callback     - Apple Music OAuth callback
GET  /api/auth/status             - Get auth status
```

### Migrations
```
POST /api/migrations              - Start new migration
GET  /api/migrations              - List all migrations
GET  /api/migrations/:id          - Get migration status
```

### Playlists
```
GET /api/playlists/:provider                  - Get user's playlists
GET /api/playlists/:provider/:playlistId      - Get playlist tracks
```

### WebSocket
```
WS  /api/ws  - Real-time migration progress
```

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── database/
│   │   ├── postgres.go          # PostgreSQL connection & migrations
│   │   └── redis.go             # Redis connection
│   ├── handlers/
│   │   ├── auth.go              # Authentication handlers
│   │   ├── migration.go         # Migration handlers
│   │   ├── playlist.go          # Playlist handlers
│   │   └── health.go            # Health check
│   ├── middleware/
│   │   └── error.go             # Error handling
│   ├── models/
│   │   └── models.go            # Data models
│   └── routes/
│       ├── auth.go              # Auth routes
│       ├── migration.go         # Migration routes
│       ├── playlist.go          # Playlist routes
│       └── websocket.go         # WebSocket routes
├── pkg/                         # TODO: External packages
│   ├── spotify/                 # Spotify API client
│   ├── apple/                   # Apple Music API client
│   └── matcher/                 # Track matching engine
├── .env.example
├── Dockerfile
├── go.mod
└── go.sum
```

## Development

```bash
# Run tests
go test ./...

# Run with hot reload (requires air)
air

# Format code
go fmt ./...

# Lint
golangci-lint run
```
