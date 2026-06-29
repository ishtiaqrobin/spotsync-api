package dto

// Response represents a reservation response
type Response struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// MyReservationResponse represents a reservation in the "my reservations" list
type MyReservationResponse struct {
	ID           uint                `json:"id"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	Zone         ReservationZoneInfo `json:"zone"`
	CreatedAt    string              `json:"created_at"`
}

// ReservationZoneInfo represents minimal zone info in a reservation response
type ReservationZoneInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
