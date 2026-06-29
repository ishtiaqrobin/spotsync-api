package reservation

import (
	"gorm.io/gorm"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/reservation/dto"
)

// Reservation represents the reservations table in the database
type Reservation struct {
	gorm.Model
	UserID       uint   `json:"user_id" gorm:"not null"`
	ZoneID       uint   `json:"zone_id" gorm:"not null"`
	LicensePlate string `json:"license_plate" gorm:"size:15;not null"`
	Status       string `json:"status" gorm:"type:varchar(15);default:'active'"`

	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Zone Zone `json:"zone,omitempty" gorm:"foreignKey:ZoneID"`
}

// TableName explicitly sets the table name
func (Reservation) TableName() string {
	return "reservations"
}

// User is a minimal reference for the relation (full model is in user domain)
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// TableName explicitly sets the table name for User struct
func (User) TableName() string {
	return "users"
}

// Zone is a minimal reference for the relation (full model is in zone domain)
type Zone struct {
	gorm.Model
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity"`
	PricePerHour  float64 `json:"price_per_hour"`
}

// TableName explicitly sets the table name for Zone struct
func (Zone) TableName() string {
	return "parking_zones"
}

// ToResponse converts a Reservation entity to a DTO response
func (r *Reservation) ToResponse() *dto.Response {
	return &dto.Response{
		ID:           r.ID,
		UserID:       r.UserID,
		ZoneID:       r.ZoneID,
		LicensePlate: r.LicensePlate,
		Status:       r.Status,
		CreatedAt:    r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    r.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// ToMyResponse converts a Reservation to a "my reservations" response with zone info
func (r *Reservation) ToMyResponse() *dto.MyReservationResponse {
	return &dto.MyReservationResponse{
		ID:           r.ID,
		LicensePlate: r.LicensePlate,
		Status:       r.Status,
		Zone: dto.ReservationZoneInfo{
			ID:   r.Zone.ID,
			Name: r.Zone.Name,
			Type: r.Zone.Type,
		},
		CreatedAt: r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
