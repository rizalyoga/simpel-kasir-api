# Categories API

## Get All Categories

- **Endpoint:** `GET /api/categories`
- **Description:** Retrieve all categories from database

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Category list",
  "data": [
    { "id": 1, "name": "Bahan Pokok", "description": "Bahan pokok sehari - hari." },
    { "id": 2, "name": "Snack", "description": "Snack atau jajanan." },
    { "id": 3, "name": "Soda", "description": "Minuman bersoda." }
  ]
}
```

---

## Get Category by ID

- **Endpoint:** `GET /api/categories/{id}`
- **Description:** Retrieve a single category by ID

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Category details",
  "data": { "id": 1, "name": "Bahan Pokok", "description": "Bahan pokok sehari - hari." }
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
  "message": "category not found",
  "data": null
}
```

---

## Create Category

- **Endpoint:** `POST /api/categories`
- **Description:** Create a new category

**Body (JSON):**
```json
{
  "name": "Susu",
  "description": "Minuman susu dan olahannya."
}
```

**Response (Success):**
```json
{
  "status": "success",
  "code": 201,
  "message": "Category added successfully",
  "data": { "id": 4, "name": "Susu", "description": "Minuman susu dan olahannya." }
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

## Update Category

- **Endpoint:** `PUT /api/categories/{id}`
- **Description:** Update an existing category

**Body (JSON):**
```json
{
  "name": "Bahan Pokok",
  "description": "Bahan-bahan pokok untuk memasak."
}
```

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Category updated successfully",
  "data": { "id": 1, "name": "Bahan Pokok", "description": "Bahan-bahan pokok untuk memasak." }
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
  "message": "category not found",
  "data": null
}
```

---

## Delete Category

- **Endpoint:** `DELETE /api/categories/{id}`
- **Description:** Delete a category

**Response (Success):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Category deleted successfully",
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
  "message": "category not found",
  "data": null
}
```

---

## Error Messages

| Scenario | Message |
|----------|---------|
| Invalid ID | "Invalid ID" |
| Category not found | "category not found" |
| Invalid request body | "Invalid request body" |
| Method not allowed | "Method not allowed" |
