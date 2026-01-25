# Kasir API

REST API for Point of Sale system built with Go.

## Description

This API provides CRUD operations for managing products and categories in a retail/kasir system. Data is stored in-memory.

## API Endpoints

| Resource | Method | Endpoint | Description |
|----------|--------|----------|-------------|
| Products | GET | `/api/products` | [View Docs](docs/products-api.md) |
| Products | GET | `/api/products/{id}` | [View Docs](docs/products-api.md) |
| Products | POST | `/api/products` | [View Docs](docs/products-api.md) |
| Products | PUT | `/api/products/{id}` | [View Docs](docs/products-api.md) |
| Products | DELETE | `/api/products/{id}` | [View Docs](docs/products-api.md) |
| Categories | GET | `/api/categories` | [View Docs](docs/categories-api.md) |
| Categories | GET | `/api/categories/{id}` | [View Docs](docs/categories-api.md) |
| Categories | POST | `/api/categories` | [View Docs](docs/categories-api.md) |
| Categories | PUT | `/api/categories/{id}` | [View Docs](docs/categories-api.md) |
| Categories | DELETE | `/api/categories/{id}` | [View Docs](docs/categories-api.md) |

## How to Run

```bash
go run main.go
```

Server runs on `http://localhost:5500`

## Response Format

All endpoints return JSON with this structure:

```json
{
  "status": 200,
  "message": "Success message",
  "data": { ... }
}
```
