# Spotsync API — Zone Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## Response Format

### Success Response (direct DTO)

Success responses return the DTO object directly without a wrapper.

**Single Zone:**

```json
{
  "id": 1,
  "name": "Terminal 1 EV Charging",
  "type": "ev_charging",
  "total_capacity": 20,
  "available_spots": 14,
  "price_per_hour": 5.5,
  "created_at": "2026-06-29T10:30:00+06:00",
  "updated_at": "2026-06-29T10:30:00+06:00"
}
```

**Array of Zones:**

```json
[
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
```

### Error Response

```json
{
  "code": 404,
  "message": "Parking zone not found"
}
```

| Field     | Type     | Description                        |
| --------- | -------- | ---------------------------------- |
| `code`    | `int`    | HTTP status code                   |
| `message` | `string` | Human-readable error message       |
| `details` | `string` | (Optional) Technical error details |

---

## 1. Get All Zones (Public)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones`
- **Headers:** None required
- **Expected Response (200 OK):**

```json
[
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
  },
  {
    "id": 3,
    "name": "Mall Covered Parking",
    "type": "covered",
    "total_capacity": 30,
    "available_spots": 30,
    "price_per_hour": 3.0,
    "created_at": "2026-06-29T09:15:00+06:00",
    "updated_at": "2026-06-29T09:15:00+06:00"
  }
]
```

> 💡 **Note:** `available_spots` is calculated dynamically: `total_capacity - active_reservations`.

---

## 2. Get Zone By ID (Public)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones/1`
- **Headers:** None required
- **Expected Response (200 OK):**

```json
{
  "id": 1,
  "name": "Terminal 1 EV Charging",
  "type": "ev_charging",
  "total_capacity": 20,
  "available_spots": 14,
  "price_per_hour": 5.5,
  "created_at": "2026-06-29T10:30:00+06:00",
  "updated_at": "2026-06-29T10:30:00+06:00"
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
  "code": 404,
  "message": "Parking zone not found"
}
```

---

## 4. Get Zone By ID — Invalid ID (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/zones/abc`
- **Headers:** None required
- **Expected Response (400 Bad Request):**

```json
{
  "code": 400,
  "message": "Invalid zone id",
  "details": "strconv.Atoi: parsing \"abc\": invalid syntax"
}
```

---

## 5. Create Zone — General Type (Admin Only)

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
  "id": 1,
  "name": "Downtown Parking",
  "type": "general",
  "total_capacity": 50,
  "available_spots": 50,
  "price_per_hour": 2.5,
  "created_at": "2026-06-29T10:30:00+06:00",
  "updated_at": "2026-06-29T10:30:00+06:00"
}
```

> 💡 **Note:** New zone has `available_spots` = `total_capacity` (no reservations yet).

---

## 6. Create Zone — EV Charging Type (Admin Only)

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
  "id": 2,
  "name": "Terminal 1 EV Charging",
  "type": "ev_charging",
  "total_capacity": 20,
  "available_spots": 20,
  "price_per_hour": 5.5,
  "created_at": "2026-06-29T10:31:00+06:00",
  "updated_at": "2026-06-29T10:31:00+06:00"
}
```

---

## 7. Create Zone — Covered Type (Admin Only)

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
  "id": 3,
  "name": "Mall Covered Parking",
  "type": "covered",
  "total_capacity": 30,
  "available_spots": 30,
  "price_per_hour": 3.0,
  "created_at": "2026-06-29T10:32:00+06:00",
  "updated_at": "2026-06-29T10:32:00+06:00"
}
```

---

## 8. Create Zone — Without Auth (Error)

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
  "code": 401,
  "message": "Missing authorization header"
}
```

---

## 9. Create Zone — Driver Token (Error, Admin Only)

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
  "code": 403,
  "message": "Admin access required"
}
```

---

## 10. Create Zone — Validation Error (Invalid Type)

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
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'CreateZoneRequest.Type' Error:Field validation for 'Type' failed on the 'oneof' tag"
}
```

---

## 11. Create Zone — Validation Error (Negative Capacity)

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
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'CreateZoneRequest.TotalCapacity' Error:Field validation for 'TotalCapacity' failed on the 'gt' tag"
}
```

---

## 12. Create Zone — Validation Error (Missing Name)

- **Method:** `POST`
- **URL:** `{{base_url}}/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "",
  "type": "general",
  "total_capacity": 10,
  "price_per_hour": 2.0
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'CreateZoneRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
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
