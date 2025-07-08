# TastyBites Restaurant API

A complete backend implementation for a restaurant food ordering system built with Go, PostgreSQL, and Docker.

## Features

- 🔐 User authentication & authorization (JWT-based)
- 🍕 Menu browsing and management
- 🪑 Table reservation and availability checking
- 📝 Order placement and management
- 👑 Admin dashboard for order and table management
- 🔄 Table status reset functionality
- 🛡️ Role-based access control (User/Admin)

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL 15
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker & Docker Compose
- **Architecture**: Clean Architecture (Repository Pattern)

## Project Structure

```
TastyBites/
├── cmd/tastybites/          # Application entry point
├── internal/
│   ├── api/                 # HTTP handlers and routes
│   │   ├── handlers/        # Request handlers
│   │   ├── middlewares/     # Authentication & logging middleware
│   │   └── dto/            # Data transfer objects
│   ├── auth/               # JWT & bcrypt utilities
│   ├── config/             # Configuration management
│   ├── models/             # Domain models
│   ├── repo/               # Repository layer
│   │   ├── interfaces/     # Repository interfaces
│   │   └── postgres/       # PostgreSQL implementations
│   ├── usecases/           # Business logic layer
│   └── utils/              # Utility functions
├── db/migrations/          # Database schema and seed data
├── docker-compose.yml      # Docker services configuration
└── README.md              # This file
```

## Setup Instructions

### Prerequisites

- Docker and Docker Compose installed
- Go 1.21+ (for development)
- curl (for testing)
- jq (optional, for pretty JSON output)

### 1. Clone and Navigate

```bash
git clone <repository-url>
cd TastyBites
```

### 2. Start Database

```bash
# Start PostgreSQL with Docker Compose
docker compose up -d

# Verify database is running
docker ps | grep postgres
```

### 3. Environment Variables (Optional)

The application uses these environment variables with sensible defaults:

```bash
export TASTYBITES_DB_USERNAME=tastybites
export TASTYBITES_DB_PASSWORD=tastybitespass
export TASTYBITES_DB_DATABASE=tastybitesdb
export TASTYBITES_DB_HOST=localhost
export TASTYBITES_DB_PORT=5432
export TASTYBITES_SERVER_HOST=localhost
export TASTYBITES_SERVER_PORT=8080
export JWT_SECRET_KEY=yoursecretkey
```

### 4. Start the API Server

```bash
# Run the Go application
go run cmd/tastybites/main.go

# Or build and run
go build -o tastybites cmd/tastybites/main.go
./tastybites
```

The server will start on `http://localhost:8080`

### 5. Verify Installation

```bash
curl http://localhost:8080/ping
# Expected: pong
```

## API Documentation

### Base URL
```
http://localhost:8080
```

### Authentication

The API uses JWT Bearer tokens for authentication. Include the token in the Authorization header:

```bash
Authorization: Bearer <your-jwt-token>
```

## API Endpoints

### 🔓 Public Endpoints

#### Health Check
```bash
curl http://localhost:8080/ping
```

#### Get Menu Items
```bash
curl http://localhost:8080/menu | jq .
```

#### Get Available Tables
```bash
curl http://localhost:8080/tables | jq .
```

#### User Registration
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }' | jq .
```

#### User Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }' | jq .
```

### 🔐 User Protected Endpoints

First, get a user token:
```bash
USER_TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "password123"}' | jq -r '.token')
```

#### Create Order
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $USER_TOKEN" \
  -d '{
    "tableId": 1,
    "items": [
      {"itemId": 1, "quantity": 2, "price": 12.99},
      {"itemId": 3, "quantity": 1, "price": 14.99}
    ]
  }' | jq .
```

#### Get User Orders
```bash
curl -H "Authorization: Bearer $USER_TOKEN" \
  http://localhost:8080/orders | jq .
```

### 👑 Admin Protected Endpoints

First, get an admin token:
```bash
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@tastybites.com", "password": "password123"}' | jq -r '.token')
```

#### Get All Orders
```bash
curl -H "Authorization: Bearer $ADMIN_TOKEN" \
  http://localhost:8080/admin/orders | jq .
```

#### Get Order by Table ID
```bash
curl -H "Authorization: Bearer $ADMIN_TOKEN" \
  "http://localhost:8080/admin/tables/?tableId=1" | jq .
```

#### Reset Table Status
```bash
curl -X PATCH http://localhost:8080/admin/tables/3 \
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq .
```

## Sample Data

The database is automatically seeded with sample data:

### Default Users
| Email | Password | Role |
|-------|----------|------|
| admin@tastybites.com | password123 | admin |
| manager@tastybites.com | password123 | admin |
| john@example.com | password123 | user |
| jane@example.com | password123 | user |

### Sample Menu Items
- Margherita Pizza ($12.99)
- Chicken Caesar Salad ($11.49)
- Beef Burger ($14.99)
- Chocolate Brownie ($6.99)
- Pepperoni Pizza ($14.99)
- And 5 more items...

### Sample Tables
- 10 tables with various seating capacities
- Tables 3, 5, 8 are initially reserved
- Others are available for booking

## Complete Workflow Examples

### 1. Customer Order Flow

```bash
# 1. Register new customer
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Smith",
    "email": "alice@example.com",
    "password": "password123"
  }'

# 2. Login to get token
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "password123"}' | jq -r '.token')

# 3. Browse menu
curl http://localhost:8080/menu

# 4. Check available tables
curl http://localhost:8080/tables

# 5. Place order
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "tableId": 2,
    "items": [
      {"itemId": 1, "quantity": 1, "price": 12.99},
      {"itemId": 4, "quantity": 2, "price": 6.99}
    ]
  }'

# 6. Check order status
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/orders
```

### 2. Admin Management Flow

```bash
# 1. Login as admin
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@tastybites.com", "password": "password123"}' | jq -r '.token')

# 2. View all orders
curl -H "Authorization: Bearer $ADMIN_TOKEN" http://localhost:8080/admin/orders

# 3. Check specific table orders
curl -H "Authorization: Bearer $ADMIN_TOKEN" \
  "http://localhost:8080/admin/tables/?tableId=2"

# 4. Reset table when customer leaves
curl -X PATCH http://localhost:8080/admin/tables/2 \
  -H "Authorization: Bearer $ADMIN_TOKEN"
```

## Testing

### Run All Tests
```bash
# Make scripts executable
chmod +x *.sh

# Run comprehensive API test
./full_test.sh

# Test table booking validation
./table_validation_test.sh

# Test table reset functionality
./table_reset_test.sh
```

### Manual Testing

Test unauthorized access:
```bash
# Should return error
curl http://localhost:8080/orders

# Should return error  
curl -H "Authorization: Bearer invalid_token" http://localhost:8080/orders
```

Test role-based access:
```bash
# User trying to access admin endpoint (should fail)
curl -H "Authorization: Bearer $USER_TOKEN" http://localhost:8080/admin/orders
```

## Error Handling

The API returns structured error responses:

```json
{
  "status": "error", 
  "message": "Description of the error"
}
```

Common HTTP status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized  
- `403` - Forbidden
- `404` - Not Found
- `409` - Conflict (e.g., table already booked)
- `500` - Internal Server Error

## Database Management

### View Data
```bash
# Connect to database
docker exec -it postgres_db psql -U tastybites -d tastybitesdb

# View tables
\dt

# View users
SELECT id, name, email, role FROM users;

# View orders
SELECT id, user_id, table_id, status, total_price FROM orders;

# Exit
\q
```

### Reset Database
```bash
# Stop and remove containers
docker compose down

# Remove data volume
docker volume rm tastybites_postgres_data

# Restart with fresh data
docker compose up -d
```

## Development

### Project Dependencies
```bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build
go build -o bin/tastybites cmd/tastybites/main.go
```

### Adding New Features

1. **Models**: Add to `internal/models/`
2. **Repository**: Add interfaces to `internal/repo/interfaces/` and implementations to `internal/repo/postgres/`
3. **Usecases**: Add business logic to `internal/usecases/`
4. **Handlers**: Add HTTP handlers to `internal/api/handlers/`
5. **Routes**: Register routes in `internal/api/routes.go`

## Production Deployment

### Environment Variables
```bash
export TASTYBITES_DB_HOST=your-db-host
export TASTYBITES_DB_USERNAME=your-db-user
export TASTYBITES_DB_PASSWORD=your-db-password
export TASTYBITES_SERVER_HOST=0.0.0.0
export TASTYBITES_SERVER_PORT=8080
```

### Docker Production Build
```bash
# Build production image
docker build -t tastybites:latest .

# Run with environment variables
docker run -p 8080:8080 \
  -e TASTYBITES_DB_HOST=your-db-host \
  -e TASTYBITES_DB_USERNAME=your-db-user \
  -e TASTYBITES_DB_PASSWORD=your-db-password \
  tastybites:latest
```

## Security Features

- ✅ JWT-based authentication
- ✅ Password hashing with bcrypt  
- ✅ Role-based authorization
- ✅ SQL injection protection (parameterized queries)
- ✅ CORS middleware
- ✅ Request logging
- ✅ Panic recovery

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## License

This project is licensed under the MIT License.

---

**🍽️ TastyBites - Delicious food ordering made simple!**
