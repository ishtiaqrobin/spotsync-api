# Spotsync API — Reservation Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## 1. Create Reservation (Driver)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/reservations`
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

- **Expected Response (201):**

```json
{
  "id": 1,
  "user_id": 1,
  "zone_id": 1,
  "license_plate": "ABC-1234",
  "status": "active",
  "created_at": "2026-06-29T10:00:00Z",
  "updated_at": "2026-06-29T10:00:00Z"
}
```

---

## 2. Create Reservation — EV Zone

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/reservations`
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

- **Expected Response (201):**

```json
{
  "id": 2,
  "user_id": 1,
  "zone_id": 2,
  "license_plate": "XYZ-9999",
  "status": "active",
  "created_at": "2026-06-29T10:01:00Z",
  "updated_at": "2026-06-29T10:01:00Z"
}
```

---

## 3. Create Reservation — Without Auth (Should Fail)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/reservations`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "zone_id": 1,
  "license_plate": "NO-AUTH-1"
}
```

- **Expected Response (401):** Unauthorized — token required

---

## 4. Create Reservation — Invalid Zone ID

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/reservations`
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

- **Expected Response (404):** Zone not found

---

## 5. Create Reservation — Validation Error (Missing Fields)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/reservations`
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

- **Expected Response (400/422):** Validation error

---

## 6. Get My Reservations (Driver)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/reservations/my-reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200):**

```json
[
  {
    "id": 1,
    "license_plate": "ABC-1234",
    "status": "active",
    "zone": {
      "id": 1,
      "name": "Downtown Parking",
      "type": "general"
    },
    "created_at": "2026-06-29T10:00:00Z"
  },
  {
    "id": 2,
    "license_plate": "XYZ-9999",
    "status": "active",
    "zone": {
      "id": 2,
      "name": "EV Station A",
      "type": "ev_charging"
    },
    "created_at": "2026-06-29T10:01:00Z"
  }
]
```

---

## 7. Get My Reservations — Without Auth (Should Fail)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/reservations/my-reservations`
- **Headers:** None
- **Expected Response (401):** Unauthorized — token required

---

## 8. Cancel Reservation (Driver)

- **Method:** `DELETE`
- **URL:** `http://localhost:8080/api/v1/reservations/1`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200):**

```json
{
  "message": "reservation cancelled successfully"
}
```

---

## 9. Cancel Reservation — Not Found

- **Method:** `DELETE`
- **URL:** `http://localhost:8080/api/v1/reservations/9999`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (404):** Reservation not found

---

## 10. Cancel Reservation — Without Auth (Should Fail)

- **Method:** `DELETE`
- **URL:** `http://localhost:8080/api/v1/reservations/1`
- **Headers:** None
- **Expected Response (401):** Unauthorized — token required

---

## 11. Get All Reservations (Admin Only)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/reservations`
- **Headers:**
  - `Authorization: Bearer {{admin_token}}`
- **Expected Response (200):**

```json
[
  {
    "id": 1,
    "user_id": 1,
    "zone_id": 1,
    "license_plate": "ABC-1234",
    "status": "cancelled",
    "created_at": "2026-06-29T10:00:00Z",
    "updated_at": "2026-06-29T10:05:00Z"
  },
  {
    "id": 2,
    "user_id": 1,
    "zone_id": 2,
    "license_plate": "XYZ-9999",
    "status": "active",
    "created_at": "2026-06-29T10:01:00Z",
    "updated_at": "2026-06-29T10:01:00Z"
  }
]
```

---

## 12. Get All Reservations — Driver Token (Should Fail)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/reservations`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (403):** Forbidden — admin access required

---

## 13. Get All Reservations — Without Auth (Should Fail)

- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/reservations`
- **Headers:** None
- **Expected Response (401):** Unauthorized — token required

---

## Postman Test Script (Optional)

Create Reservation এর **Tests** tab এ এই script যোগ করলে reservation id automatically save হবে:

```javascript
const jsonData = pm.response.json();
pm.environment.set("reservation_id", jsonData.id);
```

Cancel করার সময় URL এ `{{reservation_id}}` ব্যবহার করুন।
