# Quick Start Guide - User Microservice

This guide will help you get the User microservice up and running in 5 minutes.

## Prerequisites

- Go 1.25.5+
- PostgreSQL running locally or via Docker
- golang-migrate CLI (for database migrations)

## Step 1: Setup Environment

```bash
# Navigate to the user service directory
cd user

# Copy environment variables
cp .env.example .env

# Edit .env if needed (default values work for local development)
```

## Step 2: Setup Database

### Option A: Using Docker

```bash
# Start PostgreSQL container
docker run -d \
  --name user-postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=user_micro \
  -p 5432:5432 \
  postgres:15-alpine
```

### Option B: Using Existing PostgreSQL

Create the database:
```sql
CREATE DATABASE user_micro;
```

## Step 3: Run Migrations

```bash
# Apply database migrations
make db/migrations/up
```

## Step 4: Install Dependencies

```bash
# Download Go dependencies
go mod download
```

## Step 5: Run the Service

```bash
# Start the gRPC server
make run/api
```

You should see:
```
DB connection pool established
Starting gRPC server User service on port :8081
```

## Testing the Service

### Using grpcurl

Install grpcurl:
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

### 1. Create a User

```bash
grpcurl -plaintext -d '{
  "email": "john@example.com",
  "name": "John Doe",
  "password": "secure_password_123"
}' localhost:8081 user.UserService/CreateUser
```

Response:
```json
{
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "email": "john@example.com",
    "name": "John Doe"
  }
}
```

### 2. Get User (with JWT)

First, you need a valid JWT token with the user ID in claims. For testing, you can generate one:

```bash
# This assumes you have a JWT token
grpcurl -plaintext -d '{
  "jwt": "your.jwt.token.here"
}' localhost:8081 user.UserService/GetUser
```

Response:
```json
{
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "email": "john@example.com",
    "name": "John Doe"
  }
}
```

## Common Issues

### Port Already in Use

If port 8081 is already in use, change it in `.env`:
```bash
PORT=8082
```

### Database Connection Failed

Check if PostgreSQL is running:
```bash
psql -h localhost -U postgres -d user_micro
```

If connection fails, verify:
- PostgreSQL is running
- Credentials in `.env` match your database
- Host and port are correct

### Migration Failed

Reset the database:
```bash
# Drop and recreate
make db/migrations/down
make db/migrations/up
```

## Architecture Overview

```
Client Request (gRPC)
        ↓
    Handler Layer (handles gRPC requests)
        ↓
    Service Layer (business logic, password hashing, JWT verification)
        ↓
    Repository Layer (database operations)
        ↓
    PostgreSQL Database
```

## Key Features Implemented

✅ **CreateUser**: Registers new users with bcrypt password hashing
✅ **GetUser**: Retrieves user by JWT token validation
✅ **JWT Verification**: Validates tokens and extracts user ID from claims
✅ **Database Migrations**: Version-controlled schema changes
✅ **Graceful Shutdown**: Proper cleanup on termination
✅ **Environment Configuration**: Flexible config via environment variables

## Next Steps

1. **Add More Features**: Login endpoint, password reset, etc.
2. **Add Tests**: Write unit and integration tests
3. **Add Logging**: Implement structured logging
4. **Add Metrics**: Prometheus metrics for monitoring
5. **Add Validation**: More robust input validation
6. **Add Middleware**: Rate limiting, request logging, etc.

## Development Workflow

```bash
# 1. Make code changes
vim internal/service/service.go

# 2. Test locally
make run/api

# 3. Run tests (when you add them)
go test ./...

# 4. Build for production
make build/api

# 5. Run binary
./bin/api
```

## Docker Deployment

```bash
# Build image
docker build -t user-service:latest .

# Run container
docker run -d \
  --name user-service \
  -p 8081:8081 \
  --env-file .env \
  user-service:latest
```

## Production Checklist

Before deploying to production:

- [ ] Change `JWT_SECRET` to a strong random value
- [ ] Enable SSL/TLS for database (`POSTGRES_SSLMODE=require`)
- [ ] Set `ENVIRONMENT=PRODUCTION`
- [ ] Configure proper logging
- [ ] Set up monitoring and alerting
- [ ] Enable database backups
- [ ] Review and optimize connection pool settings
- [ ] Implement rate limiting
- [ ] Add request validation middleware
- [ ] Set up CI/CD pipeline

## Support

For issues or questions:
- Check the main [README.md](README.md) for detailed documentation
- Review the code structure and comments
- Examine the proto file for API definitions: `../proto/user.proto`

Happy coding! 🚀