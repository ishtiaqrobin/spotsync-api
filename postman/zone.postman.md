# Spotsync API — Zone Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## 1. Get All Zones (Public)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/zones`
- **Headers:** None required
- **Expected Response (200):**

```json
[
  {
    "id": 1,
    "name": "Downtown Parking",
    "type": "general",
    "total_capacity": 50,
    "available_spots": 30,
    "price_per_hour": 2.5,
    "created_at": "2026-06-29T09:00:00Z"
  },
  {
    "id": 2,
    "name": "EV Station A",
    "type": "ev_charging",
    "total_capacity": 10,
    "available_spots": 5,
    "price_per_hour": 5.0,
    "created_at": "2026-06-29T09:00:00Z"
  }
]
```

---

## 2. Get Zone By ID (Public)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/zones/1`
- **Headers:** None required
- **Expected Response (200):**

```json
{
  "id": 1,
  "name": "Downtown Parking",
  "type": "general",
  "total_capacity": 50,
  "available_spots": 30,
  "price_per_hour": 2.5,
  "created_at": "2026-06-29T09:00:00Z"
}
```

---

## 3. Get Zone By ID — Not Found

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/zones/9999`
- **Headers:** None required
- **Expected Response (404):** Zone not found error

---

## 4. Create Zone (Admin Only)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Covered Parking B",
  "type": "covered",
  "total_capacity": 20,
  "price_per_hour": 3.0
}
```

- **Expected Response (201):**

```json
{
  "id": 3,
  "name": "Covered Parking B",
  "type": "covered",
  "total_capacity": 20,
  "available_spots": 20,
  "price_per_hour": 3.0,
  "created_at": "2026-06-29T10:00:00Z"
}
```

---

## 5. Create Zone — General Type

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "Mall Parking",
  "type": "general",
  "total_capacity": 100,
  "price_per_hour": 1.5
}
```

- **Expected Response (201):**

```json
{
  "id": 4,
  "name": "Mall Parking",
  "type": "general",
  "total_capacity": 100,
  "available_spots": 100,
  "price_per_hour": 1.5,
  "created_at": "2026-06-29T10:01:00Z"
}
```

---

## 6. Create Zone — EV Charging Type

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{admin_token}}`
- **Body (raw JSON):**

```json
{
  "name": "EV Hub Central",
  "type": "ev_charging",
  "total_capacity": 15,
  "price_per_hour": 6.0
}
```

- **Expected Response (201):**

```json
{
  "id": 5,
  "name": "EV Hub Central",
  "type": "ev_charging",
  "total_capacity": 15,
  "available_spots": 15,
  "price_per_hour": 6.0,
  "created_at": "2026-06-29T10:02:00Z"
}
```

---

## 7. Create Zone — Without Auth (Should Fail)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
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

- **Expected Response (401):** Unauthorized — token required

---

## 8. Create Zone — Driver Token (Should Fail, Admin Only)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
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

- **Expected Response (403):** Forbidden — admin access required

---

## 9. Create Zone — Validation Error (Invalid Type)

- **Method:** `POST`
- `URL:` `http://localhost:8080/api/v1/zones`
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

- **Expected Response (400/422):** Validation error — type must be one of: general, ev_charging, covered

---

## 10. Create Zone — Validation Error (Negative Capacity)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/zones`
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

- **Expected Response (400/422):** Validation error — total_capacity must be greater than 0
