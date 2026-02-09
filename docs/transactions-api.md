# Transactions API

## Checkout

- **Endpoint:** `POST /api/checkout`
- **Description:** Create a new transaction from checkout items. Calculates total amount, updates product stock, and records transaction details.

**Body (JSON):**
```json
{
  "items": [
    { "product_id": 1, "quantity": 2 },
    { "product_id": 3, "quantity": 1 }
  ]
}
```

**Response (Success):**
```json
{
  "status": "success",
  "code": 201,
  "message": "Checkout successful",
  "data": {
    "id": 1,
    "total_amount": 11400,
    "created_at": "2026-02-09T10:30:00Z",
    "details": [
      {
        "id": 1,
        "transaction_id": 1,
        "product_id": 1,
        "product_name": "Indomie",
        "quantity": 2,
        "subtotal": 7000
      },
      {
        "id": 2,
        "transaction_id": 1,
        "product_id": 3,
        "product_name": "Gula",
        "quantity": 1,
        "subtotal": 16000
      }
    ]
  }
}
```

**Response (Error - Product Not Found):**
```json
{
  "status": "error",
  "code": 500,
  "message": "product id 99 not found",
  "data": null
}
```

**Response (Error - Insufficient Stock):**
```json
{
  "status": "error",
  "code": 500,
  "message": "insufficient stock for product id 1",
  "data": null
}
```

**Response (Error - Invalid Body):**
```json
{
  "status": "error",
  "code": 400,
  "message": "Invalid request body",
  "data": null
}
```

**Response (Error - Method Not Allowed):**
```json
{
  "status": "error",
  "code": 405,
  "message": "Method not allowed",
  "data": null
}
```

---

## Database Tables

### transactions
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL | Primary key |
| total_amount | INT | Total transaction amount |
| created_at | TIMESTAMP | Transaction timestamp |

### transaction_details
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL | Primary key |
| transaction_id | INT | Foreign key to transactions |
| product_id | INT | Foreign key to products |
| quantity | INT | Number of items |
| subtotal | INT | Line item total (price × quantity) |

---

## Error Messages

| Scenario | Message |
|----------|---------|
| Product not found | "product id %d not found" |
| Insufficient stock | "insufficient stock for product id %d" |
| Invalid request body | "Invalid request body" |
| Method not allowed | "Method not allowed" |
