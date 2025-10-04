# Getting Started

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Environment Setup

Create a `.env` file in the root directory:

```env
PLATFORM=dev
DB_URL=postgres://postgres:@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=your-secret-key
POLKA_KEY=your-polka-key
```

## Database Setup

1. Create the database:

```bash
createdb -U postgres chirpy
```

2. Run migrations:

```bash
# Using make
make migrate-up

# Manual migration
# Follow instructions in sql/schema/README.md
```

## Running the Server

1. Install dependencies:

```bash
go mod tidy
```

2. Start the server:

```bash
# Using make
make run

# Manual start
go run main.go
```

The server will start on http://localhost:8080

## Development

- API endpoints are in the `/api` directory
- Database queries are in `/sql/queries`
- Authentication logic is in `/internal/auth`
- Database operations are in `/internal/database`

## Testing

Run tests:

```bash
# All tests
go test ./...

# Specific package
go test ./internal/auth

# With coverage
go test -cover ./...
```
