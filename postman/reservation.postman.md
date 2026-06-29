# Spotsync API — Reservation Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## Response Format

### Success Response (direct DTO)

Success responses return the DTO object directly without a wrapper.

**Single Reservation:**

```json
{
  "id": 1,
  "user_id": 1,
  "zone_id": 1,
  "license_plate": "ABC-1234",
  "status": "active",
  "created_at": "2026-06-29T15:30:00+06:00",
  "updated_at": "2026-06-29T15:30:00+06:00"
}
```

**My Reservations (with zone info):**

```json
[
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
  }
]
```

### Error Response

```json
{
  "code": 409,
  "message": "Parking zone is full"
}
```

| Field     | Type     | Description                        |
| --------- | -------- | ---------------------------------- |
| `code`    | `int`    | HTTP status code                   |
| `message` | `string` | Human-readable error message       |
| `details` | `string` | (Optional) Technical error details |

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
  "id": 1,
  "user_id": 1,
  "zone_id": 1,
  "license_plate": "ABC-1234",
  "status": "active",
  "created_at": "2026-06-29T15:30:00+06:00",
  "updated_at": "2026-06-29T15:30:00+06:00"
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
  "id": 2,
  "user_id": 1,
  "zone_id": 2,
  "license_plate": "XYZ-9999",
  "status": "active",
  "created_at": "2026-06-29T15:31:00+06:00",
  "updated_at": "2026-06-29T15:31:00+06:00"
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
  "code": 401,
  "message": "Missing authorization header"
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
  "code": 404,
  "message": "Parking zone not found"
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
  "code": 409,
  "message": "Parking zone is full"
}
```

> 💡 **Note:** This error occurs when `active_reservations >= total_capacity` for the zone.

---

## 6. Create Reservation — Validation Error (Missing Fields)

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
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'CreateReservationRequest.ZoneID' Error:Field validation for 'ZoneID' failed on the 'required' tag"
}
```

---

## 7. Create Reservation — License Plate Too Long (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Content-Type: application/json`
  - `Authorization: Bearer {{driver_token}}`
- **Body (raw JSON):**

```json
{
  "zone_id": 1,
  "license_plate": "THIS-PLATE-IS-TOO-LONG-123"
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'CreateReservationRequest.LicensePlate' Error:Field validation for 'LicensePlate' failed on the 'max' tag"
}
```

---

## 8. Get My Reservations (Driver)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
[
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
```

---

## 9. Get My Reservations — Empty List

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
[]
```

> 💡 **Note:** Returns empty array when user has no reservations.

---

## 10. Get My Reservations — Without Auth (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations/my-reservations`
- **Headers:** None
- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Missing authorization header"
}
```

---

## 11. Cancel Reservation (Driver — Own Reservation)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/1`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "message": "Reservation cancelled successfully"
}
```

---

## 12. Cancel Reservation — Not Found (Error)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/9999`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (404 Not Found):**

```json
{
  "code": 404,
  "message": "Reservation not found"
}
```

---

## 13. Cancel Reservation — Forbidden (Not Owner)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/2`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (403 Forbidden):**

```json
{
  "code": 403,
  "message": "You can only cancel your own reservations"
}
```

> 💡 **Note:** This error occurs when a driver tries to cancel another user's reservation.

---

## 14. Cancel Reservation — Without Auth (Error)

- **Method:** `DELETE`
- **URL:** `{{base_url}}/reservations/1`
- **Headers:** None
- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Missing authorization header"
}
```

---

## 15. Get All Reservations (Admin Only)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Authorization: Bearer {{admin_token}}`
- **Expected Response (200 OK):**

```json
[
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
```

---

## 16. Get All Reservations — Driver Token (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (403 Forbidden):**

```json
{
  "code": 403,
  "message": "Admin access required"
}
```

---

## 17. Get All Reservations — Without Auth (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/reservations`
- **Headers:** None
- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Missing authorization header"
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
pm.environment.set("reservation_id", jsonData.id);
```

Cancel করার সময় URL এ `{{reservation_id}}` ব্যবহার করুন।

---

## Environment Variables

| Variable         | Initial Value                  | Description                                     |
| ---------------- | ------------------------------ | ----------------------------------------------- |
| `base_url`       | `http://localhost:8080/api/v1` | API base URL                                    |
| `driver_token`   | (empty)                        | Driver JWT token (set after driver login)       |
| `admin_token`    | (empty)                        | Admin JWT token (set after admin login)         |
| `reservation_id` | (empty)                        | Reservation ID (set after creating reservation) |
