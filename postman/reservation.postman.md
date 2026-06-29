# Spotsync API — Reservation Endpoints (Postman Test Guide)

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

## 1. Create Reservation (Driver)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 1,
  "license_plate": "ABC-1234"
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "Reservation confirmed successfully",
  "data": {
    "id": 1,
    "user_id": 1,
    "zone_id": 1,
    "license_plate": "ABC-1234",
    "status": "active",
    "created_at": "2026-06-29T15:30:00+06:00",
    "updated_at": "2026-06-29T15:30:00+06:00"
  }
}
```

---

## 2. Create Reservation — EV Zone

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 2,
  "license_plate": "XYZ-9999"
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "Reservation confirmed successfully",
  "data": {
    "id": 2,
    "user_id": 1,
    "zone_id": 2,
    "license_plate": "XYZ-9999",
    "status": "active",
    "created_at": "2026-06-29T15:31:00+06:00",
    "updated_at": "2026-06-29T15:31:00+06:00"
  }
}
```

---

## 3. Create Reservation — Without Auth (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "zone_id": 1,
  "license_plate": "NO-AUTH-1"
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

## 4. Create Reservation — Invalid Zone ID (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 9999,
  "license_plate": "ABC-1234"
}
```

- **Expected Response (404 Not Found):**

```json
{
  "success": false,
  "message": "Parking zone not found",
  "errors": null
}
```

---

## 5. Create Reservation — Zone Full (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 1,
  "license_plate": "FULL-0001"
}
```

- **Expected Response (409 Conflict):**

```json
{
  "success": false,
  "message": "Parking zone is full",
  "errors": null
}
```

---

## 6. Create Reservation — Validation Error (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 0,
  "license_plate": ""
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
        "field": "Zone ID",
        "message": "Zone ID is required"
      },
      {
        "field": "License Plate",
        "message": "License Plate is required"
      }
    ]
  }
}
```

---

## 7. Get My Reservations (Driver)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "My reservations retrieved successfully",
  "data": [
    {
      "id": 1,
      "license_plate": "ABC-1234",
      "status": "active",
      "zone": {
        "id": 1,
        "name": "Terminal 1 EV Charging",
        "type": "ev_charging"
      },
      "created_at": "2026-06-29T15:30:00+06:00"
    },
    {
      "id": 2,
      "license_plate": "XYZ-9999",
      "status": "active",
      "zone": {
        "id": 2,
        "name": "Downtown Parking",
        "type": "general"
      },
      "created_at": "2026-06-29T15:31:00+06:00"
    }
  ]
}
```

---

## 8. Get My Reservations — Empty List

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "My reservations retrieved successfully",
  "data": []
}
```

---

## 9. Get My Reservations — Without Auth (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:** None
- **Expected Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Missing authorization header",
  "errors": null
}
```

---

## 10. Cancel Reservation (Driver — Own Reservation)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/1`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "Reservation cancelled successfully",
  "data": null
}
```

---

## 11. Cancel Reservation — Not Found (Error)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/9999`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (404 Not Found):**

```json
{
  "success": false,
  "message": "Reservation not found",
  "errors": null
}
```

---

## 12. Cancel Reservation — Forbidden (Not Owner)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/2`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (403 Forbidden):**

```json
{
  "success": false,
  "message": "You can only cancel your own reservations",
  "errors": null
}
```

---

## 13. Get All Reservations (Admin Only)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Authorization: Bearer {{admin_token}}`
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "All reservations retrieved successfully",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "zone_id": 1,
      "license_plate": "ABC-1234",
      "status": "cancelled",
      "created_at": "2026-06-29T15:30:00+06:00",
      "updated_at": "2026-06-29T15:35:00+06:00"
    },
    {
      "id": 2,
      "user_id": 1,
      "zone_id": 2,
      "license_plate": "XYZ-9999",
      "status": "active",
      "created_at": "2026-06-29T15:31:00+06:00",
      "updated_at": "2026-06-29T15:31:00+06:00"
    }
  ]
}
```

---

## 14. Get All Reservations — Driver Token (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (403 Forbidden):**

```json
{
  "success": false,
  "message": "Admin access required",
  "errors": null
}
```

---

## Reservation Status Reference

| Status      | Description                                   |
| ----------- | --------------------------------------------- |
| `active`    | Currently active reservation                  |
| `completed` | Reservation completed (parking session ended) |
| `cancelled` | Reservation cancelled by user or admin        |

---

## Postman Test Script (Optional)

Create Reservation এর **Tests** tab এ এই script যোগ করলে reservation id automatically save হবে:

```javascript
const jsonData = pm.response.json();
pm.environment.set("reservation_id", jsonData.data.id);
```

---

## Environment Variables

| Variable         | Initial Value                  | Description                                     |
| ---------------- | ------------------------------ | ----------------------------------------------- |
| `base_url`       | `http://localhost:8080/api/v1` | API base URL                                    |
| `driver_token`   | (empty)                        | Driver JWT token (set after driver login)       |
| `admin_token`    | (empty)                        | Admin JWT token (set after admin login)         |
| `reservation_id` | (empty)                        | Reservation ID (set after creating reservation) |
