# Spotsync API — Auth Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## 1. Register (Driver)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/register`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456",
  "role": "driver"
}
```

- **Expected Response (201):**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "driver",
  "created_at": "2026-06-29T10:00:00Z",
  "updated_at": "2026-06-29T10:00:00Z"
}
```

---

## 2. Register (Admin)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/register`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "Admin User",
  "email": "admin@example.com",
  "password": "admin123",
  "role": "admin"
}
```

- **Expected Response (201):**

```json
{
  "id": 2,
  "name": "Admin User",
  "email": "admin@example.com",
  "role": "admin",
  "created_at": "2026-06-29T10:01:00Z",
  "updated_at": "2026-06-29T10:01:00Z"
}
```

---

## 3. Login (Driver)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "john@example.com",
  "password": "123456"
}
```

- **Expected Response (200):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver"
  }
}
```

> 💡 **Tip:** Login response থেকে `token` কপি করে নিচের protected requests এ `Authorization` header এ ব্যবহার করুন।

---

## 4. Login (Admin)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

- **Expected Response (200):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 2,
    "name": "Admin User",
    "email": "admin@example.com",
    "role": "admin"
  }
}
```

> 💡 **Tip:** Admin token দিয়ে শুধুমাত্র admin-only endpoints (Create Zone, Get All Reservations) access করা যাবে।

---

## 5. Register — Validation Error (Missing Fields)

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/register`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "",
  "email": "invalid-email",
  "password": "123"
}
```

- **Expected Response (400/422):** Validation error message

---

## 6. Login — Invalid Credentials

- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "john@example.com",
  "password": "wrongpassword"
}
```

- **Expected Response (401):** Invalid credentials error

---

## Postman Test Script (Optional)

Login request এর **Tests** tab এ এই script যোগ করলে token automatically environment variable এ save হবে:

```javascript
const jsonData = pm.response.json();
pm.environment.set("driver_token", jsonData.token);
```

Admin login এর জন্য:

```javascript
const jsonData = pm.response.json();
pm.environment.set("admin_token", jsonData.token);
```
