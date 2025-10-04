# Chirpy API Server 🐦

A modern, Go-based social messaging API server that allows users to post short messages called "chirps". Think of it as a minimalist Twitter-like backend service. Built by following a guide.

## Features ✨

- **Authentication & Authorization**

  - JWT-based authentication
  - Refresh token support
  - Secure password hashing

- **User Management**

  - User registration and login
  - Profile updates
  - Premium tier (Chirpy Red) support

- **Chirps**

  - Create, read, and delete chirps
  - Filter chirps by author
  - Sort by creation date
  - Message validation

- **Database**
  - PostgreSQL for reliable data storage
  - Efficient queries with proper indexing
  - Data consistency with foreign key constraints

## Quick Start 🚀

### Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher
- Make (optional)

### Installation

1. Clone the repository

   ```bash
   git clone https://github.com/lucho2027/chirpy.git
   cd chirpy
   ```

2. Set up environment variables

   ```bash
   cp .env.example .env
   # Edit .env with your configurations:
   # DB_URL=postgres://postgres:@localhost:5432/chirpy?sslmode=disable
   # JWT_SECRET=your-secret-key
   # POLKA_KEY=your-polka-key
   ```

3. Initialize the database

   ```bash
   createdb -U postgres chirpy
   ```

4. Run database migrations

   ```bash
   # Navigate to sql/schema directory
   cd sql/schema
   goose postgres://<DB_URL> up
   ```

5. Generate SQLC code

   ```bash
   # From project root
   sqlc generate
   ```

6. Install dependencies

   ```bash
   go mod tidy
   ```

7. Start the server
   ```bash
   make run
   # or
   go run main.go
   ```

The server will be available at `http://localhost:8080`

## API Documentation 📚

Detailed API documentation is available in the [/docs](/docs) directory:

- [Getting Started](/docs/getting-started.md)
- [API Endpoints](/docs/endpoints.md)
- [Authentication](/docs/authentication.md)
- [Database Schema](/docs/database.md)

## Project Structure 🏗️

```
.
├── api/            # API handlers and route definitions
├── internal/       # Internal packages
│   ├── auth/      # Authentication logic
│   └── database/  # Database operations
├── sql/           # SQL queries and schemas
│   ├── queries/   # SQLC queries
│   └── schema/    # Database migrations
└── docs/          # Documentation
```

## Development 🛠️

### Adding a New Endpoint

1. Add SQL migrations in `sql/schema/`
2. Run migrations: `goose postgres://<DB_URL> up`
3. Add SQL queries in `sql/queries/`
4. Generate SQLC code: `sqlc generate`
5. Implement the handler in `api/`
6. Add tests
7. Update documentation

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/auth

# Run with coverage
go test -cover ./...
```

## Security 🔒

- Passwords are hashed using bcrypt
- JWT tokens for secure authentication
- Input validation for all endpoints
- HTTPS required in production
- Database queries using prepared statements

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
