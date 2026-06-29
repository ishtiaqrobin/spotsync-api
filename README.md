# 🚗 SpotSync – Smart Parking & EV Charging Reservation API

> A centralized platform for busy airports and malls to manage parking zones, specifically handling the high-demand reservation of limited EV charging spots.

---

## 🌐 Live URL

**Backend:** `https://spotsync-api.onrender.com` _(Update after deployment)_

---

## 🛠️ Technology Stack

| Technology     | Package                                  | Purpose                                    |
| -------------- | ---------------------------------------- | ------------------------------------------ |
| **Go**         | `go 1.22+`                               | Programming language                       |
| **Echo**       | `github.com/labstack/echo/v4`            | High performance, minimalist web framework |
| **GORM**       | `gorm.io/gorm`                           | ORM for Go                                 |
| **PostgreSQL** | `gorm.io/driver/postgres`                | Relational database (NeonDB)               |
| **Validator**  | `github.com/go-playground/validator/v10` | Struct validation, integrated with Echo    |
| **JWT**        | `github.com/golang-jwt/jwt/v5`           | Standard token generation & verification   |
| **bcrypt**     | `golang.org/x/crypto/bcrypt`             | Password hashing                           |
| **Godotenv**   | `github.com/joho/godotenv`               | Environment variable management            |
| **Air**        | `github.com/air-verse/air`               | Hot reloading during development           |

---

## ✨ Features

### Authentication & Authorization

- User registration (driver/admin roles)
- JWT-based authentication
- Role-based access control (driver, admin)
- Protected routes with middleware

### Parking Zones Management

- Create parking zones (admin only)
- View all zones with dynamic available spots calculation
- Support for zone types: `general`, `ev_charging`, `covered`
- Dynamic pricing per hour

### Reservations System

- Reserve parking spots with concurrency-safe transactions
- Row-level locking (`FOR UPDATE`) to prevent over-capacity
- View personal reservations with zone details
- Cancel own reservations (driver)
- View all reservations (admin)

---

## 🏛️ Project Structure (Domain-Driven Clean Architecture)

```
spotsync-api/
├── cmd/
│   └── main.go                    # Application entry point
├── internal/
│   ├── auth/                      # JWT service interface + implementation
│   │   └── jwt.go
│   ├── config/                    # Configuration & database connection
│   │   ├── config.go
│   │   └── db.go
│   ├── domain/                    # Domain modules (each with own layers)
│   │   ├── user/                  # User/Auth domain
│   │   │   ├── dto/
│   │   │   │   ├── request.go
│   │   │   │   └── response.go
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   └── register.go        # Route registration + DI
│   │   ├── zone/                  # Parking Zone domain
│   │   │   ├── dto/
│   │   │   │   ├── request.go
│   │   │   │   └── response.go
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   └── register.go
│   │   └── reservation/           # Reservation domain
│   │       ├── dto/
│   │       │   ├── request.go
│   │       │   └── response.go
│   │       ├── entity.go
│   │       ├── repository.go
│   │       ├── service.go
│   │       ├── handler.go
│   │       └── register.go
│   ├── httpresponse/              # Standardized response helpers
│   │   ├── response.go
│   │   └── error.go
│   ├── middleware/                # Auth + Role middleware
│   │   └── auth.go
│   ├── server/                    # HTTP server setup
│   │   └── http.go
│   └── validation/                # Validation error parsing
│       └── error.go
├── postman/                       # Postman test guides
│   ├── auth.postman.md
│   ├── zone.postman.md
│   └── reservation.postman.md
├── doc/                           # Documentation
│   └── deploy.md
├── ref-golang/                    # Reference project (for learning)
├── .env                           # Environment variables (not in git)
├── .env.example                   # Example environment file
├── .gitignore
├── .air.toml                      # Air hot-reload config
├── CONCEPTS.md                    # Project concepts & keywords
├── go.mod
├── go.sum
└── README.md
```

---

## 🏗️ Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        HTTP Request                             │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Echo Server (server/http.go)                 │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Middleware: Logger → Recover → CORS → JWTAuth → Role    │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Handler Layer                               │
│ (Bind request → Validate DTO → Extract JWT claims → Call Service│
│  → Return JSON response)                                        │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Service Layer                              │
│  (Business logic: Hash passwords, Generate JWT, Check capacity, │
│   Enforce rules)                                                │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Repository Layer                             │
│  (Database operations: CRUD, Transactions, Row Locks)           │
└─────────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                      PostgreSQL (NeonDB)                        │
└─────────────────────────────────────────────────────────────────┘
```

### Dependency Injection Flow

Each domain wires its own dependencies in `register.go`:

```
Repository → Service → Handler → Routes
```

---

## 🔐 User Roles & Permissions

| Role       | Allowed Actions                                                                                                                               |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| **driver** | • Register and log in<br>• View all parking zones and availability<br>• Reserve a parking/EV spot<br>• View and cancel their own reservations |
| **admin**  | • All driver permissions<br>• Create parking zones<br>• Set pricing for zones<br>• View all reservations in the system                        |

---

## 🗄️ Database Schema

### Table: `users`

| Field        | Type           | Constraints                 |
| ------------ | -------------- | --------------------------- |
| `id`         | `SERIAL`       | Primary Key, Auto-increment |
| `name`       | `VARCHAR(100)` | NOT NULL                    |
| `email`      | `VARCHAR(255)` | UNIQUE, NOT NULL            |
| `password`   | `VARCHAR(255)` | NOT NULL (bcrypt hash)      |
| `role`       | `VARCHAR(10)`  | DEFAULT 'driver'            |
| `created_at` | `TIMESTAMP`    | Auto-generated              |
| `updated_at` | `TIMESTAMP`    | Auto-refreshed              |

### Table: `parking_zones`

| Field            | Type           | Constraints                              |
| ---------------- | -------------- | ---------------------------------------- |
| `id`             | `SERIAL`       | Primary Key, Auto-increment              |
| `name`           | `VARCHAR(100)` | NOT NULL                                 |
| `type`           | `VARCHAR(20)`  | NOT NULL (general, ev_charging, covered) |
| `total_capacity` | `INTEGER`      | NOT NULL, > 0                            |
| `price_per_hour` | `DECIMAL`      | NOT NULL, > 0                            |
| `created_at`     | `TIMESTAMP`    | Auto-generated                           |
| `updated_at`     | `TIMESTAMP`    | Auto-refreshed                           |

### Table: `reservations`

| Field           | Type          | Constraints                                     |
| --------------- | ------------- | ----------------------------------------------- |
| `id`            | `SERIAL`      | Primary Key, Auto-increment                     |
| `user_id`       | `INTEGER`     | Foreign Key → users.id                          |
| `zone_id`       | `INTEGER`     | Foreign Key → parking_zones.id                  |
| `license_plate` | `VARCHAR(15)` | NOT NULL                                        |
| `status`        | `VARCHAR(15)` | DEFAULT 'active' (active, completed, cancelled) |
| `created_at`    | `TIMESTAMP`   | Auto-generated                                  |
| `updated_at`    | `TIMESTAMP`   | Auto-refreshed                                  |

---

## 🌐 API Endpoints

### Base URL

```
http://localhost:8080/api/v1
```

### Health Check

| Method | URL       | Access |
| ------ | --------- | ------ |
| `GET`  | `/health` | Public |

---

### 🔹 Authentication Endpoints

| Method | URL                     | Access        | Description           |
| ------ | ----------------------- | ------------- | --------------------- |
| `POST` | `/api/v1/auth/register` | Public        | Register new user     |
| `POST` | `/api/v1/auth/login`    | Public        | Login user            |
| `GET`  | `/api/v1/auth/me`       | Authenticated | Get current user info |

---

### 🔹 Parking Zone Endpoints

| Method | URL                 | Access | Description     |
| ------ | ------------------- | ------ | --------------- |
| `GET`  | `/api/v1/zones`     | Public | Get all zones   |
| `GET`  | `/api/v1/zones/:id` | Public | Get zone by ID  |
| `POST` | `/api/v1/zones`     | Admin  | Create new zone |

---

### 🔹 Reservation Endpoints

| Method   | URL                                    | Access        | Description          |
| -------- | -------------------------------------- | ------------- | -------------------- |
| `POST`   | `/api/v1/reservations`                 | Authenticated | Create reservation   |
| `GET`    | `/api/v1/reservations/my-reservations` | Authenticated | Get my reservations  |
| `DELETE` | `/api/v1/reservations/:id`             | Authenticated | Cancel reservation   |
| `GET`    | `/api/v1/reservations`                 | Admin         | Get all reservations |

---

## 📋 Response Format

### Success Response

All success responses follow a standardized wrapper format:

```json
{
  "success": true,
  "message": "Operation description",
  "data": { ... }
}
```

| Field     | Type           | Description                    |
| --------- | -------------- | ------------------------------ |
| `success` | `boolean`      | Always `true` for success      |
| `message` | `string`       | Human-readable success message |
| `data`    | `object/array` | Response data (DTO object)     |

**Example - Register User (201 Created):**

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

**Example - Login (200 OK):**

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

**Example - Get All Zones (200 OK):**

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
      "created_at": "2026-06-29T10:30:00+06:00"
    }
  ]
}
```

---

### Error Response

All error responses follow a standardized wrapper format:

```json
{
  "success": false,
  "message": "Error description",
  "errors": "Error details or validation errors"
}
```

| Field     | Type                 | Description                                                                                   |
| --------- | -------------------- | --------------------------------------------------------------------------------------------- |
| `success` | `boolean`            | Always `false` for errors                                                                     |
| `message` | `string`             | Human-readable error message                                                                  |
| `errors`  | `string/object/null` | Error details (string for simple errors, object for validation errors, `null` for no details) |

**Example - Validation Error (400 Bad Request):**

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "errors": [
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

**Example - Authentication Error (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Missing authorization header",
  "errors": null
}
```

**Example - Forbidden Error (403 Forbidden):**

```json
{
  "success": false,
  "message": "Admin access required",
  "errors": null
}
```

**Example - Not Found Error (404 Not Found):**

```json
{
  "success": false,
  "message": "Parking zone not found",
  "errors": null
}
```

**Example - Conflict Error (409 Conflict):**

```json
{
  "success": false,
  "message": "Parking zone is full",
  "errors": null
}
```

---

### HTTP Status Codes

| Code  | Usage                                                |
| ----- | ---------------------------------------------------- |
| `200` | Successful GET, DELETE                               |
| `201` | Successful POST (resource created)                   |
| `400` | Validation errors, invalid input, duplicate resource |
| `401` | Missing, expired, or invalid JWT token               |
| `403` | Valid token but insufficient role/permissions        |
| `404` | Requested resource does not exist                    |
| `409` | Business logic conflict (e.g., Zone is full)         |
| `500` | Unexpected server or database error                  |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher
- PostgreSQL (NeonDB, Supabase, or local)
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/ishtiaqrobin/spotsync-api.git
cd spotsync-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment Variables

Create a `.env` file in the root directory:

```env
# Option 1: Direct DSN (recommended for NeonDB)
DSN=postgresql://user:password@host.neon.tech/dbname?sslmode=require

# Option 2: Individual DB components
DB_HOST=your-neondb-host.neon.tech
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=spotsync
DB_SSLMODE=require

# JWT Secret (generate with: node -e "console.log(require('crypto').randomBytes(32).toString('hex'))")
JWT_SECRET=your-super-secret-key

# Server Port
PORT=8080
```

### 4. Run the Application

**Without hot-reload:**

```bash
go run ./cmd/main.go
```

**With hot-reload (Air):**

```bash
# Install air if not already installed
go install github.com/air-verse/air@latest

# Run with air
air
```

### 5. Verify Installation

Visit `http://localhost:8080/health` — you should see:

```json
{
  "status": "ok"
}
```

---

## 🧪 Testing with Postman

Detailed Postman test guides are available in the `postman/` directory:

- [`postman/auth.postman.md`](postman/auth.postman.md) — Authentication endpoints
- [`postman/zone.postman.md`](postman/zone.postman.md) — Parking zone endpoints
- [`postman/reservation.postman.md`](postman/reservation.postman.md) — Reservation endpoints

### Postman Environment Variables

| Variable       | Description                    |
| -------------- | ------------------------------ |
| `base_url`     | `http://localhost:8080/api/v1` |
| `driver_token` | Set after driver login         |
| `admin_token`  | Set after admin login          |

---

## 🔒 Security Features

- **Password Hashing:** bcrypt with default cost (10)
- **JWT Authentication:** HS256 signed tokens with 24-hour expiry
- **Role-Based Access Control:** Middleware enforces admin-only routes
- **Password Never Exposed:** `json:"-"` tag prevents password in responses
- **CORS Enabled:** Cross-origin requests supported

---

## ⚡ Concurrency Handling

The reservation system uses **GORM Transactions** with **Row-Level Locking** (`FOR UPDATE`) to prevent the "EV Spot Bottleneck" race condition:

```go
db.Transaction(func(tx *gorm.DB) error {
    // 1. Lock the zone row
    tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&zone, zoneID)

    // 2. Count active reservations
    tx.Model(&Reservation{}).Where("zone_id = ? AND status = ?", zoneID, "active").Count(&count)

    // 3. Check capacity
    if activeCount >= zone.TotalCapacity {
        return ErrZoneFull
    }

    // 4. Create reservation
    tx.Create(&reservation)
    return nil
})
```

---

## 📬 Deployment

### Deploy to Render/Railway/Fly.io

1. Push code to GitHub
2. Connect repository to Render/Railway/Fly.io
3. Set environment variables in dashboard
4. Deploy

### Database Setup (NeonDB)

1. Go to [neon.tech](https://neon.tech) and create a project
2. Copy the connection string
3. Set as `DSN` in your environment variables

---

## 📝 License

This project is built for educational purposes as part of the Level2-B6 Mission-9 Assignment.

---

## 👨‍💻 Author

**Ishtiaq Robin**

- GitHub: [@ishtiaqrobin](https://github.com/ishtiaqrobin)
- Project: [spotsync-api](https://github.com/ishtiaqrobin/spotsync-api)
- Live: [spotsync-api](https://spotsync-api-uzl2.onrender.com)

---

**Built using Go, Echo, and GORM**
