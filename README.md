# go-better-api

This project was created with [Better Fullstack](https://github.com/Marve10s/Better-Fullstack), a high-performance Go stack.

## Features

- **Go** - Fast, reliable, and efficient programming language
- **Gin** - High-performance HTTP web framework
- **GORM** - Full-featured ORM for Go
- **golang-jwt** - JWT token creation and validation

## Prerequisites

- [Go](https://go.dev/) 1.22 or higher

## Getting Started

First, copy the environment file:

```bash
cp .env.example .env
```

Then, install dependencies and run the server:

```bash
go mod tidy
go run cmd/server/main.go
```

The server will be running at [http://localhost:8080](http://localhost:8080).

## Database Setup

This project uses GORM with SQLite by default. To configure the database:

1. Copy the environment file:

```bash
cp .env.example .env
```

2. Update `DATABASE_URL` in `.env` with your database connection string.

Supported databases:

- SQLite (default): `DATABASE_URL=./data.db`
- PostgreSQL: `DATABASE_URL=postgres://user:pass@localhost:5432/dbname`

## Project Structure

```
go-better-api/
├── go.mod                # Module definition
├── cmd/
│   └── server/           # HTTP server entry point
│       └── main.go
├── internal/
│   ├── database/         # Database configuration
│   │   └── database.go
│   ├── models/           # GORM models
│   │   └── models.go
│   └── handlers/         # HTTP handlers
│       └── handlers.go
├── .env.example          # Environment variables template
└── .gitignore
```

## Available Commands

- `go build ./...`: Build all packages
- `go run cmd/server/main.go`: Run the server
- `go test ./...`: Run all tests
- `go fmt ./...`: Format code
- `go vet ./...`: Run static analysis
