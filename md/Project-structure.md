Project Structure for Golang Conceptual

```
haddiBanga/
├── cmd/
│   └── main.go                  # Entry point
├── internal/
│   ├── config/
│   │   ├── config.go            # Loads .env into Config struct
│   │   └── db.go                # Opens GORM DB connection
│   ├── auth/
│   │   └── jwt.go               # JWT service (generate + validate tokens)
│   ├── middlewares/
│   │   └── auth.go              # JWT auth middleware for protected routes
│   ├── httpresponse/
│   │   └── error.go             # Standard error response struct
│   ├── server/
│   │   └── http.go              # Creates Echo, auto-migrates DB, registers routes
│   └── domain/
│       ├── user/                # User registration, login, refresh, /me
│       │   ├── entity.go        # User model + bcrypt methods
│       │   ├── repository.go    # DB queries
│       │   ├── service.go       # Business logic
│       │   ├── handler.go       # HTTP handlers
│       │   ├── register.go      # Route registration
│       │   └── dto/             # Request/response structs
│       ├── mango/               # Mango inventory
│       │   ├── entity.go        # Mango model
│       │   ├── repository.go
│       │   ├── service.go
│       │   ├── handler.go
│       │   ├── register.go
│       │   └── dto/
│       └── order/               # Order placement
│           ├── entity.go        # Order model (pending/confirmed/cancelled)
│           ├── repository.go
│           ├── service.go
│           ├── handler.go
│           ├── register.go
│           └── dto/
├── .env                         # Environment variables (not in git)
├── .air.toml                    # Air hot reload config
├── go.mod                       # Module definition + dependency list
├── go.sum                       # Dependency checksums
└── mangoshop.postman_collection.json  # Postman collection for testing
```

Project Structure for Go tickets

```
gotickets/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── auth/                   # JWT token create/validate logic
│   ├── config/                 # Environment and database config
│   ├── domain/
│   │   ├── booking/            # Booking feature
│   │   │   └── dto/            # Booking request/response DTOs
│   │   ├── event/              # Event feature
│   │   │   └── dto/            # Event request/response DTOs
│   │   └── user/               # User auth/profile feature
│   │       └── dto/            # User request/response DTOs
│   ├── httpresponse/           # Common error response shape
│   ├── middlewares/            # Auth middleware
│   └── server/                 # Echo server setup
├── .air.toml                   # Air config for live reload
├── .env.example                # Example environment variables
├── go.mod                      # Go module and dependencies
└── go.sum
```
