# Kasir API

REST API for Point of Sale system built with Go.

## Description

This API provides CRUD operations for managing products and categories in a retail/kasir system. Features include product search by name and checkout transactions. Data is stored in Supabase PostgreSQL database.

## Architecture

```
HTTP Request → Handlers → Services → Repositories → Database (Supabase)
```

### Layers

| Layer            | Responsibility                             |
| ---------------- | ------------------------------------------ |
| **Handlers**     | HTTP request/response, routing, validation |
| **Services**     | Business logic, data transformation        |
| **Repositories** | Database operations, SQL queries           |
| **Database**     | PostgreSQL (Supabase)                      |

## API Endpoints

### Products

| Method | Endpoint             | Description                                     | Docs                                           |
| ------ | -------------------- | ----------------------------------------------- | ---------------------------------------------- |
| GET    | `/api/products`      | Get all products (with optional `?name=` query) | [Link](docs/products-api.md#get-all-products)  |
| POST   | `/api/products`      | Create a new product                            | [Link](docs/products-api.md#create-product)    |
| GET    | `/api/products/{id}` | Get product by ID                               | [Link](docs/products-api.md#get-product-by-id) |
| PUT    | `/api/products/{id}` | Update product                                  | [Link](docs/products-api.md#update-product)    |
| DELETE | `/api/products/{id}` | Delete product                                  | [Link](docs/products-api.md#delete-product)    |

### Categories

| Method | Endpoint               | Description           | Docs                                              |
| ------ | ---------------------- | --------------------- | ------------------------------------------------- |
| GET    | `/api/categories`      | Get all categories    | [Link](docs/categories-api.md#get-all-categories) |
| POST   | `/api/categories`      | Create a new category | [Link](docs/categories-api.md#create-category)    |
| GET    | `/api/categories/{id}` | Get category by ID    | [Link](docs/categories-api.md#get-category-by-id) |
| PUT    | `/api/categories/{id}` | Update category       | [Link](docs/categories-api.md#update-category)    |
| DELETE | `/api/categories/{id}` | Delete category       | [Link](docs/categories-api.md#delete-category)    |

### Transactions

| Method | Endpoint        | Description                 | Docs                                      |
| ------ | --------------- | --------------------------- | ----------------------------------------- |
| POST   | `/api/checkout` | Create checkout transaction | [Link](docs/transactions-api.md#checkout) |

### Reports

| Method | Endpoint            | Description           | Docs                                                   |
| ------ | ------------------- | --------------------- | ------------------------------------------------------ |
| GET    | `/api/report/today` | Today's sales summary | [Link](docs/report-api.md#todays-sales-summary)        |
| GET    | `/api/report`       | Sales by date range   | [Link](docs/report-api.md#sales-summary-by-date-range) |

## Database

Uses Supabase PostgreSQL with pgx driver.

### Environment Variables

Create `.env` file:

```env
PORT=6543
DB_CONN=postgresql://postgres.qlybeoshtgfmsvkkgsjv:%40password@host:6543/postgres?sslmode=require
```

## How to Run

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

Server runs on `http://localhost:6543`

## Response Format

All endpoints return JSON with this structure:

```json
{
  "status": "success",
  "code": 200,
  "message": "Success message",
  "data": { ... }
}
```

### Response Status Values

| Status    | Meaning           |
| --------- | ----------------- |
| `success` | Request succeeded |
| `error`   | Request failed    |

### HTTP Status Codes

| Code | Meaning                     |
| ---- | --------------------------- |
| 200  | OK                          |
| 201  | Created                     |
| 400  | Bad Request (invalid input) |
| 404  | Not Found                   |
| 405  | Method Not Allowed          |
| 500  | Internal Server Error       |

## Project Structure

```
kasir-api-bootcamp/
├── common/
│   ├── errors/           # Custom error types
│   └── handlers/         # Shared response utilities
├── database/
│   └── database.go       # Database connection (pgx)
├── repositories/
│   ├── product_repository.go
│   ├── category_repository.go
│   └── transaction_repository.go
├── services/
│   ├── product_service.go
│   ├── category_service.go
│   └── transaction_service.go
├── handlers/
│   ├── product_handler.go
│   ├── category_handler.go
│   ├── transaction_handler.go
│   └── report_handler.go
├── models/
│   ├── products.go
│   ├── categories.go
│   ├── transactions.go
│   └── report.go
├── docs/
│   ├── products-api.md
│   ├── categories-api.md
│   ├── transactions-api.md
│   └── report-api.md
├── main.go
└── .env
```

## Documentation

- [Products API](docs/products-api.md)
- [Categories API](docs/categories-api.md)
- [Transactions API](docs/transactions-api.md)
- [Report API](docs/report-api.md)
