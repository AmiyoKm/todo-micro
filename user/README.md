# User Microservice

A gRPC-based user management microservice built with Go. This service handles user creation and authentication using JWT tokens.

## Features

- **User Registration**: Create new users with email, name, and password
- **User Authentication**: Retrieve user information using JWT tokens
- **Password Security**: Passwords are hashed using bcrypt
- **JWT Verification**: Validates and decodes JWT tokens to extract user information
- **PostgreSQL Database**: Persistent storage for user data

## Tech Stack

- **Go 1.25.5**: Programming language
- **gRPC**: Communication protocol
- **Protocol Buffers**: Data serialization
- **PostgreSQL**: Database
- **golang-jwt/jwt**: JWT token handling
- **bcrypt**: Password hashing
- **golang-migrate**: Database migrations

## Project Structure

```
user/
├── cmd/
│   └── api/
│       ├── main.go          # Application entry point
│       └── server.go        # gRPC server setup
├── configs/
│   └── config.go            # Configuration structs
├── gen/
│   └── userpb/              # Generated protobuf code
├── internal/
│   ├── handler/             # gRPC handlers
│   ├── service/             # Business logic
│   ├── repository/          # Database access layer
│   ├── model/               # Domain models
│   ├── infra/
│   │   └── db/              # Database connection
│   └── migrations/          # Database migrations
├── utils/                   # Utility functions
├── .env.example             # Example environment variables
├── Dockerfile               # Container image definition
├── Makefile                 # Build and development tasks
└── go.mod                   # Go module dependencies
```

## Prerequisites

- Go 1.25.5 or higher
- PostgreSQL 12 or higher
- golang-migrate CLI tool
- Protocol Buffer compiler (protoc)

## Installation

1. Clone the repository and navigate to the user service:
```bash
cd user
```

2. Copy the example environment file and configure it:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Install dependencies:
```bash
go mod download
```

4. Run database migrations:
```bash
make db/migrations/up
```

## Configuration

Environment variables (see `.env.example`):

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8081` |
| `ENVIRONMENT` | Environment (DEVELOPMENT/PRODUCTION) | `DEVELOPMENT` |
| `POSTGRES_HOST` | PostgreSQL host | `localhost` |
| `POSTGRES_PORT` | PostgreSQL port | `5432` |
| `POSTGRES_DB` | Database name | `user_micro` |
| `POSTGRES_USER` | Database user | `postgres` |
| `POSTGRES_PASSWORD` | Database password | `password` |
| `JWT_SECRET` | Secret key for JWT verification | `your-secret-key` |

## Running the Service

### Development Mode

```bash
make run/api
```

Or directly with Go:

```bash
go run ./cmd/api
```

### Production Mode

Build the binary:

```bash
make build/api
```

Run the binary:

```bash
./bin/api
```

### Using Docker

Build the image:

```bash
docker build -t user-service .
```

Run the container:

```bash
docker run -p 8081:8081 --env-file .env user-service
```

## API Reference

### gRPC Service

The service implements the following RPC methods defined in `proto/user.proto`:

#### CreateUser

Creates a new user account.

**Request:**
```protobuf
message CreateUserRequest {
    string email = 1;
    string name = 2;
    string password = 3;
}
```

**Response:**
```protobuf
message CreateUserResponse {
    User user = 1;
}
```

#### GetUser

Retrieves user information using a JWT token.

**Request:**
```protobuf
message GetUserRequest {
    string jwt = 1;
}
```

**Response:**
```protobuf
message GetUserResponse {
    User user = 1;
}
```

## Database Schema

### Users Table

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| `email` | VARCHAR(255) | UNIQUE, NOT NULL |
| `name` | VARCHAR(255) | NOT NULL |
| `password_hash` | TEXT | NOT NULL |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |

**Indexes:**
- `idx_users_email` on `email` column

## Development

### Available Make Commands

```bash
make help                     # Show all available commands
make run/api                  # Run the API server
make build/api               # Build the API binary
make db/migrations/new name=<name>  # Create a new migration
make db/migrations/up        # Apply all migrations
make db/migrations/down      # Rollback all migrations
make db/migrations/rollback  # Rollback specific number of migrations
```

### Adding New Migrations

```bash
make db/migrations/new name=add_user_fields
```

This creates two files:
- `000002_add_user_fields.up.sql` - Forward migration
- `000002_add_user_fields.down.sql` - Rollback migration

## Security

- Passwords are hashed using bcrypt with default cost (10)
- JWT tokens are verified using HMAC-SHA256
- Database connections support SSL/TLS
- Environment variables should be kept secure
- Never commit `.env` files to version control

## Error Handling

The service returns standard gRPC status codes:

- `OK` - Successful operation
- `INVALID_ARGUMENT` - Missing or invalid request parameters
- `UNAUTHENTICATED` - Invalid or expired JWT token
- `NOT_FOUND` - User not found
- `INTERNAL` - Internal server error

## Testing

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]