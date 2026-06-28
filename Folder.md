```
spotsync-api/
├── cmd/
│ └── main.go # entry point, DI wiring হবে এখানে
│
├── config/
│ ├── config.go # env loader (godotenv)
│ └── database.go # GORM postgres connection setup
│
├── models/
│ ├── user.go
│ ├── parking_zone.go
│ └── reservation.go
│
├── dto/
│ ├── auth_dto.go # RegisterRequest, LoginRequest, LoginResponse
│ ├── zone_dto.go # CreateZoneRequest, ZoneResponse
│ └── reservation_dto.go # CreateReservationRequest, ReservationResponse
│
├── repository/
│ ├── user_repository.go
│ ├── zone_repository.go
│ └── reservation_repository.go # transaction + row-lock logic এখানে
│
├── service/
│ ├── auth_service.go # bcrypt hash, JWT generate
│ ├── zone_service.go
│ └── reservation_service.go # business rules, capacity check
│
├── handler/
│ ├── auth_handler.go
│ ├── zone_handler.go
│ └── reservation_handler.go
│
├── middleware/
│ ├── jwt_middleware.go # token verify, inject claims into context
│ └── role_middleware.go # admin-only guard
│
├── routes/
│ └── routes.go # Echo route registration, grouped by /api/v1
│
├── utils/
│ ├── response.go # SuccessResponse / ErrorResponse helpers
│ ├── jwt_util.go # GenerateToken, ParseToken
│ └── errors.go # custom errors (ErrZoneFull, ErrNotFound, etc.)
│
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── air.toml # hot-reload config (Air)
└── README.md
```
