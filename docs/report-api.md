# Report API

## Today's Sales Summary

- **Endpoint:** `GET /api/report/today`
- **Description:** Get sales summary for today (current date)

**Response (Success):**

```json
{
  "status": "success",
  "code": 200,
  "message": "Today's sales summary",
  "data": {
    "total_revenue": 45000,
    "total_transaksi": 5,
    "produk_terlaris": {
      "nama": "Indomie Goreng",
      "qty_terjual": 12
    }
  }
}
```

---

## Sales Summary by Date Range

- **Endpoint:** `GET /api/report`
- **Query Parameters:**
  - `start_date` - Start date (YYYY-MM-DD format)
  - `end_date` - End date (YYYY-MM-DD format)

**Example Request:**

```
GET /api/report?start_date=2026-01-01&end_date=2026-02-10
```

**Response (Success):**

```json
{
  "status": "success",
  "code": 200,
  "message": "Sales summary",
  "data": {
    "total_revenue": 150000,
    "total_transaksi": 15,
    "produk_terlaris": {
      "nama": "Indomie Goreng",
      "qty_terjual": 35
    }
  }
}
```

**Response (Error - Missing Parameters):**

```json
{
  "status": "error",
  "code": 400,
  "message": "start_date and end_date are required",
  "data": null
}
```

---

## Response Fields

| Field                         | Type   | Description              |
| ----------------------------- | ------ | ------------------------ |
| `total_revenue`               | INT    | Total transaction amount |
| `total_transaksi`             | INT    | Number of transactions   |
| `produk_terlaris`             | OBJECT | Best selling product     |
| `produk_terlaris.nama`        | STRING | Product name             |
| `produk_terlaris.qty_terjual` | INT    | Total quantity sold      |

---

## Error Messages

| Scenario                    | Message                                |
| --------------------------- | -------------------------------------- |
| Missing start_date/end_date | "start_date and end_date are required" |
| Method not allowed          | "Method not allowed"                   |
