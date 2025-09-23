# REST-API

A lightweight **User Management System** that provides full CRUD operations for user data.

## Introduction
This application is built with the **Gin** web framework in **Go**, offering high performance and easy scalability.  
Data is persisted in **PostgreSQL**, and the project follows a layered architecture (handlers → use cases → repositories).

## Features
- Extensible, clean design
- Fast and resource-efficient
- Graceful shutdown and error handling
- Configurable via environment variables

## Prerequisites
- Go 1.21+  
- PostgreSQL 13+  
- (Optional) Make for running provided build/test targets

## Installation
```bash
git clone https://github.com/safwanone2n/SRE-BOOTCAMP.git
cd SRE-BOOTCAMP/REST-API
go mod tidy
```


 ## Environment setup

### create a .env file (or export environment variables)

```env
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/restapi?sslmode=disable
cat > .env <<EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=restapi
EOF
```

## Running the Application
``` bash
go run ./cmd/server
```

## build the Application

```make
make build
```



## The API will be available at: http://localhost:8080

## Usage Examples


### Create a user
```bash
curl -X POST http://localhost:8080/users/create \
     -H "Content-Type: application/json" \
     -d '{"first_name":"John","last_name":"Doe","email":"john@example.com","phone_number":"1234567890"}'
```

### List users
``` bash
curl -X POST http://localhost:8080/users/list \
     -H "Content-Type: application/json" \
     -d '{"filter_params":{"offset":0,"limit":10}}'
```

## Project structure

```bash
.
├── cmd/server        # Application entry point
├── internal
│   ├── user          # Handlers, use cases, and repository
│   └── shared        # Common utilities and helpers
├── db/sqlc           # SQLC generated code
└── Makefile
```

