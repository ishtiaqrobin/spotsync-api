package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}

type ReservationResponse struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// Used in "Get My Reservations" — nested zone info
type ReservationZoneInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type MyReservationResponse struct {
	ID           uint                `json:"id"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	Zone         ReservationZoneInfo `json:"zone"`
	CreatedAt    string              `json:"created_at"`
}
