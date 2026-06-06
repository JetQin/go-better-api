# go-better-api

This file provides context about the project for AI assistants.

## Project Overview

- **Ecosystem**: Go

## Tech Stack

- Web Framework: gin
- Database: gorm
- Logging: logrus
- Auth Library: jwt

## Project Structure

```
go-better-api/
├── go.mod           # Module definition
├── cmd/
│   └── server/      # Server entry point
├── internal/        # Internal packages
```

## Common Commands

- `go mod tidy` - Install dependencies
- `go run cmd/server/main.go` - Start server
- `go test ./...` - Run tests
- `go fmt ./...` - Format code

## Maintenance

Keep CLAUDE.md updated when:

- Adding/removing dependencies
- Changing project structure
- Adding new features or services
- Modifying build/dev workflows

AI assistants should suggest updates to this file when they notice relevant changes.
