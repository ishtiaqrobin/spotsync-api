# Spotsync API — Auth Endpoints (Postman Test Guide)

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

## 1. Register (Driver)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/register`
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

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver",
    "created_at": "2026-06-29T10:00:00+06:00",
    "updated_at": "2026-06-29T10:00:00+06:00"
  }
}
```

---

## 2. Register (Admin)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/register`
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

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 2,
    "name": "Admin User",
    "email": "admin@example.com",
    "role": "admin",
    "created_at": "2026-06-29T10:01:00+06:00",
    "updated_at": "2026-06-29T10:01:00+06:00"
  }
}
```

---

## 3. Register — Without Role (Defaults to Driver)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/register`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "Simple User",
  "email": "simple@example.com",
  "password": "123456"
}
```

- **Expected Response (201 Created):**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 3,
    "name": "Simple User",
    "email": "simple@example.com",
    "role": "driver",
    "created_at": "2026-06-29T10:02:00+06:00",
    "updated_at": "2026-06-29T10:02:00+06:00"
  }
}
```

---

## 4. Login (Driver)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "john@example.com",
  "password": "123456"
}
```

- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "driver"
    }
  }
}
```

---

## 5. Login (Admin)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 2,
      "name": "Admin User",
      "email": "admin@example.com",
      "role": "admin"
    }
  }
}
```

---

## 6. Register — Duplicate Email (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/register`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "John Doe 2",
  "email": "john@example.com",
  "password": "123456"
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Email already registered",
  "errors": null
}
```

---

## 7. Register — Validation Error (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/register`
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

- **Expected Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "errors": [
      {
        "field": "Name",
        "message": "Name is required"
      },
      {
        "field": "Email",
        "message": "Email must be a valid email address"
      },
      {
        "field": "Password",
        "message": "Password must be at least 6 characters"
      }
    ]
  }
}
```

---

## 8. Login — Invalid Credentials (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "john@example.com",
  "password": "wrongpassword"
}
```

- **Expected Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Invalid email or password",
  "errors": null
}
```

---

## 9. Get Me (Protected Route)

- **Method:** `GET`
- **URL:** `{{base_url}}/auth/me`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver",
    "created_at": "2026-06-29T10:00:00+06:00",
    "updated_at": "2026-06-29T10:00:00+06:00"
  }
}
```

---

## 10. Get Me — Without Token (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/auth/me`
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

## Postman Test Script (Optional)

Login request এর **Tests** tab এ এই script যোগ করলে token automatically environment variable এ save হবে:

```javascript
const jsonData = pm.response.json();
pm.environment.set("driver_token", jsonData.data.token);
```

---

## Environment Variables

| Variable       | Initial Value                  | Description                    |
| -------------- | ------------------------------ | ------------------------------ |
| `base_url`     | `http://localhost:8080/api/v1` | API base URL                   |
| `driver_token` | (empty)                        | Will be set after driver login |
| `admin_token`  | (empty)                        | Will be set after admin login  |
