# REST User Service

A Go-based REST API service for user management with MongoDB integration.

## Features

- CRUD operations for user management
- MongoDB database integration
- Request logging middleware
- Input validation
- Docker support

## Tech Stack

- Go
- MongoDB
- Docker
- Gorilla Mux Router

## Prerequisites

- Go 1.x
- Docker
- MongoDB

## Getting Started

1. Start MongoDB:

```bash
docker start mongodb
```

2. Run the application:

```bash
go run cmd/api/main.go
```

## API Endpoints
```
GET     - /persons       - Get all users
GET     - /person/{id}   - Get user by ID
POST    - /person        - Create new user
PUT     - /person/{id}   - Update user
DELETE  - /person/{id}   - Delete user
```
## Project Structure
```
go-rest-api/
├── cmd/
│ └── api/
│ └── main.go
├── internal/
│ ├── handlers/
│ ├── models/
│ ├── database/
│ ├── middleware/
│ └── routes/
└── pkg/
└── utils/
```

## Running Tests

```bash
go test ./...
```
