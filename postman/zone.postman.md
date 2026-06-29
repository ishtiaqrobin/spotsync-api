# Spotsync API — Zone Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation description",
  "data": { ... }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error description",
  "errors": "Error details"
}
```

---

## 1. Get All Zones (Public)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones`
- **Headers:** None required
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "Parking zones retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Terminal 1 EV Charging",
      "type": "ev_charging",
      "total_capacity": 20,
      "available_spots": 14,
      "price_per_hour": 5.5,
      "created_at": "2026-06-29T10:30:00+06:00",
      "updated_at": "2026-06-29T10:30:00+06:00"
    },
    {
      "id": 2,
      "name": "Downtown Parking",
      "type": "general",
      "total_capacity": 50,
      "available_spots": 30,
      "price_per_hour": 2.5,
      "created_at": "2026-06-29T09:00:00+06:00",
      "updated_at": "2026-06-29T09:00:00+06:00"
    }
  ]
}
```

---

## 2. Get Zone By ID (Public)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones/1`
- **Headers:** None required
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "Parking zone retrieved successfully",
  "data": {
    "id": 1,
    "name": "Terminal 1 EV Charging",
    "type": "ev_charging",
    "total_capacity": 20,
    "available_spots": 14,
    "price_per_hour": 5.5,
    "created_at": "2026-06-29T10:30:00+06:00",
    "updated_at": "2026-06-29T10:30:00+06:00"
  }
}
```

---

## 3. Get Zone By ID — Not Found (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones/9999`
- **Headers:** None required
- **Expected Response (404 Not Found):**

```json
{
  "success": false,
  "message": "Parking zone not found",
  "errors": null
}
```

---

## 4. Create Zone — General Type (Admin Only)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Downtown Parking",
  "type": "general",
  "total_capacity": 50,
  "price_per_hour": 2.5
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "Parking zone created successfully",
  "data": {
    "id": 1,
    "name": "Downtown Parking",
    "type": "general",
    "total_capacity": 50,
    "available_spots": 50,
    "price_per_hour": 2.5,
    "created_at": "2026-06-29T10:30:00+06:00",
    "updated_at": "2026-06-29T10:30:00+06:00"
  }
}
```

---

## 5. Create Zone — EV Charging Type (Admin Only)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Terminal 1 EV Charging",
  "type": "ev_charging",
  "total_capacity": 20,
  "price_per_hour": 5.5
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "Parking zone created successfully",
  "data": {
    "id": 2,
    "name": "Terminal 1 EV Charging",
    "type": "ev_charging",
    "total_capacity": 20,
    "available_spots": 20,
    "price_per_hour": 5.5,
    "created_at": "2026-06-29T10:31:00+06:00",
    "updated_at": "2026-06-29T10:31:00+06:00"
  }
}
```

---

## 6. Create Zone — Covered Type (Admin Only)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Mall Covered Parking",
  "type": "covered",
  "total_capacity": 30,
  "price_per_hour": 3.0
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "Parking zone created successfully",
  "data": {
    "id": 3,
    "name": "Mall Covered Parking",
    "type": "covered",
    "total_capacity": 30,
    "available_spots": 30,
    "price_per_hour": 3.0,
    "created_at": "2026-06-29T10:32:00+06:00",
    "updated_at": "2026-06-29T10:32:00+06:00"
  }
}
```

---

## 7. Create Zone — Without Auth (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "Unauthorized Zone",
  "type": "general",
  "total_capacity": 10,
  "price_per_hour": 2.0
}
```

- **Expected Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Missing authorization header",
  "errors": null
}
```

---

## 8. Create Zone — Driver Token (Error, Admin Only)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Driver Zone Attempt",
  "type": "general",
  "total_capacity": 10,
  "price_per_hour": 2.0
}
```

- **Expected Response (403 Forbidden):**

```json
{
  "success": false,
  "message": "Admin access required",
  "errors": null
}
```

---

## 9. Create Zone — Validation Error (Invalid Type)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Invalid Zone",
  "type": "invalid_type",
  "total_capacity": 10,
  "price_per_hour": 2.0
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "errors": [
      {
        "field": "Type",
        "message": "Type must be one of: general, ev_charging, covered"
      }
    ]
  }
}
```

---

## 10. Create Zone — Validation Error (Negative Capacity)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Bad Capacity Zone",
  "type": "general",
  "total_capacity": -5,
  "price_per_hour": 2.0
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "errors": [
      {
        "field": "Total Capacity",
        "message": "Total Capacity must be greater than 0"
      }
    ]
  }
}
```

---

## Zone Types Reference

| Type          | Description                        |
| ------------- | ---------------------------------- |
| `general`     | Standard parking spots             |
| `ev_charging` | Electric vehicle charging stations |
| `covered`     | Covered/indoor parking             |

---

## Environment Variables

| Variable       | Initial Value                  | Description                               |
| -------------- | ------------------------------ | ----------------------------------------- |
| `base_url`     | `http://localhost:8080/api/v1` | API base URL                              |
| `admin_token`  | (empty)                        | Admin JWT token (set after admin login)   |
| `driver_token` | (empty)                        | Driver JWT token (set after driver login) |
