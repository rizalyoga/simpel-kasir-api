# Products API

## Get All Products

- **Endpoint:** `GET /api/products`
- **Description:** Retrieve all products from database
- **Query Parameters (Optional):**
  - `name` - Search products by name (case-insensitive partial match)

**Response (Success - All Products):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product list",
  "data": [
    { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 },
    { "id": 2, "name": "Kecap Manis Sedap", "price": 3000, "stock": 40 },
    { "id": 3, "name": "Gula", "price": 16000, "stock": 20 }
  ]
}
```

**Response (Success - Search by Name):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product list",
  "data": [
    { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 },
    { "id": 4, "name": "Indomie Goreng", "price": 4000, "stock": 25 }
  ]
}
```

---

## Get Product by ID

- **Endpoint:** `GET /api/products/{id}`
- **Description:** Retrieve a single product by ID

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product details",
  "data": { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 }
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": "error",
  "code": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": "error",
  "code": 404,
  "message": "product not found",
  "data": null
}
```

---

## Create Product

- **Endpoint:** `POST /api/products`
- **Description:** Create a new product

**Body (JSON):**
```json
{
  "name": "Teh Botol",
  "price": 5000,
  "stock": 25
}
```

**Response (Success):**
```json
{
  "status": "success",
  "code": 201,
  "message": "Product added successfully",
  "data": { "id": 4, "name": "Teh Botol", "price": 5000, "stock": 25 }
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

---

## Update Product

- **Endpoint:** `PUT /api/products/{id}`
- **Description:** Update an existing product

**Body (JSON):**
```json
{
  "name": "Indomie Goreng",
  "price": 4000,
  "stock": 50
}
```

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product updated successfully",
  "data": { "id": 1, "name": "Indomie Goreng", "price": 4000, "stock": 50 }
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": "error",
  "code": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": "error",
  "code": 404,
  "message": "product not found",
  "data": null
}
```

---

## Delete Product

- **Endpoint:** `DELETE /api/products/{id}`
- **Description:** Delete a product

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product deleted successfully",
  "data": null
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": "error",
  "code": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": "error",
  "code": 404,
  "message": "product not found",
  "data": null
}
```

---

## Error Messages

| Scenario | Message |
|----------|---------|
| Invalid ID | "Invalid ID" |
| Product not found | "product not found" |
| Invalid request body | "Invalid request body" |
| Method not allowed | "Method not allowed" |
