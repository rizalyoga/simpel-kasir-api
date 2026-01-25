# Products API

## Get All Products

- **Endpoint:** `GET /api/products`
- **Description:** Retrieve all products

**Response (Success):**
```json
{
  "status": 200,
  "message": "Product list",
  "data": [
    { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 },
    { "id": 2, "name": "Kecap Manis Sedap", "price": 3000, "stock": 40 },
    { "id": 3, "name": "Gula", "price": 16000, "stock": 20 }
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
  "status": 200,
  "message": "Products details",
  "data": { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 }
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": 404,
  "message": "Product not found",
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
  "status": 201,
  "message": "Product added successfully",
  "data": { "id": 4, "name": "Teh Botol", "price": 5000, "stock": 25 }
}
```

**Response (Error - Invalid Body):**
```json
{
  "status": 400,
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
  "name": "Inmie Goreng",
  "price": 4000,
  "stock": 50
}
```

**Response (Success):**
```json
{
  "status": 200,
  "message": "Product updated successfully",
  "data": { "id": 1, "name": "Inmie Goreng", "price": 4000, "stock": 50 }
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": 404,
  "message": "Product not found",
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
  "status": 200,
  "message": "Product deleted successfully",
  "data": { "id": 1, "name": "Indomie", "price": 3500, "stock": 40 }
}
```

**Response (Error - Invalid ID):**
```json
{
  "status": 400,
  "message": "Invalid ID",
  "data": null
}
```

**Response (Error - Not Found):**
```json
{
  "status": 404,
  "message": "Product not found",
  "data": null
}
```
