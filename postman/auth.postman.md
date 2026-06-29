# Spotsync API — Auth Endpoints (Postman Test Guide)

Base URL: `http://localhost:8080/api/v1`

---

## Response Format

### Success Response (direct DTO)

Success responses return the DTO object directly without a wrapper.

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "driver",
  "created_at": "2026-06-29T10:00:00+06:00",
  "updated_at": "2026-06-29T10:00:00+06:00"
}
```

### Error Response

Error responses use the standardized error format from `httpresponse.Error`.

```json
{
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

| Field     | Type     | Description                        |
| --------- | -------- | ---------------------------------- |
| `code`    | `int`    | HTTP status code                   |
| `message` | `string` | Human-readable error message       |
| `details` | `string` | (Optional) Technical error details |

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
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "driver",
  "created_at": "2026-06-29T10:00:00+06:00",
  "updated_at": "2026-06-29T10:00:00+06:00"
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
  "id": 2,
  "name": "Admin User",
  "email": "admin@example.com",
  "role": "admin",
  "created_at": "2026-06-29T10:01:00+06:00",
  "updated_at": "2026-06-29T10:01:00+06:00"
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
  "id": 3,
  "name": "Simple User",
  "email": "simple@example.com",
  "role": "driver",
  "created_at": "2026-06-29T10:02:00+06:00",
  "updated_at": "2026-06-29T10:02:00+06:00"
}
```

> 💡 **Note:** `role` field is optional. If omitted, defaults to `"driver"`.

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
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuYW1lIjoiSm9obiBEb2UiLCJlbWFpbCI6ImpvaG5AZXhhbXBsZS5jb20iLCJyb2xlIjoiZHJpdmVyIiwidG9rZW5fdHlwZSI6ImFjY2VzcyIsImV4cCI6MTc1MTI1MDQwMCwiaWF0IjoxNzUxMTY0MDAwLCJpc3MiOiJzcG90c3luYyJ9.xxx...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver"
  }
}
```

> 💡 **Tip:** Copy the `token` value and use it in subsequent requests as `Authorization: Bearer <token>`.

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
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJuYW1lIjoiQWRtaW4gVXNlciIsImVtYWlsIjoiYWRtaW5AZXhhbXBsZS5jb20iLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNzUxMjUwNDAwLCJpYXQiOjE3NTExNjQwMDAsImlzcyI6InNwb3RzeW5jIn0.xxx...",
  "user": {
    "id": 2,
    "name": "Admin User",
    "email": "admin@example.com",
    "role": "admin"
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
  "password": "123456",
  "role": "driver"
}
```

- **Expected Response (400 Bad Request):**

```json
{
  "code": 400,
  "message": "Email already registered"
}
```

---

## 7. Register — Validation Error (Missing Fields)

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
  "code": 400,
  "message": "Validation failed",
  "details": "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
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
  "code": 401,
  "message": "Invalid email or password"
}
```

---

## 9. Login — User Not Found (Error)

- **Method:** `POST`
- **URL:** `{{base_url}}/auth/login`
- **Headers:**
  - `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "email": "nonexistent@example.com",
  "password": "123456"
}
```

- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Invalid email or password"
}
```

---

## 10. Get Me (Protected Route)

- **Method:** `GET`
- **URL:** `{{base_url}}/auth/me`
- **Headers:**
  - `Authorization: Bearer {{driver_token}}`
- **Expected Response (200 OK):**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "driver",
  "created_at": "2026-06-29T10:00:00+06:00",
  "updated_at": "2026-06-29T10:00:00+06:00"
}
```

---

## 11. Get Me — Without Token (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/auth/me`
- **Headers:** None
- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Missing authorization header"
}
```

---

## 12. Get Me — Invalid Token (Error)

- **Method:** `GET`
- **URL:** `{{base_url}}/auth/me`
- **Headers:**
  - `Authorization: Bearer invalidtoken123`
- **Expected Response (401 Unauthorized):**

```json
{
  "code": 401,
  "message": "Invalid or expired token"
}
```

---

## Postman Test Script (Optional)

Login request এর **Tests** tab এ এই script যোগ করলে token automatically environment variable এ save হবে:

```javascript
// For Driver Login
const jsonData = pm.response.json();
pm.environment.set("driver_token", jsonData.token);

// For Admin Login (use this instead)
// const jsonData = pm.response.json();
// pm.environment.set("admin_token", jsonData.token);
```

---

## Environment Variables

Create a Postman environment with these variables:

| Variable       | Initial Value                  | Description                    |
| -------------- | ------------------------------ | ------------------------------ |
| `base_url`     | `http://localhost:8080/api/v1` | API base URL                   |
| `driver_token` | (empty)                        | Will be set after driver login |
| `admin_token`  | (empty)                        | Will be set after admin login  |
