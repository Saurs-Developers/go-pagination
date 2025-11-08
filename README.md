# Go Pagination

A lightweight and generic pagination helper for Go applications that simplifies offset/limit queries and provides strongly-typed responses for APIs.

---

## 📦 Overview

This package provides:
- A simple `Pagination` struct for input (page, size, order, sort).
- A computed `Query` struct for database operations.
- A generic `Response[T]` struct for returning paginated data from APIs.
- Built-in validation with descriptive errors (`ErrPageOutOfRange`, `ErrSizeOutOfRange`).

---

## ⚙️ Installation

```bash
go get github.com/yourusername/pagination
```

---

## 🧩 Usage Example

### 1. Basic Pagination Flow

```go
package main

import (
    "fmt"
    "github.com/yourusername/pagination"
)

func main() {
    p := pagination.Pagination{Page: 2, Size: 10, OrderBy: "name", SortBy: "ASC"}

    totalItems := 95
    q, err := p.Values(totalItems)
    if err != nil {
        panic(err)
    }

    fmt.Printf("OFFSET: %d, LIMIT: %d\n", q.Offset, q.Limit)
    fmt.Printf("Order: %s %s\n", q.OrderBy, q.SortBy)
}
```

Output:
```
OFFSET: 10, LIMIT: 10
Order: name ASC
```

---

### 2. Building Database Queries

You can easily integrate `Query` values with SQL builders (like `squirrel`) or ORM libraries:

```go
sql := fmt.Sprintf("SELECT * FROM users ORDER BY %s %s LIMIT %d OFFSET %d",
    q.OrderBy, q.SortBy, q.Limit, q.Offset)
```

---

### 3. API Response Struct

For APIs, use the provided `Response[T]` type:

```go
users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}

res := pagination.Response[[]User]{}
res.Pagination.Page = p.Page
res.Pagination.Size = p.Size
res.Pagination.TotalPages = q.TotalPages
res.Pagination.TotalItems = totalItems
res.Pagination.NextPage = q.NextPage
res.Pagination.PrevPage = q.PrevPage
res.Pagination.OrderBy = q.OrderBy
res.Pagination.SortBy = q.SortBy
res.Data = users
```

---

## 🧠 API Reference

### **type Pagination**
Input structure defining user pagination preferences.

| Field | Type | Description |
|-------|------|-------------|
| `Page` | `int` | Current page (1-indexed). |
| `Size` | `int` | Number of items per page. |
| `OrderBy` | `string` | Column name to order results by. |
| `SortBy` | `string` | Sort direction ("ASC" or "DESC"). |

### **func (p *Pagination) Values(totalItems int) (Query, error)**

Computes database-safe pagination metadata.

Returns `Query` and validates the input:

- Returns `ErrSizeOutOfRange` if `Size <= 0`
- Returns `ErrPageOutOfRange` if `Page > TotalPages`

### **type Query**
Structure returned by `Values()` for DB queries.

| Field | Type | Description |
|--------|------|-------------|
| `Offset` | `int` | SQL offset value. |
| `Limit` | `int` | SQL limit value. |
| `TotalPages` | `int` | Computed total number of pages. |
| `OrderBy` | `string` | Order column. |
| `SortBy` | `string` | Sort direction. |
| `NextPage` | `int` | Next page index. |
| `PrevPage` | `int` | Previous page index. |

### **type Response[T any]**
Generic API-friendly response structure with embedded pagination metadata.

```go
type Response[T any] struct {
    Pagination struct {
        Page       int
        Size       int
        TotalPages int
        TotalItems int
        NextPage   int
        PrevPage   int
        OrderBy    string
        SortBy     string
    }
    Data T
}
```

---

## ⚠️ Errors

| Error | Meaning |
|--------|----------|
| `ErrPageOutOfRange` | Requested page number exceeds available pages. |
| `ErrSizeOutOfRange` | Page size is zero or negative. |

---

## ✅ Example Output

Example serialized response:

```json
{
  "pagination": {
    "page": 2,
    "size": 10,
    "total_pages": 5,
    "total_items": 45,
    "next_page": 3,
    "prev_page": 1,
    "order_by": "created_at",
    "sort_by": "DESC"
  },
  "data": [
    { "id": "1", "name": "Alice" },
    { "id": "2", "name": "Bob" }
  ]
}
```

---

## 🧩 License

MIT — feel free to use and modify.