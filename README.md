# POS API

This is a RESTful API for a Point of Sale (POS) system, built using Golang and PostgreSQL.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Database Migration](#database-migration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Requirements
Before installing, ensure you have the following installed:
- **Go** (1.18+)
- **PostgreSQL** (latest version recommended)
- **Golang Migrate** (for database migrations)
- **Git**

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repository/pos-api.git
   cd pos-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Configuration
This project uses `config.yml` for configuration. Before running the application, ensure you have a `config.yml` file in the root directory with the following structure:

```yaml
server:
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "your_username" # Edit this section
  password: "your_password" # Edit this section
  dbname: "your_db" # Edit this section
  sslmode: "disable"

jwt:
  secret: "your_secret_jwt" # Edit this section
```

## Database Migration
Run database migrations using Golang Migrate:
```sh
migrate -path migrations -database "postgres://your_db_user:your_db_password@localhost:5432/your_db_name?sslmode=disable" -verbose up
```

## Running the Application
Run the API server with:
```sh
go run main.go
```
The server will start on `http://localhost:8080` (or the port specified in `config.yml`).

## API Endpoints
| Method | Endpoint            | Description               |
|--------|--------------------|---------------------------|
| POST   | `/api/transactions` | Create a new transaction |
| GET    | `/api/products`     | Get all products         |
| GET    | `/api/categories`   | Get all categories       |

## Testing
You can test the API using [Postman](https://www.postman.com/) or `cURL`.

Example request to create a transaction:
```sh
curl -X POST "http://localhost:8080/api/transactions" -H "Content-Type: application/json" -d '{"user_id": 1, "payment_method": "cash", "products": [{"product_id": 2, "quantity": 3}]}'
```

---
Feel free to contribute or report issues! 🚀

